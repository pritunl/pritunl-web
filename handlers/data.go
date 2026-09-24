package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/pritunl/pritunl-web/request"
	"github.com/pritunl/pritunl-web/utils"
)

func dataKeyGet(c *gin.Context) {
	orgId := utils.FilterId(c.Params.ByName("org_id"))
	userId := utils.FilterStr(c.Params.ByName("user_id"), 512)

	req := &request.Request{
		Method: "GET",
		Path:   "/data/" + orgId + "/" + userId,
	}

	req.Do(c)
}

func dataServerKeyGet(c *gin.Context) {
	orgId := utils.FilterId(c.Params.ByName("org_id"))
	userId := utils.FilterId(c.Params.ByName("user_id"))
	serverId := utils.FilterStr(c.Params.ByName("server_id"), 512)

	req := &request.Request{
		Method: "GET",
		Path:   "/data/" + orgId + "/" + userId + "/" + serverId,
	}

	req.Do(c)
}
