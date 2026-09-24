package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/pritunl/pritunl-web/request"
	"github.com/pritunl/pritunl-web/utils"
)

func subscriptionGet(c *gin.Context) {
	req := &request.Request{
		Method: "GET",
		Path:   "/subscription",
	}

	req.Do(c)
}

func subscriptionStylesGet(c *gin.Context) {
	plan := utils.FilterId(c.Params.ByName("plan"))
	ver := utils.FilterStr(c.Params.ByName("ver"), 512)

	req := &request.Request{
		Method: "GET",
		Path:   "/subscription/styles/" + plan + "/" + ver,
	}

	req.Do(c)
}

type subscriptionPostData struct {
	License string `json:"license"`
}

func (d *subscriptionPostData) Filter() {
	d.License = utils.FilterText(d.License, 16384)
}

func subscriptionPost(c *gin.Context) {
	data := &subscriptionPostData{}

	req := &request.Request{
		Method: "POST",
		Path:   "/subscription",
		Json:   data,
	}

	req.Do(c)
}

type subscriptionPutData struct {
	Card      string `json:"card"`
	Email     string `json:"email"`
	Plan      string `json:"plan"`
	PromoCode string `json:"promo_code"`
	Cancel    bool   `json:"cancel"`
}

func (d *subscriptionPutData) Filter() {
	d.Card = utils.FilterId(d.Card)
	d.Email = utils.FilterStr(d.Email, 1024)
	d.Plan = utils.FilterId(d.Plan)
	d.PromoCode = utils.FilterStr(d.PromoCode, 1024)
}

func subscriptionPut(c *gin.Context) {
	data := &subscriptionPutData{}

	req := &request.Request{
		Method: "PUT",
		Path:   "/subscription",
		Json:   data,
	}

	req.Do(c)
}

func subscriptionDelete(c *gin.Context) {
	req := &request.Request{
		Method: "DELETE",
		Path:   "/subscription",
	}

	req.Do(c)
}
