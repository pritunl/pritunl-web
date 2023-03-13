package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/pritunl/pritunl-web/request"
	"github.com/pritunl/pritunl-web/utils"
)

func hostGet(c *gin.Context) {
	hostId := utils.FilterStr(c.Params.ByName("host_id"), 128)
	if hostId != "" {
		hostId = "/" + hostId
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
		Path:   "/host" + hostId,
		Query:  query,
	}

	req.Do(c)
}

type hostPutData struct {
	Name              string `json:"name"`
	PublicAddress     string `json:"public_address"`
	PublicAddress6    string `json:"public_address6"`
	RoutedSubnet6     string `json:"routed_subnet6"`
	RoutedSubnet6Wg   string `json:"routed_subnet6_wg"`
	ProxyNdp          bool   `json:"proxy_ndp"`
	LocalAddress      string `json:"local_address"`
	LocalAddress6     string `json:"local_address6"`
	LinkAddress       string `json:"link_address"`
	SyncAddress       string `json:"sync_address"`
	AvailabilityGroup string `json:"availability_group"`
	InstanceId        string `json:"instance_id"`
}

func hostPut(c *gin.Context) {
	hostId := utils.FilterStr(c.Params.ByName("host_id"), 128)
	data := &hostPutData{}

	req := &request.Request{
		Method: "PUT",
		Path:   "/host/" + hostId,
		Json:   data,
	}

	req.Do(c)
}

func hostDelete(c *gin.Context) {
	hostId := utils.FilterStr(c.Params.ByName("host_id"), 128)

	req := &request.Request{
		Method: "DELETE",
		Path:   "/host/" + hostId,
	}

	req.Do(c)
}

func hostUsageGet(c *gin.Context) {
	hostId := utils.FilterStr(c.Params.ByName("host_id"), 128)
	period := utils.FilterStr(c.Params.ByName("period"), 128)

	req := &request.Request{
		Method: "GET",
		Path:   "/host/" + hostId + "/usage/" + period,
	}

	req.Do(c)
}
