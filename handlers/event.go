package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/issuu/pritunl-web/request"
)

func eventGet(c *gin.Context) {
	cursor := c.Params.ByName("cursor")
	if cursor != "" {
		cursor = "/" + cursor
	}

	req := &request.Request{
		Method: "GET",
		Path:   "/event" + cursor,
	}

	req.Do(c)
}
