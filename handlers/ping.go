package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/issuu/pritunl-web/request"
)

func pingGet(c *gin.Context) {
	req := &request.Request{
		Method: "GET",
		Path:   "/ping",
	}

	req.Do(c)
}

func checkGet(c *gin.Context) {
	req := &request.Request{
		Method: "GET",
		Path:   "/check",
	}

	req.Do(c)
}
