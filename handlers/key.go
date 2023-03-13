package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/pritunl/pritunl-web/request"
	"github.com/pritunl/pritunl-web/utils"
)

func keyGet(c *gin.Context) {
	param1 := utils.FilterStr(c.Params.ByName("param1"), 128)
	param2 := utils.FilterStr(c.Params.ByName("param2"), 128)
	param3 := utils.FilterStr(c.Params.ByName("param3"), 128)
	param4 := utils.FilterStr(c.Params.ByName("param4"), 128)
	param5 := utils.FilterStr(c.Params.ByName("param5"), 128)

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

	if param1 == "request" || param1 == "callback" {
		req.RawQuery = c.Request.URL.RawQuery
	}

	req.Do(c)
}

func keyOncGet(c *gin.Context) {
	param1 := utils.FilterStr(c.Params.ByName("param1"), 128)
	param2 := utils.FilterStr(c.Params.ByName("param2"), 128)

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
	keyId := utils.FilterStr(c.Params.ByName("key_id"), 128)
	data := &userKeyPinPutData{}

	req := &request.Request{
		Method: "PUT",
		Path:   "/key_pin/" + keyId,
		Json:   data,
	}

	req.Do(c)
}

func keyShortGet(c *gin.Context) {
	shortCode := utils.FilterStr(c.Params.ByName("short_code"), 128)

	req := &request.Request{
		Method: "GET",
		Path:   "/k/" + shortCode,
	}

	req.Do(c)
}

func keyShortDelete(c *gin.Context) {
	shortCode := utils.FilterStr(c.Params.ByName("short_code"), 128)

	req := &request.Request{
		Method: "DELETE",
		Path:   "/k/" + shortCode,
	}

	req.Do(c)
}

func keyApiShortGet(c *gin.Context) {
	shortCode := utils.FilterStr(c.Params.ByName("short_code"), 128)

	req := &request.Request{
		Method: "GET",
		Path:   "/ku/" + shortCode,
	}

	req.Do(c)
}

type keyWgPutPostData struct {
	Data            string `json:"data"`
	Nonce           string `json:"nonce"`
	PublicKey       string `json:"public_key"`
	Signature       string `json:"signature"`
	DeviceSignature string `json:"device_signature"`
}

func keyWgPut(c *gin.Context) {
	orgId := utils.FilterStr(c.Params.ByName("org_id"), 128)
	userId := utils.FilterStr(c.Params.ByName("user_id"), 128)
	serverId := utils.FilterStr(c.Params.ByName("server_id"), 128)
	data := &keyWgPutPostData{}

	req := &request.Request{
		Method: "PUT",
		Path:   "/key/wg/" + orgId + "/" + userId + "/" + serverId,
		Json:   data,
	}

	req.Do(c)
}

func keyWgPost(c *gin.Context) {
	orgId := utils.FilterStr(c.Params.ByName("org_id"), 128)
	userId := utils.FilterStr(c.Params.ByName("user_id"), 128)
	serverId := utils.FilterStr(c.Params.ByName("server_id"), 128)
	data := &keyWgPutPostData{}

	req := &request.Request{
		Method: "POST",
		Path:   "/key/wg/" + orgId + "/" + userId + "/" + serverId,
		Json:   data,
	}

	req.Do(c)
}

type keyOvpnPostData struct {
	Data            string `json:"data"`
	Nonce           string `json:"nonce"`
	PublicKey       string `json:"public_key"`
	Signature       string `json:"signature"`
	DeviceSignature string `json:"device_signature"`
}

func keyOvpnPost(c *gin.Context) {
	orgId := utils.FilterStr(c.Params.ByName("org_id"), 128)
	userId := utils.FilterStr(c.Params.ByName("user_id"), 128)
	serverId := utils.FilterStr(c.Params.ByName("server_id"), 128)
	data := &keyOvpnPostData{}

	req := &request.Request{
		Method: "POST",
		Path:   "/key/ovpn/" + orgId + "/" + userId + "/" + serverId,
		Json:   data,
	}

	req.Do(c)
}

type keyOvpnWaitPostData struct {
	Data            string `json:"data"`
	Nonce           string `json:"nonce"`
	PublicKey       string `json:"public_key"`
	Signature       string `json:"signature"`
	DeviceSignature string `json:"device_signature"`
}

func keyOvpnWaitPost(c *gin.Context) {
	orgId := utils.FilterStr(c.Params.ByName("org_id"), 128)
	userId := utils.FilterStr(c.Params.ByName("user_id"), 128)
	serverId := utils.FilterStr(c.Params.ByName("server_id"), 128)
	data := &keyOvpnWaitPostData{}

	req := &request.Request{
		Method: "POST",
		Path:   "/key/ovpn_wait/" + orgId + "/" + userId + "/" + serverId,
		Json:   data,
	}

	req.Do(c)
}

type keyWgWaitPostData struct {
	Data            string `json:"data"`
	Nonce           string `json:"nonce"`
	PublicKey       string `json:"public_key"`
	Signature       string `json:"signature"`
	DeviceSignature string `json:"device_signature"`
}

func keyWgWaitPost(c *gin.Context) {
	orgId := utils.FilterStr(c.Params.ByName("org_id"), 128)
	userId := utils.FilterStr(c.Params.ByName("user_id"), 128)
	serverId := utils.FilterStr(c.Params.ByName("server_id"), 128)
	data := &keyWgWaitPostData{}

	req := &request.Request{
		Method: "POST",
		Path:   "/key/wg_wait/" + orgId + "/" + userId + "/" + serverId,
		Json:   data,
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

type keyDuoPostData struct {
	Token    string `json:"token"`
	Passcode string `json:"passcode"`
}

func keyDuoPost(c *gin.Context) {
	data := &keyDuoPostData{}

	req := &request.Request{
		Method: "POST",
		Path:   "/key/duo",
		Json:   data,
	}

	req.Do(c)
}

type keyYubicoPostData struct {
	Token string `json:"token"`
	Key   string `json:"key"`
}

func keyYubicoPost(c *gin.Context) {
	data := &keyYubicoPostData{}

	req := &request.Request{
		Method: "POST",
		Path:   "/key/yubico",
		Json:   data,
	}

	req.Do(c)
}
