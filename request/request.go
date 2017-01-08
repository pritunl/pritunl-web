package request

import (
	"bytes"
	"encoding/json"
	"github.com/dropbox/godropbox/errors"
	"github.com/gin-gonic/gin"
	"github.com/pritunl/pritunl-web/constants"
	"github.com/pritunl/pritunl-web/errortypes"
	"io"
	"net/http"
	"net/url"
)

var client = &http.Client{
	CheckRedirect: func(req *http.Request, via []*http.Request) error {
		return http.ErrUseLastResponse
	},
}

type Request struct {
	Method   string
	Path     string
	Headers  []string
	Query    map[string]string
	RawQuery string
	Json     interface{}
}

func (r *Request) Do(c *gin.Context) {
	reqUrl := "http://" + constants.InternalHost + r.Path

	var body io.Reader

	if r.Json != nil {
		if c.ContentType() != "application/json" {
			err := errortypes.RequestError{
				errors.New("request: Invalid content type"),
			}
			c.AbortWithError(500, err)
			return
		}

		err := c.BindJSON(r.Json)
		if err != nil {
			return
		}

		data, err := json.Marshal(r.Json)
		if err != nil {
			err = errortypes.RequestError{
				errors.Wrap(err, "request: Json marshal error"),
			}
			c.AbortWithError(500, err)
			return
		}

		body = bytes.NewBuffer(data)
	}

	req, err := http.NewRequest(r.Method, reqUrl, body)
	if err != nil {
		err = errortypes.RequestError{
			errors.Wrap(err, "request: Create request failed"),
		}
		c.AbortWithError(500, err)
		return
	}

	forwardUrl := url.URL{
		Scheme: constants.Scheme,
		Host:   c.Request.Host,
	}

	if r.Query != nil {
		query := req.URL.Query()

		for key, val := range r.Query {
			query.Add(key, val)
		}

		req.URL.RawQuery = query.Encode()
	}

	if r.RawQuery != "" {
		req.URL.RawQuery = r.RawQuery
	}

	req.Header.Set("PR-Forwarded-Header",
		req.Header.Get(constants.ReverseProxyHeader))
	req.Header.Set("PR-Forwarded-Url", forwardUrl.String())
	req.Header.Set("PR-Forwarded-For",
		parseRemoteAddr(c.Request.RemoteAddr))

	copyHeader(req, c.Request, "Auth-Token")
	copyHeader(req, c.Request, "Auth-Timestamp")
	copyHeader(req, c.Request, "Auth-Nonce")
	copyHeader(req, c.Request, "Auth-Signature")

	copyHeader(req, c.Request, "Cookie")
	copyHeader(req, c.Request, "Csrf-Token")

	if r.Headers != nil {
		for _, key := range r.Headers {
			copyHeader(req, c.Request, key)
		}
	}

	if r.Json != nil {
		req.Header.Set("Content-Type", "application/json")
	}

	resp, err := client.Do(req)
	if err != nil {
		err = errortypes.RequestError{
			errors.Wrap(err, "request: Request failed"),
		}
		c.AbortWithError(500, err)
		return
	}
	defer resp.Body.Close()

	copyHeaders(c.Writer.Header(), resp.Header)
	c.Writer.Header().Del("Server")
	c.Writer.WriteHeader(resp.StatusCode)
	io.Copy(c.Writer, resp.Body)
}
