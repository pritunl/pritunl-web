package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/pritunl/pritunl-web/request"
	"github.com/pritunl/pritunl-web/utils"
)

func deviceUnregisteredGet(c *gin.Context) {
	req := &request.Request{
		Method: "GET",
		Path:   "/device/unregistered",
	}

	req.Do(c)
}

type deviceRegisterPutData struct {
	Name   string `json:"name"`
	RegKey string `json:"reg_key"`
}

func (d *deviceRegisterPutData) Filter() {
	d.Name = utils.FilterStr(d.Name, 1024)
	d.RegKey = utils.FilterId(d.RegKey)
}

func deviceRegisterPut(c *gin.Context) {
	data := &deviceRegisterPutData{}

	orgId := utils.FilterId(c.Params.ByName("org_id"))
	userId := utils.FilterId(c.Params.ByName("user_id"))
	deviceId := utils.FilterId(c.Params.ByName("device_id"))

	req := &request.Request{
		Method: "PUT",
		Path:   "/device/register/" + orgId + "/" + userId + "/" + deviceId,
		Json:   data,
	}

	req.Do(c)
}

func deviceRegisterDelete(c *gin.Context) {
	orgId := utils.FilterId(c.Params.ByName("org_id"))
	userId := utils.FilterId(c.Params.ByName("user_id"))
	deviceId := utils.FilterId(c.Params.ByName("device_id"))

	req := &request.Request{
		Method: "DELETE",
		Path:   "/device/register/" + orgId + "/" + userId + "/" + deviceId,
	}

	req.Do(c)
}
