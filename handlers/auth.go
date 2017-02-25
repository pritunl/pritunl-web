package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/pritunl/pritunl-web/request"
)

type authSessionPostData struct {
	Username  string `json:"username"`
	Password  string `json:"password"`
	YubicoKey string `json:"yubico_key"`
	OtpCode   string `json:"otp_code"`
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
