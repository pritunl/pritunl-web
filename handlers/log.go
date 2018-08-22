package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/issuu/pritunl-web/request"
)

func logGet(c *gin.Context) {
	req := &request.Request{
		Method: "GET",
		Path:   "/log",
	}

	req.Do(c)
}

func logsGet(c *gin.Context) {
	req := &request.Request{
		Method: "GET",
		Path:   "/logs",
	}

	req.Do(c)
}
