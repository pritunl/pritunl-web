package handlers

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/pritunl/pritunl-web/request"
	"github.com/pritunl/pritunl-web/utils"
)

func keyGet(c *gin.Context) {
	param1 := utils.FilterStr(c.Params.ByName("param1"), 512)
	param2 := utils.FilterStr(c.Params.ByName("param2"), 512)
	param3 := utils.FilterStr(c.Params.ByName("param3"), 512)
	param4 := utils.FilterStr(c.Params.ByName("param4"), 512)
	param5 := utils.FilterStr(c.Params.ByName("param5"), 512)

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

	if param1 == "sync" {
		vInt, _ := strconv.Atoi(c.Query("ver"))
		if vInt != 0 {
			req.Query = map[string]string{
				"ver": strconv.Itoa(vInt),
			}
		}
	}

	if param1 == "request" || param1 == "callback" {
		req.RawQuery = utils.FilterOpen(c.Request.URL.RawQuery)
	}

	req.Do(c)
}

type userKeyPinPutData struct {
	Pin        string `json:"pin"`
	CurrentPin string `json:"current_pin"`
}

func (d *userKeyPinPutData) Filter() {
	d.Pin = utils.FilterStr(d.Pin, 1024)
	d.CurrentPin = utils.FilterStr(d.CurrentPin, 1024)
}

func keyPinPut(c *gin.Context) {
	keyId := utils.FilterId(c.Params.ByName("key_id"))
	data := &userKeyPinPutData{}

	req := &request.Request{
		Method: "PUT",
		Path:   "/key_pin/" + keyId,
		Json:   data,
	}

	req.Do(c)
}

func keyShortGet(c *gin.Context) {
	shortCode := utils.FilterId(c.Params.ByName("short_code"))

	req := &request.Request{
		Method: "GET",
		Path:   "/k/" + shortCode,
	}

	req.Do(c)
}

func keyShortDelete(c *gin.Context) {
	shortCode := utils.FilterId(c.Params.ByName("short_code"))

	req := &request.Request{
		Method: "DELETE",
		Path:   "/k/" + shortCode,
	}

	req.Do(c)
}

func keyApiShortGet(c *gin.Context) {
	shortCode := utils.FilterId(c.Params.ByName("short_code"))

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

func (d *keyWgPutPostData) Filter() {
	d.Data = utils.FilterBase64(d.Data)
	d.Nonce = utils.FilterBase64(d.Nonce)
	d.PublicKey = utils.FilterBase64(d.PublicKey)
	d.Signature = utils.FilterBase64(d.Signature)
	d.DeviceSignature = utils.FilterBase64(d.DeviceSignature)
}

func keyWgPut(c *gin.Context) {
	orgId := utils.FilterId(c.Params.ByName("org_id"))
	userId := utils.FilterId(c.Params.ByName("user_id"))
	serverId := utils.FilterId(c.Params.ByName("server_id"))
	data := &keyWgPutPostData{}

	req := &request.Request{
		Method: "PUT",
		Path:   "/key/wg/" + orgId + "/" + userId + "/" + serverId,
		Json:   data,
	}

	req.Do(c)
}

func keyWgPost(c *gin.Context) {
	orgId := utils.FilterId(c.Params.ByName("org_id"))
	userId := utils.FilterId(c.Params.ByName("user_id"))
	serverId := utils.FilterId(c.Params.ByName("server_id"))
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

func (d *keyOvpnPostData) Filter() {
	d.Data = utils.FilterBase64(d.Data)
	d.Nonce = utils.FilterBase64(d.Nonce)
	d.PublicKey = utils.FilterBase64(d.PublicKey)
	d.Signature = utils.FilterBase64(d.Signature)
	d.DeviceSignature = utils.FilterBase64(d.DeviceSignature)
}

func keyOvpnPost(c *gin.Context) {
	orgId := utils.FilterId(c.Params.ByName("org_id"))
	userId := utils.FilterId(c.Params.ByName("user_id"))
	serverId := utils.FilterId(c.Params.ByName("server_id"))
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

func (d *keyOvpnWaitPostData) Filter() {
	d.Data = utils.FilterBase64(d.Data)
	d.Nonce = utils.FilterBase64(d.Nonce)
	d.PublicKey = utils.FilterBase64(d.PublicKey)
	d.Signature = utils.FilterBase64(d.Signature)
	d.DeviceSignature = utils.FilterBase64(d.DeviceSignature)
}

func keyOvpnWaitPost(c *gin.Context) {
	orgId := utils.FilterId(c.Params.ByName("org_id"))
	userId := utils.FilterId(c.Params.ByName("user_id"))
	serverId := utils.FilterId(c.Params.ByName("server_id"))
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

func (d *keyWgWaitPostData) Filter() {
	d.Data = utils.FilterBase64(d.Data)
	d.Nonce = utils.FilterBase64(d.Nonce)
	d.PublicKey = utils.FilterBase64(d.PublicKey)
	d.Signature = utils.FilterBase64(d.Signature)
	d.DeviceSignature = utils.FilterBase64(d.DeviceSignature)
}

func keyWgWaitPost(c *gin.Context) {
	orgId := utils.FilterId(c.Params.ByName("org_id"))
	userId := utils.FilterId(c.Params.ByName("user_id"))
	serverId := utils.FilterId(c.Params.ByName("server_id"))
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

func (d *ssoAuthenticatePostData) Filter() {
	d.Username = utils.FilterStr(d.Username, 1024)
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
		RawQuery: utils.FilterOpen(c.Request.URL.RawQuery),
	}

	req.Do(c)
}

type ssoDuoPostData struct {
	Token    string `json:"token"`
	Passcode string `json:"passcode"`
}

func (d *ssoDuoPostData) Filter() {
	d.Token = utils.FilterId(d.Token)
	d.Passcode = utils.FilterId(d.Passcode)
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

func (d *ssoYubicoPostData) Filter() {
	d.Token = utils.FilterId(d.Token)
	d.Key = utils.FilterId(d.Key)
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

func (d *keyDuoPostData) Filter() {
	d.Token = utils.FilterId(d.Token)
	d.Passcode = utils.FilterId(d.Passcode)
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

func (d *keyYubicoPostData) Filter() {
	d.Token = utils.FilterId(d.Token)
	d.Key = utils.FilterId(d.Key)
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
