package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/pritunl/pritunl-web/request"
	"github.com/pritunl/pritunl-web/utils"
)

func usersGet(c *gin.Context) {
	orgId := utils.FilterId(c.Params.ByName("org_id"))

	query := map[string]string{}

	page := utils.FilterId(c.Query("page"))
	if page != "" {
		query["page"] = page
	}

	lastActive := utils.FilterId(c.Query("last_active"))
	if lastActive != "" {
		query["last_active"] = lastActive
	}

	search := utils.FilterStr(c.Query("search"), 1024)
	if search != "" {
		query["search"] = search
	}

	limit := utils.FilterId(c.Query("limit"))
	if limit != "" {
		query["limit"] = limit
	}

	req := &request.Request{
		Method: "GET",
		Path:   "/user/" + orgId,
		Query:  query,
	}

	req.Do(c)
}

func userGet(c *gin.Context) {
	orgId := utils.FilterId(c.Params.ByName("org_id"))
	userId := utils.FilterId(c.Params.ByName("user_id"))

	req := &request.Request{
		Method: "GET",
		Path:   "/user/" + orgId + "/" + userId,
	}

	req.Do(c)
}

type userPortForwardingData struct {
	Protocol string `json:"protocol"`
	Port     string `json:"port"`
	Dport    string `json:"dport"`
}

func (d *userPortForwardingData) Filter() {
	d.Protocol = utils.FilterId(d.Protocol)
	d.Port = utils.FilterId(d.Port)
	d.Dport = utils.FilterId(d.Dport)
}

type userPostData struct {
	Name            string                   `json:"name"`
	Email           string                   `json:"email"`
	AuthType        string                   `json:"auth_type"`
	YubicoId        string                   `json:"yubico_id"`
	Groups          []string                 `json:"groups"`
	Pin             string                   `json:"pin"`
	Disabled        bool                     `json:"disabled"`
	NetworkLinks    []string                 `json:"network_links"`
	BypassSecondary bool                     `json:"bypass_secondary"`
	ClientToClient  bool                     `json:"client_to_client"`
	MacAddresses    []string                 `json:"mac_addresses"`
	DnsServers      []string                 `json:"dns_servers"`
	DnsSuffix       string                   `json:"dns_suffix"`
	PortForwarding  []userPortForwardingData `json:"port_forwarding"`
}

func (d *userPostData) Filter() {
	d.Name = utils.FilterStr(d.Name, 1024)
	d.Email = utils.FilterStr(d.Email, 1024)
	d.AuthType = utils.FilterId(d.AuthType)
	d.YubicoId = utils.FilterId(d.YubicoId)
	d.Pin = utils.FilterStr(d.Pin, 1024)
	d.DnsSuffix = utils.FilterStr(d.DnsSuffix, 1024)

	for i, group := range d.Groups {
		d.Groups[i] = utils.FilterStr(group, 1024)
	}

	for i, networkLink := range d.NetworkLinks {
		d.NetworkLinks[i] = utils.FilterStr(networkLink, 1024)
	}

	for i, macAddress := range d.MacAddresses {
		d.MacAddresses[i] = utils.FilterId(macAddress)
	}

	for i, dnsServer := range d.DnsServers {
		d.DnsServers[i] = utils.FilterDomain(dnsServer)
	}

	for i := range d.PortForwarding {
		d.PortForwarding[i].Filter()
	}
}

type userMultiPostData []*userPostData

func (d *userMultiPostData) Filter() {
	for _, user := range *d {
		if user != nil {
			user.Filter()
		}
	}
}

func userPost(c *gin.Context) {
	orgId := utils.FilterId(c.Params.ByName("org_id"))
	data := &userPostData{}

	req := &request.Request{
		Method: "POST",
		Path:   "/user/" + orgId,
		Json:   data,
	}

	req.Do(c)
}

func userMultiPost(c *gin.Context) {
	orgId := utils.FilterId(c.Params.ByName("org_id"))
	data := userMultiPostData{}

	req := &request.Request{
		Method: "POST",
		Path:   "/user/" + orgId + "/multi",
		Json:   &data,
	}

	req.Do(c)
}

