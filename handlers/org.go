package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/pritunl/pritunl-web/request"
	"github.com/pritunl/pritunl-web/utils"
)

func orgGet(c *gin.Context) {
	orgId := utils.FilterStr(c.Params.ByName("org_id"), 128)
	if orgId != "" {
		orgId = "/" + orgId
	}

	var query map[string]string
	page := c.Query("page")
	if page != "" {
		query = map[string]string{
			"page": page,
		}
	}

	req := &request.Request{
		Method: "GET",
		Path:   "/organization" + orgId,
		Query:  query,
	}

	req.Do(c)
}

type orgPostData struct {
	Name    string `json:"name"`
	AuthApi bool   `json:"auth_api"`
}

func orgPost(c *gin.Context) {
	data := &orgPostData{}

	req := &request.Request{
		Method: "POST",
		Path:   "/organization",
		Json:   data,
	}

	req.Do(c)
}

type orgPutData struct {
	Name       string `json:"name"`
	AuthApi    bool   `json:"auth_api"`
	AuthToken  bool   `json:"auth_token"`
	AuthSecret bool   `json:"auth_secret"`
}

func orgPut(c *gin.Context) {
	orgId := utils.FilterStr(c.Params.ByName("org_id"), 128)
	data := &orgPutData{}

	req := &request.Request{
		Method: "PUT",
		Path:   "/organization/" + orgId,
		Json:   data,
	}

	req.Do(c)
}

func orgDelete(c *gin.Context) {
	orgId := utils.FilterStr(c.Params.ByName("org_id"), 128)

	req := &request.Request{
		Method: "DELETE",
		Path:   "/organization/" + orgId,
	}

	req.Do(c)
}
