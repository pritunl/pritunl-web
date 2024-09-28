package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/pritunl/pritunl-web/request"
	"github.com/pritunl/pritunl-web/utils"
)

func serverGet(c *gin.Context) {
	serverId := utils.FilterStr(c.Params.ByName("server_id"), 128)
	if serverId != "" {
		serverId = "/" + serverId
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
		Path:   "/server" + serverId,
		Query:  query,
	}

	req.Do(c)
}

type serverPostPutData struct {
	Name             string      `json:"name"`
	Network          string      `json:"network"`
	NetworkWg        string      `json:"network_wg"`
	NetworkMode      string      `json:"network_mode"`
	NetworkStart     string      `json:"network_start"`
	NetworkEnd       string      `json:"network_end"`
	RestrictRoutes   bool        `json:"restrict_routes"`
	Wg               bool        `json:"wg"`
	Ipv6             bool        `json:"ipv6"`
	Ipv6Firewall     bool        `json:"ipv6_firewall"`
	DynamicFirewall  bool        `json:"dynamic_firewall"`
	GeoSort          bool        `json:"geo_sort"`
	ForceConnect     bool        `json:"force_connect"`
	DeviceAuth       bool        `json:"device_auth"`
	BindAddress      string      `json:"bind_address"`
	Protocol         string      `json:"protocol"`
	Port             int         `json:"port"`
	PortWg           int         `json:"port_wg"`
	DhParamBits      int         `json:"dh_param_bits"`
	Groups           []string    `json:"groups"`
	MultiDevice      bool        `json:"multi_device"`
	DnsServers       []string    `json:"dns_servers"`
	SearchDomain     string      `json:"search_domain"`
	InterClient      bool        `json:"inter_client"`
	PingInterval     int         `json:"ping_interval"`
	PingTimeout      int         `json:"ping_timeout"`
	LinkPingInterval int         `json:"link_ping_interval"`
	LinkPingTimeout  int         `json:"link_ping_timeout"`
	InactiveTimeout  int         `json:"inactive_timeout"`
	SessionTimeout   int         `json:"session_timeout"`
	AllowedDevices   string      `json:"allowed_devices"`
	MaxClients       int         `json:"max_clients"`
	MaxDevices       int         `json:"max_devices"`
	ReplicaCount     int         `json:"replica_count"`
	Vxlan            bool        `json:"vxlan"`
	DnsMapping       bool        `json:"dns_mapping"`
	RouteDns         bool        `json:"route_dns"`
	Debug            bool        `json:"debug"`
	SsoAuth          bool        `json:"sso_auth"`
	OtpAuth          bool        `json:"otp_auth"`
	LzoCompression   bool        `json:"lzo_compression"`
	Cipher           string      `json:"cipher"`
	Hash             string      `json:"hash"`
	BlockOutsideDns  bool        `json:"block_outside_dns"`
	JumboFrames      bool        `json:"jumbo_frames"`
	PreConnectMsg    string      `json:"pre_connect_msg"`
	Policy           string      `json:"policy"`
	MssFix           interface{} `json:"mss_fix"`
	Multihome        bool        `json:"multihome"`
}

func serverPost(c *gin.Context) {
	data := &serverPostPutData{}

	switch data.MssFix.(type) {
	case string, int:
	default:
		data.MssFix = nil
	}

	req := &request.Request{
		Method: "POST",
		Path:   "/server",
		Json:   data,
	}

	req.Do(c)
}

func serverPut(c *gin.Context) {
	serverId := utils.FilterStr(c.Params.ByName("server_id"), 128)
	data := &serverPostPutData{}

	switch data.MssFix.(type) {
	case string, int:
	default:
		data.MssFix = nil
	}

	req := &request.Request{
		Method: "PUT",
		Path:   "/server/" + serverId,
		Json:   data,
	}

	req.Do(c)
}

func serverDelete(c *gin.Context) {
	serverId := utils.FilterStr(c.Params.ByName("server_id"), 128)

	req := &request.Request{
		Method: "DELETE",
		Path:   "/server/" + serverId,
	}

	req.Do(c)
}

func serverOrgGet(c *gin.Context) {
	serverId := utils.FilterStr(c.Params.ByName("server_id"), 128)

	req := &request.Request{
		Method: "GET",
		Path:   "/server/" + serverId + "/organization",
	}

	req.Do(c)
}

func serverOrgPut(c *gin.Context) {
	serverId := utils.FilterStr(c.Params.ByName("server_id"), 128)
	orgId := utils.FilterStr(c.Params.ByName("org_id"), 128)

	req := &request.Request{
		Method: "PUT",
		Path:   "/server/" + serverId + "/organization/" + orgId,
	}

	req.Do(c)
}

func serverOrgDelete(c *gin.Context) {
	serverId := utils.FilterStr(c.Params.ByName("server_id"), 128)
	orgId := utils.FilterStr(c.Params.ByName("org_id"), 128)

	req := &request.Request{
		Method: "DELETE",
		Path:   "/server/" + serverId + "/organization/" + orgId,
	}

	req.Do(c)
}

func serverRouteGet(c *gin.Context) {
	serverId := utils.FilterStr(c.Params.ByName("server_id"), 128)

	req := &request.Request{
		Method: "GET",
		Path:   "/server/" + serverId + "/route",
	}

	req.Do(c)
}

type serverRoutePostPutData struct {
	Network      string `json:"network"`
	Comment      string `json:"comment"`
	Metric       int    `json:"metric"`
	Nat          bool   `json:"nat"`
	NatInterface string `json:"nat_interface"`
	NatNetmap    string `json:"nat_netmap"`
	Advertise    bool   `json:"advertise"`
	VpcRegion    string `json:"vpc_region"`
	VpcId        string `json:"vpc_id"`
	NetGateway   bool   `json:"net_gateway"`
}

