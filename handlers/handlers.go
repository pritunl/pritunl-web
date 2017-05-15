package handlers

import (
	"fmt"
	"github.com/Sirupsen/logrus"
	"github.com/dropbox/godropbox/errors"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Recovery(c *gin.Context) {
	defer func() {
		if r := recover(); r != nil {
			logrus.WithFields(logrus.Fields{
				"error": errors.New(fmt.Sprintf("%s", r)),
			}).Error("handlers: Handler panic")
			c.Writer.WriteHeader(http.StatusInternalServerError)
		}
	}()

	c.Next()
}

func Errors(c *gin.Context) {
	c.Next()
	for _, err := range c.Errors {
		logrus.WithFields(logrus.Fields{
			"error": err,
		}).Error("handlers: Handler error")
	}
}

func Register(engine *gin.Engine) {
	engine.Use(Recovery)
	engine.Use(Errors)

	engine.GET("/admin", adminGet)
	engine.GET("/admin/:admin_id", adminGet)
	engine.PUT("/admin/:admin_id", adminPut)
	engine.POST("/admin", adminPost)
	engine.DELETE("/admin/:admin_id", adminDelete)
	engine.GET("/admin/:admin_id/audit", adminAuditGet)

	engine.POST("/auth/session", authSessionPost)
	engine.DELETE("/auth/session", authSessionDelete)
	engine.GET("/state", authStateGet)

	engine.GET("/event", eventGet)
	engine.GET("/event/:cursor", eventGet)

	engine.GET("/host", hostGet)
	engine.GET("/host/:host_id", hostGet)
	engine.PUT("/host/:host_id", hostPut)
	engine.DELETE("/host/:host_id", hostDelete)
	engine.GET("/host/:host_id/usage/:period", hostUsageGet)

	engine.GET("/key/:param1", keyGet)
	engine.GET("/key/:param1/:param2", keyGet)
	engine.GET("/key/:param1/:param2/:param3/", keyGet)
	engine.GET("/key/:param1/:param2/:param3/:param4", keyGet)
	engine.GET("/key/:param1/:param2/:param3/:param4/:param5", keyGet)
	engine.GET("/key_onc/:param1", keyOncGet)
	engine.GET("/key_onc/:param1/:param2", keyOncGet)
	engine.PUT("/key_pin/:key_id", keyPinPut)
	engine.GET("/k/:short_code", keyShortGet)
	engine.DELETE("/k/:short_code", keyShortDelete)
	engine.GET("/ku/:short_code", keyApiShortGet)
	engine.POST("/sso/authenticate", ssoAuthenticatePost)
	engine.GET("/sso/request", ssoRequestGet)
	engine.GET("/sso/callback", ssoCallbackGet)
	engine.POST("/sso/duo", ssoDuoPost)
	engine.POST("/sso/yubico", ssoYubicoPost)

	engine.GET("/link", linkGet)
	engine.POST("/link", linkPost)
	engine.PUT("/link/:link_id", linkPut)
	engine.DELETE("/link/:link_id", linkDelete)
	engine.GET("/link/:link_id/location", linkLocationGet)
	engine.POST("/link/:link_id/location", linkLocationPost)
	engine.PUT("/link/:link_id/location/:location_id", linkLocationPut)
	engine.DELETE("/link/:link_id/location/:location_id", linkLocationDelete)
	engine.POST("/link/:link_id/location/:location_id/route",
		linkLocationRoutePost)
	engine.PUT("/link/:link_id/location/:location_id/route/:route_id",
		linkLocationRoutePut)
	engine.DELETE("/link/:link_id/location/:location_id/route/:route_id",
		linkLocationRouteDelete)
	engine.POST("/link/:link_id/location/:location_id/host",
		linkLocationHostPost)
	engine.PUT("/link/:link_id/location/:location_id/host/:host_id",
		linkLocationHostPut)
	engine.DELETE("/link/:link_id/location/:location_id/host/:host_id",
		linkLocationHostDelete)

	engine.GET("/log", logGet)
	engine.GET("/logs", logsGet)

	engine.GET("/organization", orgGet)
	engine.GET("/organization/:org_id", orgGet)
	engine.POST("/organization", orgPost)
	engine.PUT("/organization/:org_id", orgPut)
	engine.DELETE("/organization/:org_id", orgDelete)

	engine.GET("/ping", pingGet)
	engine.GET("/check", checkGet)

	engine.GET("/server", serverGet)
	engine.GET("/server/:server_id", serverGet)
	engine.POST("/server", serverPost)
	engine.PUT("/server/:server_id", serverPut)
	engine.DELETE("/server/:server_id", serverDelete)
	engine.GET("/server/:server_id/organization", serverOrgGet)
	engine.PUT("/server/:server_id/organization/:org_id", serverOrgPut)
	engine.DELETE("/server/:server_id/organization/:org_id", serverOrgDelete)
	engine.GET("/server/:server_id/route", serverRouteGet)
	engine.POST("/server/:server_id/route", serverRoutePost)
	engine.PUT("/server/:server_id/route/:route_net", serverRoutePut)
	engine.DELETE("/server/:server_id/route/:route_net", serverRouteDelete)
	engine.GET("/server/:server_id/host", serverHostGet)
	engine.PUT("/server/:server_id/host/:host_id", serverHostPut)
	engine.DELETE("/server/:server_id/host/:host_id", serverHostDelete)
	engine.GET("/server/:server_id/link", serverLinkGet)
	engine.PUT("/server/:server_id/link/:link_id", serverLinkPut)
	engine.DELETE("/server/:server_id/link/:link_id", serverLinkDelete)
	engine.PUT("/server/:server_id/operation/:operation", serverOperationPut)
	engine.GET("/server/:server_id/output", serverOutputGet)
	engine.DELETE("/server/:server_id/output", serverOutputDelete)
	engine.GET("/server/:server_id/link_output", serverLinkOutputGet)
	engine.DELETE("/server/:server_id/link_output", serverLinkOutputDelete)
	engine.GET("/server/:server_id/bandwidth/:period", serverBandwidthGet)

	engine.GET("/settings", settingsGet)
	engine.PUT("/settings", settingsPut)
	engine.GET("/settings/zones", settingsZonesGet)

	engine.GET("/setup", setupGet)
	engine.GET("/upgrade", upgradeGet)
	engine.GET("/setup/s/fredoka-one.eot", setupFredokaEotStaticGet)
	engine.GET("/setup/s/ubuntu-bold.eot", setupUbuntuEotStaticGet)
	engine.GET("/setup/s/fredoka-one.woff", setupFredokaWoffStaticGet)
	engine.GET("/setup/s/ubuntu-bold.woff", setupUbuntuWoffStaticGet)
	engine.PUT("/setup/mongodb", setupMongoPut)
	engine.GET("/setup/upgrade", setupUpgradeGet)

	engine.GET("/s/*path", staticPathGet)
	engine.GET("/fredoka-one.eot", fredokaEotStaticGet)
	engine.GET("/ubuntu-bold.eot", ubuntuEotStaticGet)
	engine.GET("/fredoka-one.woff", fredokaWoffStaticGet)
	engine.GET("/ubuntu-bold.woff", ubuntuWoffStaticGet)
	engine.GET("/logo.png", logoStaticGet)
	engine.GET("/robots.txt", robotsStaticGet)
	engine.GET("/", rootGet)
	engine.GET("/login", loginGet)

	engine.GET("/status", statusGet)

	engine.GET("/subscription", subscriptionGet)
	engine.GET("/subscription/styles/:plan/:ver", subscriptionStylesGet)
	engine.POST("/subscription", subscriptionPost)
	engine.PUT("/subscription", subscriptionPut)
	engine.DELETE("/subscription", subscriptionDelete)

	engine.GET("/user/:org_id", usersGet)
	engine.GET("/user/:org_id/:user_id", userGet)
	engine.POST("/user/:org_id", userPost)
	engine.POST("/user/:org_id/multi", userMultiPost)
	engine.PUT("/user/:org_id/:user_id", userPut)
	engine.DELETE("/user/:org_id/:user_id", userDelete)
	engine.PUT("/user/:org_id/:user_id/otp_secret", userOtpSecretPut)
	engine.GET("/user/:org_id/:user_id/audit", userAuditGet)
}
