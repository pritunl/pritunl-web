package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/issuu/pritunl-web/request"
)

func linkGet(c *gin.Context) {
	var query map[string]string
	page := c.Query("page")
	if page != "" {
		query = map[string]string{
			"page": page,
		}
	}

	req := &request.Request{
		Method: "GET",
		Path:   "/link",
		Query:  query,
	}

	req.Do(c)
}

type linkPostData struct {
	Name   string `json:"name"`
	Type   string `json:"type"`
	Status string `json:"status"`
	Ipv6   bool   `json:"ipv6"`
}

func linkPost(c *gin.Context) {
	data := &linkPostData{}

	req := &request.Request{
		Method: "POST",
		Path:   "/link",
		Json:   data,
	}

	req.Do(c)
}

type linkPutData struct {
	Name   string `json:"name"`
	Status string `json:"status"`
	Key    bool   `json:"key"`
	Ipv6   bool   `json:"ipv6"`
}

type linkStatePutData struct {
	Version       string            `json:"version"`
	PublicAddress string            `json:"public_address"`
	LocalAddress  string            `json:"local_address"`
	Address6      string            `json:"address6"`
	Provider      string            `json:"provider"`
	Status        map[string]string `json:"status"`
	Errors        []string          `json:"errors"`
}

func linkPut(c *gin.Context) {
	var req *request.Request
	linkId := c.Params.ByName("link_id")

	if linkId == "state" {
		data := &linkStatePutData{}

		req = &request.Request{
			Method: "PUT",
			Path:   "/link/state",
			Json:   data,
		}
	} else {
		data := &linkPutData{}

		req = &request.Request{
			Method: "PUT",
			Path:   "/link/" + linkId,
			Json:   data,
		}
	}

	req.Do(c)
}

func linkDelete(c *gin.Context) {
	linkId := c.Params.ByName("link_id")

	req := &request.Request{
		Method: "DELETE",
		Path:   "/link/" + linkId,
	}

	req.Do(c)
}

func linkLocationGet(c *gin.Context) {
	linkId := c.Params.ByName("link_id")

	req := &request.Request{
		Method: "GET",
		Path:   "/link/" + linkId + "/location",
	}

	req.Do(c)
}

type linkLocationPostData struct {
	Name     string `json:"name"`
	LinkId   string `json:"link_id"`
	Location string `json:"location"`
}

func linkLocationPost(c *gin.Context) {
	linkId := c.Params.ByName("link_id")
	data := &linkLocationPostData{}

	req := &request.Request{
		Method: "POST",
		Path:   "/link/" + linkId + "/location",
		Json:   data,
	}

	req.Do(c)
}

type linkLocationPutData struct {
	Name     string `json:"name"`
	LinkId   string `json:"link_id"`
	Location string `json:"location"`
}

func linkLocationPut(c *gin.Context) {
	linkId := c.Params.ByName("link_id")
	locationId := c.Params.ByName("location_id")
	data := &linkLocationPutData{}

	req := &request.Request{
		Method: "PUT",
		Path:   "/link/" + linkId + "/location/" + locationId,
		Json:   data,
	}

	req.Do(c)
}

func linkLocationDelete(c *gin.Context) {
	linkId := c.Params.ByName("link_id")
	locationId := c.Params.ByName("location_id")

	req := &request.Request{
		Method: "DELETE",
		Path:   "/link/" + linkId + "/location/" + locationId,
	}

	req.Do(c)
}

type linkLocationRoutePostData struct {
	Network string `json:"network"`
}

func linkLocationRoutePost(c *gin.Context) {
	linkId := c.Params.ByName("link_id")
	locationId := c.Params.ByName("location_id")
	data := &linkLocationRoutePostData{}

	req := &request.Request{
		Method: "POST",
		Path:   "/link/" + linkId + "/location/" + locationId + "/route",
		Json:   data,
	}

	req.Do(c)
}

type linkLocationRoutePutData struct {
	Network string `json:"network"`
}

func linkLocationRoutePut(c *gin.Context) {
	linkId := c.Params.ByName("link_id")
	locationId := c.Params.ByName("location_id")
	routeId := c.Params.ByName("route_id")
	data := &linkLocationRoutePutData{}

	req := &request.Request{
		Method: "PUT",
		Path: "/link/" + linkId + "/location/" + locationId +
			"/route/" + routeId,
		Json: data,
	}

	req.Do(c)
}

