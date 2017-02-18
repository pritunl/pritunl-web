package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/pritunl/pritunl-web/request"
)

func keyGet(c *gin.Context) {
	param1 := c.Params.ByName("param1")
	param2 := c.Params.ByName("param2")
	param3 := c.Params.ByName("param3")
	param4 := c.Params.ByName("param4")
	param5 := c.Params.ByName("param5")

	path := "/key/" + param1

	if param2 != "" {
		path += "/" + param2
		if param3 != "" {
			path += "/" + param3
			if param4 != "" {
				path += "/" + param4
				if param5 != "" {
					path += "/" + param5
				}
			}
		}
	}

	req := &request.Request{
		Method: "GET",
		Path:   path,
	}

	req.Do(c)
}

func keyOncGet(c *gin.Context) {
	param1 := c.Params.ByName("param1")
	param2 := c.Params.ByName("param2")

	path := "/key_onc/" + param1

	if param2 != "" {
		path += "/" + param2
	}

	req := &request.Request{
		Method: "GET",
		Path:   path,
	}

	req.Do(c)
}

type userKeyPinPutData struct {
	Pin        string `json:"pin"`
	CurrentPin string `json:"current_pin"`
}

func keyPinPut(c *gin.Context) {
	keyId := c.Params.ByName("key_id")
	data := &userKeyPinPutData{}

	req := &request.Request{
		Method: "PUT",
		Path:   "/key_pin/" + keyId,
		Json:   data,
	}

	req.Do(c)
}

func keyShortGet(c *gin.Context) {
	shortCode := c.Params.ByName("short_code")

	req := &request.Request{
		Method: "GET",
		Path:   "/k/" + shortCode,
	}

	req.Do(c)
}

func keyShortDelete(c *gin.Context) {
	shortCode := c.Params.ByName("short_code")

	req := &request.Request{
		Method: "DELETE",
		Path:   "/k/" + shortCode,
	}

	req.Do(c)
}

func keyApiShortGet(c *gin.Context) {
	shortCode := c.Params.ByName("short_code")

	req := &request.Request{
		Method: "GET",
		Path:   "/ku/" + shortCode,
	}

	req.Do(c)
}

type ssoAuthenticatePostData struct {
	Username string `json:"username"`
}

func ssoAuthenticatePost(c *gin.Context) {
	data := &ssoAuthenticatePostData{}

	req := &request.Request{
		Method: "POST",
		Path:   "/sso/authenticate",
		Json:   data,
	}

	req.Do(c)
}

func ssoRequestGet(c *gin.Context) {
	req := &request.Request{
		Method: "GET",
		Path:   "/sso/request",
	}

	req.Do(c)
}

func ssoCallbackGet(c *gin.Context) {
	req := &request.Request{
		Method:   "GET",
		Path:     "/sso/callback",
		RawQuery: c.Request.URL.RawQuery,
	}

	req.Do(c)
}

type ssoDuoPostData struct {
	Token    string `json:"token"`
	Passcode string `json:"passcode"`
}

func ssoDuoPost(c *gin.Context) {
	data := &ssoDuoPostData{}

	req := &request.Request{
		Method: "POST",
		Path:   "/sso/duo",
		Json:   data,
	}

	req.Do(c)
}

type ssoYubicoPostData struct {
	Token string `json:"token"`
	Key   string `json:"key"`
}

func ssoYubicoPost(c *gin.Context) {
	data := &ssoYubicoPostData{}

	req := &request.Request{
		Method: "POST",
		Path:   "/sso/yubico",
		Json:   data,
	}

	req.Do(c)
}
