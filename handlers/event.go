package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/pritunl/pritunl-web/request"
	"github.com/pritunl/pritunl-web/utils"
)

func eventGet(c *gin.Context) {
	cursor := utils.FilterId(c.Params.ByName("cursor"))
	if cursor != "" {
		cursor = "/" + cursor
	}

	req := &request.Request{
		Method: "GET",
		Path:   "/event" + cursor,
	}

	req.Do(c)
}
