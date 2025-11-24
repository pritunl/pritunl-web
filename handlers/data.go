package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/pritunl/pritunl-web/request"
	"github.com/pritunl/pritunl-web/utils"
)

func dataKeyGet(c *gin.Context) {
	orgId := utils.FilterStr(c.Params.ByName("org_id"), 128)
	userId := utils.FilterStr(c.Params.ByName("user_id"), 128)

	req := &request.Request{
		Method: "GET",
		Path:   "/data/" + orgId + "/" + userId,
	}

	req.Do(c)
}

func dataServerKeyGet(c *gin.Context) {
	orgId := utils.FilterStr(c.Params.ByName("org_id"), 128)
	userId := utils.FilterStr(c.Params.ByName("user_id"), 128)
	serverId := utils.FilterStr(c.Params.ByName("server_id"), 128)

	req := &request.Request{
		Method: "GET",
		Path:   "/data/" + orgId + "/" + userId + "/" + serverId,
	}

	req.Do(c)
}
