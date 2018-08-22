package main

import (
	"crypto/tls"
	"fmt"
	log "github.com/sirupsen/logrus"
	logrus_syslog "github.com/sirupsen/logrus/hooks/syslog"
	"log/syslog"
	"github.com/gin-gonic/gin"
	"github.com/issuu/pritunl-web/constants"
	"github.com/issuu/pritunl-web/handlers"
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
	hook, hook_err := logrus_syslog.NewSyslogHook("", "", syslog.LOG_DEBUG, "")
	if hook_err != nil {
		log.Error("Unable to connect to local syslog daemon")
	} else {
		log.AddHook(hook)
	}

	if constants.RedirectServer == "true" && constants.BindPort != "80" {
		go func() {
			server := http.Server{
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
							http.Error(w, "", http.StatusInternalServerError)
							return
						}
						defer resp.Body.Close()

						copyHeader(w.Header(), resp.Header)
						w.WriteHeader(resp.StatusCode)
						io.Copy(w, resp.Body)

						return
					} else if strings.HasPrefix(req.URL.Path, "/check") ||
						strings.HasPrefix(req.URL.Path, "/ping") {

						w.Header().Set(
							"Content-Type",
							"text/plain; charset=utf-8",
						)
						w.WriteHeader(200)
						fmt.Fprintln(w, "200 OK")

						return
					}

					if constants.ServerHost != "" {
						req.URL.Host = constants.ServerHost
					} else {
						req.URL.Host = req.Host
					}
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
				log.WithFields(log.Fields{
					"error": err,
				}).Error("main: Redirect server error")
			}
		}()
	}

	gin.SetMode(gin.ReleaseMode)

	router := gin.New()
	handlers.Register(router)

	server := http.Server{
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
		log.WithFields(log.Fields{
			"error": err,
		}).Error("main: Server error")
		panic(err)
	}
}
