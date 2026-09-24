package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/pritunl/pritunl-web/request"
	"github.com/pritunl/pritunl-web/utils"
)

func linkGet(c *gin.Context) {
	var query map[string]string
	page := utils.FilterId(c.Query("page"))
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
	Name           string `json:"name"`
	Type           string `json:"type"`
	Status         string `json:"status"`
	Protocol       string `json:"protocol"`
	WgPort         int    `json:"wg_port"`
	Ipv6           bool   `json:"ipv6"`
	HostCheck      bool   `json:"host_check"`
	Action         string `json:"action"`
	PreferredIke   string `json:"preferred_ike"`
	PreferredEsp   string `json:"preferred_esp"`
	ForcePreferred bool   `json:"force_preferred"`
}

func (d *linkPostData) Filter() {
	d.Name = utils.FilterStr(d.Name, 1024)
	d.Type = utils.FilterId(d.Type)
	d.Status = utils.FilterId(d.Status)
	d.Protocol = utils.FilterId(d.Protocol)
	d.Action = utils.FilterId(d.Action)
	d.PreferredIke = utils.FilterId(d.PreferredIke)
	d.PreferredEsp = utils.FilterId(d.PreferredEsp)
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
	Name           string `json:"name"`
	Status         string `json:"status"`
	Protocol       string `json:"protocol"`
	WgPort         int    `json:"wg_port"`
	Key            bool   `json:"key"`
	Ipv6           bool   `json:"ipv6"`
	HostCheck      bool   `json:"host_check"`
	Action         string `json:"action"`
	PreferredIke   string `json:"preferred_ike"`
	PreferredEsp   string `json:"preferred_esp"`
	ForcePreferred bool   `json:"force_preferred"`
}

func (d *linkPutData) Filter() {
	d.Name = utils.FilterStr(d.Name, 1024)
	d.Status = utils.FilterId(d.Status)
	d.Protocol = utils.FilterId(d.Protocol)
	d.Action = utils.FilterId(d.Action)
	d.PreferredIke = utils.FilterId(d.PreferredIke)
	d.PreferredEsp = utils.FilterId(d.PreferredEsp)
}

type linkStateHostData struct {
	State   bool `json:"state"`
	Latency int  `json:"latency"`
}

type linkStatePutData struct {
	Timestamp     int64                        `json:"timestamp"`
	Version       string                       `json:"version"`
	PublicAddress string                       `json:"public_address"`
	LocalAddress  string                       `json:"local_address"`
	Address6      string                       `json:"address6"`
	Provider      string                       `json:"provider"`
	WgPublicKey   string                       `json:"wg_public_key"`
	Status        map[string]string            `json:"status"`
	Hosts         map[string]linkStateHostData `json:"hosts"`
	Errors        []string                     `json:"errors"`
}

func (d *linkStatePutData) Filter() {
	d.Version = utils.FilterStr(d.Version, 1024)
	d.PublicAddress = utils.FilterDomain(d.PublicAddress)
	d.LocalAddress = utils.FilterDomain(d.LocalAddress)
	d.Address6 = utils.FilterDomain(d.Address6)
	d.Provider = utils.FilterStr(d.Provider, 1024)
	d.WgPublicKey = utils.FilterBase64(d.WgPublicKey)

	if d.Status != nil {
		status := make(map[string]string, len(d.Status))
		for key, val := range d.Status {
			status[utils.FilterId(key)] = utils.FilterStr(val, 1024)
		}
		d.Status = status
	}

	if d.Hosts != nil {
		hosts := make(map[string]linkStateHostData, len(d.Hosts))
		for key, val := range d.Hosts {
			hosts[utils.FilterId(key)] = val
		}
		d.Hosts = hosts
	}

	for i, e := range d.Errors {
		d.Errors[i] = utils.FilterStr(e, 1024)
	}
}

func linkStatePut(c *gin.Context) {
	data := &linkStatePutData{}

	req := &request.Request{
		Method: "PUT",
		Path:   "/link/state",
		Json:   data,
	}

	req.Do(c)
}

func linkStateDelete(c *gin.Context) {
	req := &request.Request{
		Method: "DELETE",
		Path:   "/link/state",
	}

	req.Do(c)
}

func linkPut(c *gin.Context) {
	linkId := utils.FilterId(c.Params.ByName("link_id"))
	data := &linkPutData{}

	req := &request.Request{
		Method: "PUT",
		Path:   "/link/" + linkId,
		Json:   data,
	}

	req.Do(c)
}

func linkDelete(c *gin.Context) {
	linkId := utils.FilterId(c.Params.ByName("link_id"))

	req := &request.Request{
		Method: "DELETE",
		Path:   "/link/" + linkId,
	}

	req.Do(c)
}

