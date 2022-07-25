package handlers

import (
	"github.com/gin-gonic/gin"
)

const robots = `User-agent: *
Disallow: /
`

func robotsGet(c *gin.Context) {
	c.String(200, robots)
}
