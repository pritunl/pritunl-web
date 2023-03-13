package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/pritunl/pritunl-web/request"
	"github.com/pritunl/pritunl-web/utils"
)

func adminGet(c *gin.Context) {
	adminId := utils.FilterStr(c.Params.ByName("admin_id"), 128)
	if adminId != "" {
		adminId = "/" + adminId
	}

	req := &request.Request{
		Method: "GET",
		Path:   "/admin" + adminId,
	}

	req.Do(c)
}

type adminPutData struct {
	Username  string `json:"username"`
	Password  string `json:"password"`
	YubikeyId string `json:"yubikey_id"`
	SuperUser bool   `json:"super_user"`
	AuthApi   bool   `json:"auth_api"`
	Token     string `json:"token"`
	Secret    string `json:"secret"`
	Disabled  bool   `json:"disabled"`
	OtpAuth   bool   `json:"otp_auth"`
	OtpSecret bool   `json:"otp_secret"`
}

func adminPut(c *gin.Context) {
	adminId := utils.FilterStr(c.Params.ByName("admin_id"), 128)
	data := &adminPutData{}

	req := &request.Request{
		Method: "PUT",
		Path:   "/admin/" + adminId,
		Json:   data,
	}

	req.Do(c)
}

type adminPostData struct {
	Username  string `json:"username"`
	Password  string `json:"password"`
	YubikeyId string `json:"yubikey_id"`
	OtpAuth   bool   `json:"otp_auth"`
	AuthApi   bool   `json:"auth_api"`
	Disabled  bool   `json:"disabled"`
	SuperUser bool   `json:"super_user"`
}

func adminPost(c *gin.Context) {
	data := &adminPostData{}

	req := &request.Request{
		Method: "POST",
		Path:   "/admin",
		Json:   data,
	}

	req.Do(c)
}

func adminDelete(c *gin.Context) {
	adminId := utils.FilterStr(c.Params.ByName("admin_id"), 128)

	req := &request.Request{
		Method: "DELETE",
		Path:   "/admin/" + adminId,
	}

	req.Do(c)
}

func adminAuditGet(c *gin.Context) {
	adminId := utils.FilterStr(c.Params.ByName("admin_id"), 128)

	req := &request.Request{
		Method: "GET",
		Path:   "/admin/" + adminId + "/audit",
	}

	req.Do(c)
}
