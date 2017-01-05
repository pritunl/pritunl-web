package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/pritunl/pritunl-web/request"
)

func usersGet(c *gin.Context) {
	orgId := c.Params.ByName("org_id")

	var query map[string]string
	page := c.Query("page")
	if page != "" {
		query = map[string]string{
			"page": page,
		}
	}

	req := &request.Request{
		Method: "GET",
		Path:   "/user/" + orgId,
		Query:  query,
	}

	req.Do(c)
}

func userGet(c *gin.Context) {
	orgId := c.Params.ByName("org_id")
	userId := c.Params.ByName("user_id")

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
	Groups          []string                 `json:"groups"`
	Pin             string                   `json:"pin"`
	Disabled        bool                     `json:"disabled"`
	NetworkLinks    []string                 `json:"network_links"`
	BypassSecondary bool                     `json:"bypass_secondary"`
	ClientToClient  bool                     `json:"client_to_client"`
	DnsServers      []string                 `json:"dns_servers"`
	DnsSuffix       string                   `json:"dns_suffix"`
	PortForwarding  []userPortForwardingData `json:"port_forwarding"`
}

func userPost(c *gin.Context) {
	orgId := c.Params.ByName("org_id")
	data := &userPostData{}

	req := &request.Request{
		Method: "POST",
		Path:   "/user/" + orgId,
		Json:   data,
	}

	req.Do(c)
}

func userMultiPost(c *gin.Context) {
	orgId := c.Params.ByName("org_id")
	data := []*userPostData{}

	req := &request.Request{
		Method: "POST",
		Path:   "/user/" + orgId,
		Json:   data,
	}

	req.Do(c)
}

type userPutData struct {
	Name            string                   `json:"name"`
	Email           string                   `json:"email"`
	Groups          []string                 `json:"groups"`
	Pin             string                   `json:"pin"`
	Disabled        bool                     `json:"disabled"`
	NetworkLinks    []string                 `json:"network_links"`
	BypassSecondary bool                     `json:"bypass_secondary"`
	ClientToClient  bool                     `json:"client_to_client"`
	DnsServers      []string                 `json:"dns_servers"`
	DnsSuffix       string                   `json:"dns_suffix"`
	PortForwarding  []userPortForwardingData `json:"port_forwarding"`
	SendKeyEmail    bool                     `json:"send_key_email"`
}

func userPut(c *gin.Context) {
	orgId := c.Params.ByName("org_id")
	userId := c.Params.ByName("user_id")
	data := &userPutData{}

	req := &request.Request{
		Method: "PUT",
		Path:   "/user/" + orgId + "/" + userId,
		Json:   data,
	}

	req.Do(c)
}

func userDelete(c *gin.Context) {
	orgId := c.Params.ByName("org_id")
	userId := c.Params.ByName("user_id")

	req := &request.Request{
		Method: "DELETE",
		Path:   "/user/" + orgId + "/" + userId,
	}

	req.Do(c)
}

func userOtpSecretPut(c *gin.Context) {
	orgId := c.Params.ByName("org_id")
	userId := c.Params.ByName("user_id")

	req := &request.Request{
		Method: "PUT",
		Path:   "/user/" + orgId + "/" + userId + "/otp_secret",
	}

	req.Do(c)
}

func userAuditGet(c *gin.Context) {
	orgId := c.Params.ByName("org_id")
	userId := c.Params.ByName("user_id")

	req := &request.Request{
		Method: "GET",
		Path:   "/user/" + orgId + "/" + userId + "/audit",
	}

	req.Do(c)
}

type authUserPutData struct {
	Username     string   `json:"username"`
	NetworkLinks []string `json:"network_links"`
}

func authUserPost(c *gin.Context) {
	data := &authUserPutData{}

	req := &request.Request{
		Method: "POST",
		Path:   "/auth/user",
		Json:   data,
	}

	req.Do(c)
}
