package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/pritunl/pritunl-web/request"
	"github.com/pritunl/pritunl-web/utils"
)

func hostGet(c *gin.Context) {
	hostId := utils.FilterId(c.Params.ByName("host_id"))
	if hostId != "" {
		hostId = "/" + hostId
	}

	var query map[string]string
	page := utils.FilterId(c.Query("page"))
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
	Priority          int    `json:"priority"`
	InstanceId        string `json:"instance_id"`
}

func (d *hostPutData) Filter() {
	d.Name = utils.FilterStr(d.Name, 1024)
	d.PublicAddress = utils.FilterDomain(d.PublicAddress)
	d.PublicAddress6 = utils.FilterDomain(d.PublicAddress6)
	d.RoutedSubnet6 = utils.FilterStr(d.RoutedSubnet6, 1024)
	d.RoutedSubnet6Wg = utils.FilterStr(d.RoutedSubnet6Wg, 1024)
	d.LocalAddress = utils.FilterDomain(d.LocalAddress)
	d.LocalAddress6 = utils.FilterDomain(d.LocalAddress6)
	d.LinkAddress = utils.FilterDomain(d.LinkAddress)
	d.SyncAddress = utils.FilterDomain(d.SyncAddress)
	d.AvailabilityGroup = utils.FilterStr(d.AvailabilityGroup, 1024)
	d.InstanceId = utils.FilterStr(d.InstanceId, 1024)
}

func hostPut(c *gin.Context) {
	hostId := utils.FilterId(c.Params.ByName("host_id"))
	data := &hostPutData{}

	req := &request.Request{
		Method: "PUT",
		Path:   "/host/" + hostId,
		Json:   data,
	}

	req.Do(c)
}

func hostDelete(c *gin.Context) {
	hostId := utils.FilterId(c.Params.ByName("host_id"))

	req := &request.Request{
		Method: "DELETE",
		Path:   "/host/" + hostId,
	}

	req.Do(c)
}

func hostUsageGet(c *gin.Context) {
	hostId := utils.FilterId(c.Params.ByName("host_id"))
	period := utils.FilterId(c.Params.ByName("period"))

	req := &request.Request{
		Method: "GET",
		Path:   "/host/" + hostId + "/usage/" + period,
	}

	req.Do(c)
}
