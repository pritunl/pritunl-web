package main

import (
	"crypto/tls"
	"encoding/base64"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"

	"github.com/dropbox/godropbox/errors"
	"github.com/gin-gonic/gin"
	"github.com/pritunl/pritunl-web/constants"
	"github.com/pritunl/pritunl-web/errortypes"
	"github.com/pritunl/pritunl-web/handlers"
	"github.com/pritunl/pritunl-web/request"
	"github.com/sirupsen/logrus"
)

func ParseRemoteAddr(remoteAddr string) (addr string) {
	addr = remoteAddr[:strings.LastIndex(remoteAddr, ":")]
	addr = strings.Replace(addr, "[", "", 1)
	addr = strings.Replace(addr, "]", "", 1)
	return
}

func copyHeader(dst, src http.Header) {
	for k, vv := range src {
		for _, v := range vv {
			dst.Add(k, v)
		}
	}
}

func main() {
	constants.ReverseProxyHeader = os.Getenv("REVERSE_PROXY_HEADER")
	constants.ReverseProxyProtoHeader = os.Getenv("REVERSE_PROXY_PROTO_HEADER")
	constants.RedirectServer = os.Getenv("REDIRECT_SERVER")
	constants.BindHost = os.Getenv("BIND_HOST")
	constants.BindPort = os.Getenv("BIND_PORT")
	constants.InternalHost = os.Getenv("INTERNAL_ADDRESS")
	constants.SslCert = os.Getenv("SSL_CERT")
	constants.SslKey = os.Getenv("SSL_KEY")
	webStrictStr := os.Getenv("WEB_STRICT")
	webSecretStr := os.Getenv("WEB_SECRET")
	os.Unsetenv("REVERSE_PROXY_HEADER")
	os.Unsetenv("REVERSE_PROXY_PROTO_HEADER")
	os.Unsetenv("REDIRECT_SERVER")
	os.Unsetenv("BIND_HOST")
	os.Unsetenv("BIND_PORT")
	os.Unsetenv("INTERNAL_ADDRESS")
	os.Unsetenv("SSL_CERT")
	os.Unsetenv("SSL_KEY")
	os.Unsetenv("WEB_STRICT")
	os.Unsetenv("WEB_SECRET")

	constants.Ssl = constants.SslCert != "" && constants.SslKey != ""
	if constants.Ssl {
		constants.Scheme = "https"
	} else {
		constants.Scheme = "http"
	}

	if webStrictStr == "false" {
		constants.WebStrict = false
	} else {
		constants.WebStrict = true
	}

	var err error
	if webSecretStr != "" {
		webSecretByt, e := base64.StdEncoding.DecodeString(webSecretStr)
		if e != nil {
			err = &errortypes.ParseError{
				errors.Wrap(e, "main: Failed to decode web secret"),
			}

			logrus.WithFields(logrus.Fields{
				"error": err,
			}).Error("main: Failed to decode web secret")

			panic(err)
		}
		constants.WebSecret = &[32]byte{}
		copy(constants.WebSecret[:], webSecretByt)
	}

	if constants.RedirectServer == "true" && constants.BindPort != "80" {
		go func() {
			logrus.WithFields(logrus.Fields{
				"port": 80,
			}).Info("main: Starting HTTP redirect server")

			server := &http.Server{
				Addr:         constants.BindHost + ":80",
				ReadTimeout:  1 * time.Minute,
				WriteTimeout: 1 * time.Minute,
				Handler: http.HandlerFunc(func(
					w http.ResponseWriter, req *http.Request) {

					if strings.HasPrefix(req.URL.Path,
						"/.well-known/acme-challenge/") {

						pathSplit := strings.Split(req.URL.Path, "/")
						token := pathSplit[len(pathSplit)-1]

						acmeUrl := url.URL{
							Scheme: "http",
							Host:   constants.InternalHost,
							Path:   "/.well-known/acme-challenge/" + token,
						}

						resp, err := http.Get(acmeUrl.String())
						if err != nil {
							http.Error(
								w,
								fmt.Sprintf(
									"%d %s",
									http.StatusInternalServerError,
									http.StatusText(
										http.StatusInternalServerError,
									),
								),
								http.StatusInternalServerError,
							)
							return
						}
						defer resp.Body.Close()

						copyHeader(w.Header(), resp.Header)
						w.Header().Del("Server")
						w.WriteHeader(resp.StatusCode)
						io.Copy(w, resp.Body)

						return
					} else if strings.HasPrefix(req.URL.Path, "/check") ||
						strings.HasPrefix(req.URL.Path, "/ping") {

						request.DoCheck(w, req)
						return
					}

					req.URL.Host = req.Host
					if constants.ReverseProxyHeader != "" &&
						req.Header.Get(constants.ReverseProxyHeader) != "" {

						req.URL.Scheme = "https"
					} else {
						req.URL.Scheme = constants.Scheme

						if constants.BindPort != "443" {
							req.URL.Host += ":" + constants.BindPort
						}
					}

					http.Redirect(w, req, req.URL.String(),
						http.StatusMovedPermanently)
				}),
			}

			err := server.ListenAndServe()
			if err != nil {
				logrus.WithFields(logrus.Fields{
					"error": err,
				}).Error("main: Redirect server error")
			}
		}()
	}

	gin.SetMode(gin.ReleaseMode)

	router := gin.New()
	handlers.Register(router)

	server := &http.Server{
		Addr:              constants.BindHost + ":" + constants.BindPort,
		Handler:           router,
		ReadTimeout:       2 * time.Minute,
		ReadHeaderTimeout: 30 * time.Second,
		WriteTimeout:      2 * time.Minute,
		IdleTimeout:       1 * time.Minute,
		MaxHeaderBytes:    500000,
	}

	if constants.Ssl {
		logrus.WithFields(logrus.Fields{
			"port": constants.BindPort,
		}).Info("main: Starting HTTPS server")

		sslCertByt, e := base64.StdEncoding.DecodeString(constants.SslCert)
		if e != nil {
			logrus.WithFields(logrus.Fields{
				"error": e,
			}).Error("main: Server cert decode error")
			panic(e)
		}

		sslKeyByt, e := base64.StdEncoding.DecodeString(constants.SslKey)
		if e != nil {
			logrus.WithFields(logrus.Fields{
				"error": e,
			}).Error("main: Server key decode error")
			panic(e)
		}

		tlsCert, e := tls.X509KeyPair(sslCertByt, sslKeyByt)
		if e != nil {
			logrus.WithFields(logrus.Fields{
				"error": e,
			}).Error("main: Server tls decode error")
			panic(e)
		}

		server.TLSConfig = &tls.Config{
			MinVersion: tls.VersionTLS12,
			MaxVersion: tls.VersionTLS13,
			CipherSuites: []uint16{
				tls.TLS_AES_128_GCM_SHA256,                        // 0x1301
				tls.TLS_AES_256_GCM_SHA384,                        // 0x1302
				tls.TLS_CHACHA20_POLY1305_SHA256,                  // 0x1303
				tls.TLS_ECDHE_ECDSA_WITH_AES_128_GCM_SHA256,       // 0xc02b
				tls.TLS_ECDHE_RSA_WITH_AES_128_GCM_SHA256,         // 0xc02f
				tls.TLS_ECDHE_ECDSA_WITH_AES_256_GCM_SHA384,       // 0xc02c
				tls.TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384,         // 0xc030
				tls.TLS_ECDHE_ECDSA_WITH_CHACHA20_POLY1305_SHA256, // 0xcca9
				tls.TLS_ECDHE_RSA_WITH_CHACHA20_POLY1305_SHA256,   // 0xcca8
			},
			Certificates: []tls.Certificate{
				tlsCert,
			},
		}

		err = server.ListenAndServeTLS("", "")
	} else {
		logrus.WithFields(logrus.Fields{
			"port": constants.BindPort,
		}).Info("main: Starting HTTP server")

		err = server.ListenAndServe()
	}
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"error": err,
		}).Error("main: Server error")
		panic(err)
	}
}
