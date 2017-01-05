package request

import (
	"net/http"
	"strings"
)

func copyHeader(dst, src *http.Request, key string) {
	val := src.Header.Get(key)
	if val != "" {
		dst.Header.Set(key, val)
	}
}

func copyHeaders(dst, src http.Header) {
	for k, vv := range src {
		for _, v := range vv {
			dst.Add(k, v)
		}
	}
}

func parseRemoteAddr(remoteAddr string) (addr string) {
	addr = remoteAddr[:strings.LastIndex(remoteAddr, ":")]
	addr = strings.Replace(addr, "[", "", 1)
	addr = strings.Replace(addr, "]", "", 1)
	return
}
