package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/pritunl/pritunl-web/request"
	"github.com/pritunl/pritunl-web/utils"
)

func usersGet(c *gin.Context) {
	orgId := utils.FilterStr(c.Params.ByName("org_id"), 128)

	query := map[string]string{}

	page := c.Query("page")
	if page != "" {
		query["page"] = page
	}

	lastActive := c.Query("last_active")
	if lastActive != "" {
		query["last_active"] = lastActive
	}

	search := c.Query("search")
	if search != "" {
		query["search"] = search
	}

	limit := c.Query("limit")
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
	orgId := utils.FilterStr(c.Params.ByName("org_id"), 128)
	userId := utils.FilterStr(c.Params.ByName("user_id"), 128)

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

func userPost(c *gin.Context) {
	orgId := utils.FilterStr(c.Params.ByName("org_id"), 128)
	data := &userPostData{}

	req := &request.Request{
		Method: "POST",
		Path:   "/user/" + orgId,
		Json:   data,
	}

	req.Do(c)
}

func userMultiPost(c *gin.Context) {
	orgId := utils.FilterStr(c.Params.ByName("org_id"), 128)
	data := []*userPostData{}

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

func userPut(c *gin.Context) {
	orgId := utils.FilterStr(c.Params.ByName("org_id"), 128)
	userId := utils.FilterStr(c.Params.ByName("user_id"), 128)
	data := &userPutData{}

	switch data.Pin.(type) {
	case string, bool:
	default:
		data.Pin = nil
	}

	req := &request.Request{
		Method: "PUT",
		Path:   "/user/" + orgId + "/" + userId,
		Json:   data,
	}

	req.Do(c)
}

func userDelete(c *gin.Context) {
	orgId := utils.FilterStr(c.Params.ByName("org_id"), 128)
	userId := utils.FilterStr(c.Params.ByName("user_id"), 128)

	req := &request.Request{
		Method: "DELETE",
		Path:   "/user/" + orgId + "/" + userId,
	}

	req.Do(c)
}

func userOtpSecretPut(c *gin.Context) {
	orgId := utils.FilterStr(c.Params.ByName("org_id"), 128)
	userId := utils.FilterStr(c.Params.ByName("user_id"), 128)

	req := &request.Request{
		Method: "PUT",
		Path:   "/user/" + orgId + "/" + userId + "/otp_secret",
	}

	req.Do(c)
}

func userAuditGet(c *gin.Context) {
	orgId := utils.FilterStr(c.Params.ByName("org_id"), 128)
	userId := utils.FilterStr(c.Params.ByName("user_id"), 128)

	req := &request.Request{
		Method: "GET",
		Path:   "/user/" + orgId + "/" + userId + "/audit",
	}

	req.Do(c)
}
