package main

import (
	"github.com/Sirupsen/logrus"
	"github.com/gin-gonic/gin"
	"github.com/pritunl/pritunl-web/constants"
	"github.com/pritunl/pritunl-web/handlers"
	"io"
	"net/http"
	"net/url"
	"strings"
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
		go http.ListenAndServe(constants.BindHost+":80", http.HandlerFunc(func(
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
					panic(err)
					http.Error(w, "", http.StatusInternalServerError)
					return
				}
				defer resp.Body.Close()

				copyHeader(w.Header(), resp.Header)
				w.WriteHeader(resp.StatusCode)
				io.Copy(w, resp.Body)

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
		}))
	}

	router := gin.New()
	handlers.Register(router)

	var err error
	if constants.Ssl {
		err = router.RunTLS(constants.BindHost+":"+constants.BindPort,
			constants.CertPath, constants.KeyPath)
	} else {
		err = router.Run(constants.BindHost + ":" + constants.BindPort)
	}
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"error": err,
		}).Error("main: Server error")
		panic(err)
	}
}
