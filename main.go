package main

import (
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"
	"strings"
)

func ParseRemoteAddr(remoteAddr string) (addr string) {
	addr = remoteAddr[:strings.LastIndex(remoteAddr, ":")]
	addr = strings.Replace(addr, "[", "", 1)
	addr = strings.Replace(addr, "]", "", 1)
	return
}

func main() {
	redirectServer := os.Getenv("REDIRECT_SERVER")
	bindHost := os.Getenv("BIND_HOST")
	bindPort := os.Getenv("BIND_PORT")
	internalHost := os.Getenv("INTERNAL_ADDRESS")
	certPath := os.Getenv("CERT_PATH")
	keyPath := os.Getenv("KEY_PATH")
	ssl := certPath != "" && keyPath != ""

	if redirectServer == "true" && bindPort != "80" {
		go http.ListenAndServe(bindHost+":80", http.HandlerFunc(func(
			w http.ResponseWriter, req *http.Request) {

			if ssl {
				req.URL.Scheme = "https"
			} else {
				req.URL.Scheme = "http"
			}

			req.URL.Host = req.Host
			if bindPort != "443" {
				req.URL.Host += ":" + bindPort
			}

			http.Redirect(w, req, req.URL.String(),
				http.StatusMovedPermanently)
		}))
	}

	proxy := &httputil.ReverseProxy{
		Director: func(req *http.Request) {
			forwardUrl := url.URL{
				Host: req.Host,
			}
			if ssl {
				forwardUrl.Scheme = "https"
			} else {
				forwardUrl.Scheme = "http"
			}

			req.Header.Set("PR-Forward-Url", forwardUrl.String())
			req.Header.Set("PR-Forward-For", ParseRemoteAddr(req.RemoteAddr))

			req.URL.Scheme = "http"
			req.URL.Host = internalHost
		},
	}

	var err error
	if ssl {
		err = http.ListenAndServeTLS(
			bindHost+":"+bindPort,
			certPath,
			keyPath,
			proxy,
		)
	} else {
		err = http.ListenAndServe(
			bindHost+":"+bindPort,
			proxy,
		)
	}
	if err != nil {
		panic(err)
	}
}
