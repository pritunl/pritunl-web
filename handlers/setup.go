package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/pritunl/pritunl-web/request"
	"github.com/pritunl/pritunl-web/utils"
)

func setupGet(c *gin.Context) {
	req := &request.Request{
		Method: "GET",
		Path:   "/setup",
	}

	req.Do(c)
}

func upgradeGet(c *gin.Context) {
	req := &request.Request{
		Method: "GET",
		Path:   "/upgrade",
	}

	req.Do(c)
}

func setupFredokaEotStaticGet(c *gin.Context) {
	req := &request.Request{
		Method: "GET",
		Path:   "/setup/s/fredoka-one.eot",
	}

	req.Do(c)
}

func setupUbuntuEotStaticGet(c *gin.Context) {
	req := &request.Request{
		Method: "GET",
		Path:   "/setup/s/ubuntu-bold.eot",
	}

	req.Do(c)
}

func setupFredokaWoffStaticGet(c *gin.Context) {
	req := &request.Request{
		Method: "GET",
		Path:   "/setup/s/fredoka-one.woff",
	}

	req.Do(c)
}

func setupUbuntuWoffStaticGet(c *gin.Context) {
	req := &request.Request{
		Method: "GET",
		Path:   "/setup/s/ubuntu-bold.woff",
	}

	req.Do(c)
}

type setupMongoPutData struct {
	SetupKey   string `json:"setup_key"`
	MongodbUri string `json:"mongodb_uri"`
}

func (d *setupMongoPutData) Filter() {
	d.SetupKey = utils.FilterId(d.SetupKey)
	d.MongodbUri = utils.FilterText(d.MongodbUri, 4096)
}

func setupMongoPut(c *gin.Context) {
	data := &setupMongoPutData{}

	req := &request.Request{
		Method: "PUT",
		Path:   "/setup/mongodb",
		Json:   data,
	}

	req.Do(c)
}

func setupUpgradeGet(c *gin.Context) {
	req := &request.Request{
		Method: "GET",
		Path:   "/setup/upgrade",
	}

	req.Do(c)
}

func successGet(c *gin.Context) {
	var query map[string]string
	mode := c.Query("mode")
	if mode == "webauth" {
		query = map[string]string{
			"mode": mode,
		}
	}

	req := &request.Request{
		Method: "GET",
		Path:   "/success",
		Query:  query,
	}

	req.Do(c)
}
