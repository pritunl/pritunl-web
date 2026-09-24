package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/pritunl/pritunl-web/request"
	"github.com/pritunl/pritunl-web/utils"
)

type authSessionPostData struct {
	Username  string `json:"username"`
	Password  string `json:"password"`
	YubicoKey string `json:"yubico_key"`
	OtpCode   string `json:"otp_code"`
}

func (d *authSessionPostData) Filter() {
	d.Username = utils.FilterStr(d.Username, 1024)
	d.Password = utils.FilterText(d.Password, 512)
	d.YubicoKey = utils.FilterId(d.YubicoKey)
	d.OtpCode = utils.FilterId(d.OtpCode)
}

func authSessionPost(c *gin.Context) {
	data := &authSessionPostData{}

	req := &request.Request{
		Method: "POST",
		Path:   "/auth/session",
		Json:   data,
	}

	req.Do(c)
}

func authSessionDelete(c *gin.Context) {
	req := &request.Request{
		Method: "DELETE",
		Path:   "/auth/session",
	}

	req.Do(c)
}

func authStateGet(c *gin.Context) {
	req := &request.Request{
		Method: "GET",
		Path:   "/state",
	}

	req.Do(c)
}
