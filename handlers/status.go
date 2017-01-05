package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/pritunl/pritunl-web/request"
)

func statusGet(c *gin.Context) {
	req := &request.Request{
		Method: "GET",
		Path:   "/status",
	}

	req.Do(c)
}
