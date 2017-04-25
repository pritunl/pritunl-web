package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/pritunl/pritunl-web/request"
)

type linkStatePutData struct {
	PublicAddress string `json:"public_address"`
	Tunnels       int    `json:"tunnels"`
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
