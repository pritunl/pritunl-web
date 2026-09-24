package handlers

import (
	"path"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/pritunl/pritunl-web/request"
	"github.com/pritunl/pritunl-web/utils"
)

func staticPathGet(c *gin.Context) {
	pth := utils.FilterStr(c.Params.ByName("path"), 512)
	pth = strings.Replace(pth, "..", "", -1)
	pth = path.Clean(pth)

	req := &request.Request{
		Method: "GET",
		Path:   "/s" + pth,
	}

	req.Do(c)
}

func fredokaEotStaticGet(c *gin.Context) {
	req := &request.Request{
		Method: "GET",
		Path:   "/fredoka-one.eot",
	}

	req.Do(c)
}

func ubuntuEotStaticGet(c *gin.Context) {
	req := &request.Request{
		Method: "GET",
		Path:   "/ubuntu-bold.eot",
	}

	req.Do(c)
}

func fredokaWoffStaticGet(c *gin.Context) {
	req := &request.Request{
		Method: "GET",
		Path:   "/fredoka-one.woff",
	}

	req.Do(c)
}

func ubuntuWoffStaticGet(c *gin.Context) {
	req := &request.Request{
		Method: "GET",
		Path:   "/ubuntu-bold.woff",
	}

	req.Do(c)
}

func logoStaticGet(c *gin.Context) {
	req := &request.Request{
		Method: "GET",
		Path:   "/logo.png",
	}

	req.Do(c)
}

func rootGet(c *gin.Context) {
	req := &request.Request{
		Method: "GET",
		Path:   "/",
	}

	req.Do(c)
}

func loginGet(c *gin.Context) {
	req := &request.Request{
		Method: "GET",
		Path:   "/login",
	}

	req.Do(c)
}
