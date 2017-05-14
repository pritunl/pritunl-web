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

type linkPostData struct {
	Name   string `json:"name"`
	Status string `json:"status"`
}

func linkPost(c *gin.Context) {
	data := &linkPutData{}

	req := &request.Request{
		Method: "POST",
		Path:   "/link",
		Json:   data,
	}

	req.Do(c)
}

type linkPutData struct {
	Name    string `json:"name"`
	Status  string `json:"status"`
	Timeout int    `json:"timeout"`
}

type linkStatePutData struct {
	Version       string            `json:"version"`
	PublicAddress string            `json:"public_address"`
	Status        map[string]string `json:"status"`
	Errors        []string          `json:"errors"`
}

func linkPut(c *gin.Context) {
	var req *request.Request
	linkId := c.Params.ByName("link_id")

	if linkId == "state" {
		data := &linkStatePutData{}

		req = &request.Request{
			Method: "PUT",
			Path:   "/link/state",
			Json:   data,
		}
	} else {
		data := &linkPutData{}

		req = &request.Request{
			Method: "PUT",
			Path:   "/link/" + linkId,
			Json:   data,
		}
	}

	req.Do(c)
}

func linkDelete(c *gin.Context) {
	linkId := c.Params.ByName("link_id")

	req := &request.Request{
		Method: "DELETE",
		Path:   "/link/" + linkId,
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