func linkLocationRouteDelete(c *gin.Context) {
	linkId := c.Params.ByName("link_id")
	locationId := c.Params.ByName("location_id")
	routeId := c.Params.ByName("route_id")

	req := &request.Request{
		Method: "DELETE",
		Path: "/link/" + linkId + "/location/" + locationId +
			"/route/" + routeId,
	}

	req.Do(c)
}

func linkLocationHostUriGet(c *gin.Context) {
	linkId := c.Params.ByName("link_id")
	locationId := c.Params.ByName("location_id")
	hostId := c.Params.ByName("host_id")

	req := &request.Request{
		Method: "GET",
		Path: "/link/" + linkId + "/location/" + locationId +
			"/host/" + hostId + "/uri",
	}

	req.Do(c)
}

func linkLocationHostConfGet(c *gin.Context) {
	linkId := c.Params.ByName("link_id")
	locationId := c.Params.ByName("location_id")
	hostId := c.Params.ByName("host_id")

	req := &request.Request{
		Method: "GET",
		Path: "/link/" + linkId + "/location/" + locationId +
			"/host/" + hostId + "/conf",
	}

	req.Do(c)
}

type linkLocationHostPostData struct {
	Name          string `json:"name"`
	Timeout       int    `json:"timeout"`
	Priority      int    `json:"priority"`
	Static        bool   `json:"static"`
	PublicAddress string `json:"public_address"`
	LocalAddress  string `json:"local_address"`
	Address6      string `json:"address6"`
}

func linkLocationHostPost(c *gin.Context) {
	linkId := c.Params.ByName("link_id")
	locationId := c.Params.ByName("location_id")
	data := &linkLocationHostPostData{}

	req := &request.Request{
		Method: "POST",
		Path:   "/link/" + linkId + "/location/" + locationId + "/host",
		Json:   data,
	}

	req.Do(c)
}

type linkLocationHostPutData struct {
	Name          string `json:"name"`
	Timeout       int    `json:"timeout"`
	Priority      int    `json:"priority"`
	Static        bool   `json:"static"`
	PublicAddress string `json:"public_address"`
	LocalAddress  string `json:"local_address"`
	Address6      string `json:"address6"`
}

func linkLocationHostPut(c *gin.Context) {
	linkId := c.Params.ByName("link_id")
	locationId := c.Params.ByName("location_id")
	hostId := c.Params.ByName("host_id")
	data := &linkLocationHostPutData{}

	req := &request.Request{
		Method: "PUT",
		Path: "/link/" + linkId + "/location/" + locationId +
			"/host/" + hostId,
		Json: data,
	}

	req.Do(c)
}

func linkLocationHostDelete(c *gin.Context) {
	linkId := c.Params.ByName("link_id")
	locationId := c.Params.ByName("location_id")
	hostId := c.Params.ByName("host_id")

	req := &request.Request{
		Method: "DELETE",
		Path: "/link/" + linkId + "/location/" + locationId +
			"/host/" + hostId,
	}

	req.Do(c)
}

type linkLocationPeerPostData struct {
	PeerId string `json:"peer_id"`
}

func linkLocationPeerPost(c *gin.Context) {
	linkId := c.Params.ByName("link_id")
	locationId := c.Params.ByName("location_id")
	data := &linkLocationPeerPostData{}

	req := &request.Request{
		Method: "POST",
		Path: "/link/" + linkId + "/location/" + locationId +
			"/peer",
		Json: data,
	}

	req.Do(c)
}

func linkLocationPeerDelete(c *gin.Context) {
	linkId := c.Params.ByName("link_id")
	locationId := c.Params.ByName("location_id")
	peerId := c.Params.ByName("peer_id")

	req := &request.Request{
		Method: "DELETE",
		Path: "/link/" + linkId + "/location/" + locationId +
			"/peer/" + peerId,
	}

	req.Do(c)
}

type linkLocationTransitPostData struct {
	TransitId string `json:"transit_id"`
}

func linkLocationTransitPost(c *gin.Context) {
	linkId := c.Params.ByName("link_id")
	locationId := c.Params.ByName("location_id")
	data := &linkLocationTransitPostData{}

	req := &request.Request{
		Method: "POST",
		Path: "/link/" + linkId + "/location/" + locationId +
			"/transit",
		Json: data,
	}

	req.Do(c)
}

func linkLocationTransitDelete(c *gin.Context) {
	linkId := c.Params.ByName("link_id")
	locationId := c.Params.ByName("location_id")
	transitId := c.Params.ByName("transit_id")

	req := &request.Request{
		Method: "DELETE",
		Path: "/link/" + linkId + "/location/" + locationId +
			"/transit/" + transitId,
	}

	req.Do(c)
}
