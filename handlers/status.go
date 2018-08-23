package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/issuu/pritunl-web/request"
)

func statusGet(c *gin.Context) {
	req := &request.Request{
		Method: "GET",
		Path:   "/status",
	}

	req.Do(c)
}
