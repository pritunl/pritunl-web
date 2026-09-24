package request

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strconv"
	"time"

	"github.com/dropbox/godropbox/errors"
	"github.com/gin-gonic/gin"
	"github.com/pritunl/pritunl-web/constants"
	"github.com/pritunl/pritunl-web/errortypes"
	"github.com/pritunl/pritunl-web/utils"
	"github.com/pritunl/tools/logger"
)

var client = &http.Client{
	CheckRedirect: func(req *http.Request, via []*http.Request) error {
		return http.ErrUseLastResponse
	},
	Timeout: 2 * time.Minute,
}

type Filterer interface {
	Filter()
}

type Request struct {
	Method   string
	Path     string
	Headers  []string
	Query    map[string]string
	RawQuery string
	Json     Filterer
}

func (r *Request) Send(c *gin.Context) (resp *http.Response, err error) {
	reqUrl := "http://" + constants.InternalHost + r.Path

	var body io.Reader

	if r.Json != nil {
		if c.ContentType() != "application/json" {
			err = errortypes.RequestError{
				errors.New("request: Invalid content type"),
			}
			c.AbortWithError(520, err)
			return
		}

		err = c.BindJSON(r.Json)
		if err != nil {
			return
		}

		r.Json.Filter()

		data, e := json.Marshal(r.Json)
		if e != nil {
			err = errortypes.RequestError{
				errors.Wrap(e, "request: Json marshal error"),
			}
			logger.WithFields(logger.Fields{
				"error": err,
			}).Error("request: Request error")
			c.AbortWithError(521, err)
			return
		}

		body = bytes.NewBuffer(data)
	}

	req, err := http.NewRequest(r.Method, reqUrl, body)
	if err != nil {
		err = errortypes.RequestError{
			errors.Wrap(err, "request: Create request failed"),
		}
		logger.WithFields(logger.Fields{
			"error": err,
		}).Error("request: Request error")
		c.AbortWithError(522, err)
		return
	}

	forwardUrl := url.URL{
		Scheme: constants.Scheme,
		Host:   utils.FilterDomain(c.Request.Host),
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

	req.Header.Set(
		"PR-Validated",
		strconv.FormatBool(c.MustGet("validated").(bool)),
	)
	req.Header.Set("PR-Forwarded-Header", filterForwardedHeader(
		c.Request.Header.Get(constants.ReverseProxyHeader)))
	req.Header.Set("PR-Forwarded-Url", forwardUrl.String())
	req.Header.Set("PR-Forwarded-For",
		parseRemoteAddr(c.Request.RemoteAddr))

	copyHeader(req, c.Request, "Auth-Token", utils.FilterId)
	copyHeader(req, c.Request, "Auth-Timestamp", utils.FilterId)
	copyHeader(req, c.Request, "Auth-Nonce", utils.FilterId)
	copyHeader(req, c.Request, "Auth-Signature", utils.FilterBase64)

	copyHeader(req, c.Request, "Cookie", utils.FilterOpen)
	copyHeader(req, c.Request, "Csrf-Token", utils.FilterId)

	if r.Headers != nil {
		for _, key := range r.Headers {
			copyHeader(req, c.Request, key, utils.FilterOpen)
		}
	}

	if r.Json != nil {
		req.Header.Set("Content-Type", "application/json")
	}

	resp, err = client.Do(req)
	if err != nil {
		err = errortypes.RequestError{
			errors.Wrap(err, "request: Request failed"),
		}
		logger.WithFields(logger.Fields{
			"error": err,
		}).Error("request: Request error")
		c.AbortWithError(502, err)
		return
	}

	return
}

func (r *Request) Do(c *gin.Context) {
	resp, err := r.Send(c)
	if err != nil {
		return
	}
	defer resp.Body.Close()

	copyHeaders(c.Writer.Header(), resp.Header)
	c.Writer.Header().Del("Server")
	c.Writer.WriteHeader(resp.StatusCode)
	io.Copy(c.Writer, resp.Body)
}

func WriteError(w http.ResponseWriter, code int, err error) {
	http.Error(w, fmt.Sprintf("%d %s", code, http.StatusText(code)), code)
}

func DoCheck(w http.ResponseWriter, r *http.Request) {
	reqUrl := "http://" + constants.InternalHost + "/check"

	req, err := http.NewRequest("GET", reqUrl, nil)
	if err != nil {
		err = errortypes.RequestError{
			errors.Wrap(err, "request: Create request failed"),
		}
		logger.WithFields(logger.Fields{
			"error": err,
		}).Error("request: Request error")
		WriteError(w, 525, err)
		return
	}

	forwardUrl := url.URL{
		Scheme: constants.Scheme,
		Host:   utils.FilterDomain(r.Host),
	}

	req.Header.Set("PR-Forwarded-Header", filterForwardedHeader(
		r.Header.Get(constants.ReverseProxyHeader)))
	req.Header.Set("PR-Forwarded-Url", forwardUrl.String())
	req.Header.Set("PR-Forwarded-For", parseRemoteAddr(r.RemoteAddr))

	resp, err := client.Do(req)
	if err != nil {
		err = errortypes.RequestError{
			errors.Wrap(err, "request: Request failed"),
		}
		logger.WithFields(logger.Fields{
			"error": err,
		}).Error("request: Request error")
		WriteError(w, 502, err)
		return
	}
	defer resp.Body.Close()

	copyHeaders(w.Header(), resp.Header)
	w.Header().Del("Server")
	w.WriteHeader(resp.StatusCode)
	io.Copy(w, resp.Body)
}
