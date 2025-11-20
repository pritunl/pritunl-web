package request

import (
	"fmt"
	"net/http"
	"net/url"
	"strings"

	"github.com/dropbox/godropbox/errors"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/render"
	"github.com/pritunl/pritunl-web/errortypes"
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

func AbortWithStatus(c *gin.Context, code int, msg string) {
	r := render.String{
		Format: fmt.Sprintf("%d %s", code, msg),
	}

	c.Status(code)
	r.WriteContentType(c.Writer)
	c.Writer.WriteHeaderNow()
	r.Render(c.Writer)
	c.Abort()
}

func AbortRedirect(c *gin.Context, path string) {
	u, err := url.Parse(c.Request.URL.String())
	if err != nil {
		err = errortypes.RequestError{
			errors.Wrap(err, "request: URL parse error"),
		}
		c.AbortWithError(418, err)
		return
	}

	u.Path = path
	u.RawQuery = ""

	c.Redirect(http.StatusFound, u.String())
	c.Abort()
}
