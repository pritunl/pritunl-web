package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/pritunl/pritunl-web/request"
)

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