type userPutData struct {
	Name            string                   `json:"name"`
	Email           string                   `json:"email"`
	AuthType        string                   `json:"auth_type"`
	YubicoId        string                   `json:"yubico_id"`
	Groups          []string                 `json:"groups"`
	Pin             interface{}              `json:"pin"`
	Disabled        bool                     `json:"disabled"`
	NetworkLinks    []string                 `json:"network_links"`
	BypassSecondary bool                     `json:"bypass_secondary"`
	ClientToClient  bool                     `json:"client_to_client"`
	MacAddresses    []string                 `json:"mac_addresses"`
	DnsServers      []string                 `json:"dns_servers"`
	DnsSuffix       string                   `json:"dns_suffix"`
	PortForwarding  []userPortForwardingData `json:"port_forwarding"`
	SendKeyEmail    bool                     `json:"send_key_email"`
}

func (d *userPutData) Filter() {
	d.Name = utils.FilterStr(d.Name, 1024)
	d.Email = utils.FilterStr(d.Email, 1024)
	d.AuthType = utils.FilterId(d.AuthType)
	d.YubicoId = utils.FilterId(d.YubicoId)
	d.DnsSuffix = utils.FilterStr(d.DnsSuffix, 1024)

	switch pin := d.Pin.(type) {
	case string:
		d.Pin = utils.FilterStr(pin, 1024)
	case bool:
	default:
		d.Pin = nil
	}

	for i, group := range d.Groups {
		d.Groups[i] = utils.FilterStr(group, 1024)
	}

	for i, networkLink := range d.NetworkLinks {
		d.NetworkLinks[i] = utils.FilterStr(networkLink, 1024)
	}

	for i, macAddress := range d.MacAddresses {
		d.MacAddresses[i] = utils.FilterId(macAddress)
	}

	for i, dnsServer := range d.DnsServers {
		d.DnsServers[i] = utils.FilterDomain(dnsServer)
	}

	for i := range d.PortForwarding {
		d.PortForwarding[i].Filter()
	}
}

func userPut(c *gin.Context) {
	orgId := utils.FilterId(c.Params.ByName("org_id"))
	userId := utils.FilterId(c.Params.ByName("user_id"))
	data := &userPutData{}

	req := &request.Request{
		Method: "PUT",
		Path:   "/user/" + orgId + "/" + userId,
		Json:   data,
	}

	req.Do(c)
}

func userDelete(c *gin.Context) {
	orgId := utils.FilterId(c.Params.ByName("org_id"))
	userId := utils.FilterId(c.Params.ByName("user_id"))

	req := &request.Request{
		Method: "DELETE",
		Path:   "/user/" + orgId + "/" + userId,
	}

	req.Do(c)
}

func userOtpSecretPut(c *gin.Context) {
	orgId := utils.FilterId(c.Params.ByName("org_id"))
	userId := utils.FilterId(c.Params.ByName("user_id"))

	req := &request.Request{
		Method: "PUT",
		Path:   "/user/" + orgId + "/" + userId + "/otp_secret",
	}

	req.Do(c)
}

func userAuditGet(c *gin.Context) {
	orgId := utils.FilterId(c.Params.ByName("org_id"))
	userId := utils.FilterId(c.Params.ByName("user_id"))

	req := &request.Request{
		Method: "GET",
		Path:   "/user/" + orgId + "/" + userId + "/audit",
	}

	req.Do(c)
}

type userDevicePutData struct {
	Name   string `json:"name"`
	RegKey string `json:"reg_key"`
}

func (d *userDevicePutData) Filter() {
	d.Name = utils.FilterStr(d.Name, 1024)
	d.RegKey = utils.FilterId(d.RegKey)
}

func userDevicePut(c *gin.Context) {
	data := &userDevicePutData{}

	orgId := utils.FilterId(c.Params.ByName("org_id"))
	userId := utils.FilterId(c.Params.ByName("user_id"))
	deviceId := utils.FilterId(c.Params.ByName("device_id"))

	req := &request.Request{
		Method: "PUT",
		Path:   "/user/" + orgId + "/" + userId + "/device/" + deviceId,
		Json:   data,
	}

	req.Do(c)
}

func userDeviceDelete(c *gin.Context) {
	orgId := utils.FilterId(c.Params.ByName("org_id"))
	userId := utils.FilterId(c.Params.ByName("user_id"))
	deviceId := utils.FilterId(c.Params.ByName("device_id"))

	req := &request.Request{
		Method: "DELETE",
		Path:   "/user/" + orgId + "/" + userId + "/device/" + deviceId,
	}

	req.Do(c)
}
