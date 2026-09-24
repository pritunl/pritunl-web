package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/pritunl/pritunl-web/request"
	"github.com/pritunl/pritunl-web/utils"
)

func adminGet(c *gin.Context) {
	adminId := utils.FilterId(c.Params.ByName("admin_id"))
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
	Username     string `json:"username"`
	Password     string `json:"password"`
	YubikeyId    string `json:"yubikey_id"`
	SuperUser    bool   `json:"super_user"`
	AuthApi      bool   `json:"auth_api"`
	Token        string `json:"token"`
	Secret       string `json:"secret"`
	Disabled     bool   `json:"disabled"`
	OtpAuth      bool   `json:"otp_auth"`
	OtpSecret    bool   `json:"otp_secret"`
	LocalOtpAuth bool   `json:"local_otp_auth"`
}

func (d *adminPutData) Filter() {
	d.Username = utils.FilterStr(d.Username, 1024)
	d.Password = utils.FilterText(d.Password, 512)
	d.YubikeyId = utils.FilterId(d.YubikeyId)
	d.Token = utils.FilterId(d.Token)
	d.Secret = utils.FilterId(d.Secret)
}

func adminPut(c *gin.Context) {
	adminId := utils.FilterId(c.Params.ByName("admin_id"))
	data := &adminPutData{}

	req := &request.Request{
		Method: "PUT",
		Path:   "/admin/" + adminId,
		Json:   data,
	}

	req.Do(c)
}

type adminPostData struct {
	Username     string `json:"username"`
	Password     string `json:"password"`
	YubikeyId    string `json:"yubikey_id"`
	OtpAuth      bool   `json:"otp_auth"`
	LocalOtpAuth bool   `json:"local_otp_auth"`
	AuthApi      bool   `json:"auth_api"`
	Disabled     bool   `json:"disabled"`
	SuperUser    bool   `json:"super_user"`
}

func (d *adminPostData) Filter() {
	d.Username = utils.FilterStr(d.Username, 1024)
	d.Password = utils.FilterText(d.Password, 512)
	d.YubikeyId = utils.FilterId(d.YubikeyId)
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
	adminId := utils.FilterId(c.Params.ByName("admin_id"))

	req := &request.Request{
		Method: "DELETE",
		Path:   "/admin/" + adminId,
	}

	req.Do(c)
}

func adminAuditGet(c *gin.Context) {
	adminId := utils.FilterId(c.Params.ByName("admin_id"))

	req := &request.Request{
		Method: "GET",
		Path:   "/admin/" + adminId + "/audit",
	}

	req.Do(c)
}