func linkLocationGet(c *gin.Context) {
	linkId := utils.FilterId(c.Params.ByName("link_id"))

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

func (d *linkLocationPostData) Filter() {
	d.Name = utils.FilterStr(d.Name, 1024)
	d.LinkId = utils.FilterId(d.LinkId)
	d.Location = utils.FilterStr(d.Location, 1024)
}

func linkLocationPost(c *gin.Context) {
	linkId := utils.FilterId(c.Params.ByName("link_id"))
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

func (d *linkLocationPutData) Filter() {
	d.Name = utils.FilterStr(d.Name, 1024)
	d.LinkId = utils.FilterId(d.LinkId)
	d.Location = utils.FilterStr(d.Location, 1024)
}

func linkLocationPut(c *gin.Context) {
	linkId := utils.FilterId(c.Params.ByName("link_id"))
	locationId := utils.FilterId(c.Params.ByName("location_id"))
	data := &linkLocationPutData{}

	req := &request.Request{
		Method: "PUT",
		Path:   "/link/" + linkId + "/location/" + locationId,
		Json:   data,
	}

	req.Do(c)
}

func linkLocationDelete(c *gin.Context) {
	linkId := utils.FilterId(c.Params.ByName("link_id"))
	locationId := utils.FilterId(c.Params.ByName("location_id"))

	req := &request.Request{
		Method: "DELETE",
		Path:   "/link/" + linkId + "/location/" + locationId,
	}

	req.Do(c)
}

type linkLocationRoutePostData struct {
	Network string `json:"network"`
}

func (d *linkLocationRoutePostData) Filter() {
	d.Network = utils.FilterStr(d.Network, 1024)
}

func linkLocationRoutePost(c *gin.Context) {
	linkId := utils.FilterId(c.Params.ByName("link_id"))
	locationId := utils.FilterId(c.Params.ByName("location_id"))
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

func (d *linkLocationRoutePutData) Filter() {
	d.Network = utils.FilterStr(d.Network, 1024)
}

func linkLocationRoutePut(c *gin.Context) {
	linkId := utils.FilterId(c.Params.ByName("link_id"))
	locationId := utils.FilterId(c.Params.ByName("location_id"))
	routeId := utils.FilterId(c.Params.ByName("route_id"))
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
	linkId := utils.FilterId(c.Params.ByName("link_id"))
	locationId := utils.FilterId(c.Params.ByName("location_id"))
	routeId := utils.FilterId(c.Params.ByName("route_id"))

	req := &request.Request{
		Method: "DELETE",
		Path: "/link/" + linkId + "/location/" + locationId +
			"/route/" + routeId,
	}

	req.Do(c)
}

func linkLocationHostUriGet(c *gin.Context) {
	linkId := utils.FilterId(c.Params.ByName("link_id"))
	locationId := utils.FilterId(c.Params.ByName("location_id"))
	hostId := utils.FilterId(c.Params.ByName("host_id"))

	req := &request.Request{
		Method: "GET",
		Path: "/link/" + linkId + "/location/" + locationId +
			"/host/" + hostId + "/uri",
	}

	req.Do(c)
}

func linkLocationHostConfGet(c *gin.Context) {
	linkId := utils.FilterId(c.Params.ByName("link_id"))
	locationId := utils.FilterId(c.Params.ByName("location_id"))
	hostId := utils.FilterId(c.Params.ByName("host_id"))

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
	Backoff       int    `json:"backoff"`
	Static        bool   `json:"static"`
	PublicAddress string `json:"public_address"`
	LocalAddress  string `json:"local_address"`
	Address6      string `json:"address6"`
	WgPublicKey   string `json:"wg_public_key"`
}

func (d *linkLocationHostPostData) Filter() {
	d.Name = utils.FilterStr(d.Name, 1024)
	d.PublicAddress = utils.FilterDomain(d.PublicAddress)
	d.LocalAddress = utils.FilterDomain(d.LocalAddress)
	d.Address6 = utils.FilterDomain(d.Address6)
	d.WgPublicKey = utils.FilterBase64(d.WgPublicKey)
}

func linkLocationHostPost(c *gin.Context) {
	linkId := utils.FilterId(c.Params.ByName("link_id"))
	locationId := utils.FilterId(c.Params.ByName("location_id"))
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
	Backoff       int    `json:"backoff"`
	Static        bool   `json:"static"`
	PublicAddress string `json:"public_address"`
	LocalAddress  string `json:"local_address"`
	Address6      string `json:"address6"`
	WgPublicKey   string `json:"wg_public_key"`
}

func (d *linkLocationHostPutData) Filter() {
	d.Name = utils.FilterStr(d.Name, 1024)
	d.PublicAddress = utils.FilterDomain(d.PublicAddress)
	d.LocalAddress = utils.FilterDomain(d.LocalAddress)
	d.Address6 = utils.FilterDomain(d.Address6)
	d.WgPublicKey = utils.FilterBase64(d.WgPublicKey)
}

func linkLocationHostPut(c *gin.Context) {
	linkId := utils.FilterId(c.Params.ByName("link_id"))
	locationId := utils.FilterId(c.Params.ByName("location_id"))
	hostId := utils.FilterId(c.Params.ByName("host_id"))
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
	linkId := utils.FilterId(c.Params.ByName("link_id"))
	locationId := utils.FilterId(c.Params.ByName("location_id"))
	hostId := utils.FilterId(c.Params.ByName("host_id"))

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

func (d *linkLocationPeerPostData) Filter() {
	d.PeerId = utils.FilterId(d.PeerId)
}

func linkLocationPeerPost(c *gin.Context) {
	linkId := utils.FilterId(c.Params.ByName("link_id"))
	locationId := utils.FilterId(c.Params.ByName("location_id"))
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
	linkId := utils.FilterId(c.Params.ByName("link_id"))
	locationId := utils.FilterId(c.Params.ByName("location_id"))
	peerId := utils.FilterId(c.Params.ByName("peer_id"))

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

func (d *linkLocationTransitPostData) Filter() {
	d.TransitId = utils.FilterId(d.TransitId)
}

func linkLocationTransitPost(c *gin.Context) {
	linkId := utils.FilterId(c.Params.ByName("link_id"))
	locationId := utils.FilterId(c.Params.ByName("location_id"))
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
	linkId := utils.FilterId(c.Params.ByName("link_id"))
	locationId := utils.FilterId(c.Params.ByName("location_id"))
	transitId := utils.FilterId(c.Params.ByName("transit_id"))

	req := &request.Request{
		Method: "DELETE",
		Path: "/link/" + linkId + "/location/" + locationId +
			"/transit/" + transitId,
	}

	req.Do(c)
}
