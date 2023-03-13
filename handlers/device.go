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
	DeviceName   string `json:"device_name"`
	DeviceRegKey string `json:"device_reg_key"`
}

func deviceRegisterPut(c *gin.Context) {
	data := &deviceRegisterPutData{}

	orgId := utils.FilterStr(c.Params.ByName("org_id"), 64)
	userId := utils.FilterStr(c.Params.ByName("user_id"), 64)
	deviceId := utils.FilterStr(c.Params.ByName("device_id"), 64)

	req := &request.Request{
		Method: "PUT",
		Path:   "/device/register/" + orgId + "/" + userId + "/" + deviceId,
		Json:   data,
	}

	req.Do(c)
}

func deviceRegisterDelete(c *gin.Context) {
	orgId := utils.FilterStr(c.Params.ByName("org_id"), 64)
	userId := utils.FilterStr(c.Params.ByName("user_id"), 64)
	deviceId := utils.FilterStr(c.Params.ByName("device_id"), 64)

	req := &request.Request{
		Method: "DELETE",
		Path:   "/device/register/" + orgId + "/" + userId + "/" + deviceId,
	}

	req.Do(c)
}
