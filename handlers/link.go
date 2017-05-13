package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/pritunl/pritunl-web/request"
)

func linkGet(c *gin.Context) {
	var query map[string]string
	page := c.Query("page")
	if page != "" {
		query = map[string]string{
			"page": page,
		}
	}

	req := &request.Request{
		Method: "GET",
		Path:   "/link",
		Query:  query,
	}

	req.Do(c)
}

type linkPutData struct {
	Name    string `json:"name"`
	Status  string `json:"status"`
	Timeout int    `json:"timeout"`
}

func linkPut(c *gin.Context) {
	linkId := c.Params.ByName("link_id")
	data := &linkPutData{}

	req := &request.Request{
		Method: "PUT",
		Path:   "/link/" + linkId,
		Json:   data,
	}

	req.Do(c)
}

func linkLocationGet(c *gin.Context) {
	linkId := c.Params.ByName("link_id")

	req := &request.Request{
		Method: "GET",
		Path:   "/link/" + linkId + "/location",
	}

	req.Do(c)
}

type linkStatePutData struct {
	Version       string            `json:"version"`
	PublicAddress string            `json:"public_address"`
	Status        map[string]string `json:"status"`
	Errors        []string          `json:"errors"`
}

func linkStatePut(c *gin.Context) {
	data := &linkStatePutData{}

	req := &request.Request{
		Method: "PUT",
		Path:   "/link/state",
		Json:   data,
	}

	req.Do(c)
}
