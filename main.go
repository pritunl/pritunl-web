package main

import (
	"crypto/tls"
	"fmt"
	"github.com/Sirupsen/logrus"
	"github.com/gin-gonic/gin"
	"github.com/pritunl/pritunl-web/constants"
	"github.com/pritunl/pritunl-web/handlers"
	"github.com/pritunl/pritunl-web/request"
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"
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
	if constants.RedirectServer == "true" && constants.BindPort != "80" {
		go func() {
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
		Addr:         constants.BindHost + ":" + constants.BindPort,
		Handler:      router,
		ReadTimeout:  2 * time.Minute,
		WriteTimeout: 2 * time.Minute,
	}

	var err error
	if constants.Ssl {
		server.TLSConfig = &tls.Config{
			MinVersion: tls.VersionTLS12,
			MaxVersion: tls.VersionTLS12,
		}

		err = server.ListenAndServeTLS(
			constants.CertPath,
			constants.KeyPath,
		)
	} else {
		err = server.ListenAndServe()
	}
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"error": err,
		}).Error("main: Server error")
		panic(err)
	}
}