func serverRoutePost(c *gin.Context) {
	serverId := utils.FilterStr(c.Params.ByName("server_id"), 128)
	data := &serverRoutePostPutData{}

	req := &request.Request{
		Method: "POST",
		Path:   "/server/" + serverId + "/route",
		Json:   data,
	}

	req.Do(c)
}

func serverRoutesPost(c *gin.Context) {
	serverId := utils.FilterStr(c.Params.ByName("server_id"), 128)
	data := []*serverRoutePostPutData{}

	req := &request.Request{
		Method: "POST",
		Path:   "/server/" + serverId + "/routes",
		Json:   &data,
	}

	req.Do(c)
}

func serverRoutePut(c *gin.Context) {
	serverId := utils.FilterStr(c.Params.ByName("server_id"), 128)
	routeNet := utils.FilterStr(c.Params.ByName("route_net"), 128)
	data := &serverRoutePostPutData{}

	req := &request.Request{
		Method: "PUT",
		Path:   "/server/" + serverId + "/route/" + routeNet,
		Json:   data,
	}

	req.Do(c)
}

func serverRouteDelete(c *gin.Context) {
	serverId := utils.FilterStr(c.Params.ByName("server_id"), 128)
	routeNet := utils.FilterStr(c.Params.ByName("route_net"), 128)

	req := &request.Request{
		Method: "DELETE",
		Path:   "/server/" + serverId + "/route/" + routeNet,
	}

	req.Do(c)
}

func serverHostGet(c *gin.Context) {
	serverId := utils.FilterStr(c.Params.ByName("server_id"), 128)

	req := &request.Request{
		Method: "GET",
		Path:   "/server/" + serverId + "/host",
	}

	req.Do(c)
}

func serverHostPut(c *gin.Context) {
	serverId := utils.FilterStr(c.Params.ByName("server_id"), 128)
	hostId := utils.FilterStr(c.Params.ByName("host_id"), 128)

	req := &request.Request{
		Method: "PUT",
		Path:   "/server/" + serverId + "/host/" + hostId,
	}

	req.Do(c)
}

func serverHostDelete(c *gin.Context) {
	serverId := utils.FilterStr(c.Params.ByName("server_id"), 128)
	hostId := utils.FilterStr(c.Params.ByName("host_id"), 128)

	req := &request.Request{
		Method: "DELETE",
		Path:   "/server/" + serverId + "/host/" + hostId,
	}

	req.Do(c)
}

func serverLinkGet(c *gin.Context) {
	serverId := utils.FilterStr(c.Params.ByName("server_id"), 128)

	req := &request.Request{
		Method: "GET",
		Path:   "/server/" + serverId + "/link",
	}

	req.Do(c)
}

type serverLinkPutData struct {
	UseLocalAddress bool `json:"use_local_address"`
}

func serverLinkPut(c *gin.Context) {
	serverId := utils.FilterStr(c.Params.ByName("server_id"), 128)
	linkId := utils.FilterStr(c.Params.ByName("link_id"), 128)
	data := &serverLinkPutData{}

	req := &request.Request{
		Method: "PUT",
		Path:   "/server/" + serverId + "/link/" + linkId,
		Json:   data,
	}

	req.Do(c)
}

func serverLinkDelete(c *gin.Context) {
	serverId := utils.FilterStr(c.Params.ByName("server_id"), 128)
	linkId := utils.FilterStr(c.Params.ByName("link_id"), 128)

	req := &request.Request{
		Method: "DELETE",
		Path:   "/server/" + serverId + "/link/" + linkId,
	}

	req.Do(c)
}

func serverOperationPut(c *gin.Context) {
	serverId := utils.FilterStr(c.Params.ByName("server_id"), 128)
	operation := utils.FilterStr(c.Params.ByName("operation"), 128)

	req := &request.Request{
		Method: "PUT",
		Path:   "/server/" + serverId + "/operation/" + operation,
	}

	req.Do(c)
}

func serverOutputGet(c *gin.Context) {
	serverId := utils.FilterStr(c.Params.ByName("server_id"), 128)

	req := &request.Request{
		Method: "GET",
		Path:   "/server/" + serverId + "/output",
	}

	req.Do(c)
}

func serverOutputDelete(c *gin.Context) {
	serverId := utils.FilterStr(c.Params.ByName("server_id"), 128)

	req := &request.Request{
		Method: "DELETE",
		Path:   "/server/" + serverId + "/output",
	}

	req.Do(c)
}

func serverLinkOutputGet(c *gin.Context) {
	serverId := utils.FilterStr(c.Params.ByName("server_id"), 128)

	req := &request.Request{
		Method: "GET",
		Path:   "/server/" + serverId + "/link_output",
	}

	req.Do(c)
}

func serverLinkOutputDelete(c *gin.Context) {
	serverId := utils.FilterStr(c.Params.ByName("server_id"), 128)

	req := &request.Request{
		Method: "DELETE",
		Path:   "/server/" + serverId + "/link_output",
	}

	req.Do(c)
}

func serverBandwidthGet(c *gin.Context) {
	serverId := utils.FilterStr(c.Params.ByName("server_id"), 128)
	period := utils.FilterStr(c.Params.ByName("period"), 128)

	req := &request.Request{
		Method: "GET",
		Path:   "/server/" + serverId + "/bandwidth/" + period,
	}

	req.Do(c)
}
