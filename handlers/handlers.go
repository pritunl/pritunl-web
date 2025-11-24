package handlers

import (
	"encoding/base64"
	"encoding/json"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/dropbox/godropbox/errors"
	"github.com/gin-gonic/gin"
	"github.com/pritunl/pritunl-web/constants"
	"github.com/pritunl/pritunl-web/errortypes"
	"github.com/pritunl/pritunl-web/request"
	"github.com/pritunl/pritunl-web/utils"
	"github.com/sirupsen/logrus"
	"golang.org/x/crypto/nacl/secretbox"
)

type Token struct {
	Id  string `json:"id"`
	Ttl int64  `json:"ttl"`
}

func Limiter(c *gin.Context) {
	c.Request.Body = http.MaxBytesReader(c.Writer, c.Request.Body, 50000)
}

func Recovery(c *gin.Context) {
	defer func() {
		if r := recover(); r != nil {
			c.AbortWithStatus(http.StatusNotExtended)
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

func authSessionEnd(c *gin.Context) {
	c.Set("validated", false)

	req := &request.Request{
		Method: "DELETE",
		Path:   "/auth/session",
	}
	resp, err := req.Send(c)
	if err != nil {
		c.AbortWithError(500, err)
		return
	}
	defer resp.Body.Close()
}

func Unauthorize(c *gin.Context) {
	c.Set("validated", false)
}

func Authorize(c *gin.Context) {
	tokenStr, err := c.Cookie("token")
	if err != nil {
		if !constants.WebStrict {
			c.Set("validated", false)
			return
		}

		authSessionEnd(c)
		if c.Request.URL.Path == "/" {
			request.AbortRedirect(c, "/login")
		} else {
			request.AbortWithStatus(c, 401, "Missing token")
		}
		return
	}

	tokenByt, err := base64.URLEncoding.DecodeString(tokenStr)
	if err != nil {
		authSessionEnd(c)
		if c.Request.URL.Path == "/" {
			request.AbortRedirect(c, "/login")
		} else {
			request.AbortWithStatus(c, 401, "Failed to decode token")
		}
		return
	}

	if len(tokenByt) < 28 {
		authSessionEnd(c)
		if c.Request.URL.Path == "/" {
			request.AbortRedirect(c, "/login")
		} else {
			request.AbortWithStatus(c, 401, "Token length invalid")
		}
		return
	}

	var nonce [24]byte
	copy(nonce[:], tokenByt[:24])
	encByt := tokenByt[24:]

	decByt, ok := secretbox.Open(nil, encByt, &nonce, constants.WebSecret)
	if !ok {
		authSessionEnd(c)
		if c.Request.URL.Path == "/" {
			request.AbortRedirect(c, "/login")
		} else {
			request.AbortWithStatus(c, 401, "Failed to decrypt token")
		}
		return
	}

	token := &Token{}

	err = json.Unmarshal(decByt, token)
	if err != nil {
		authSessionEnd(c)
		if c.Request.URL.Path == "/" {
			request.AbortRedirect(c, "/login")
		} else {
			request.AbortWithStatus(c, 401, "Failed to unmarshal token")
		}
		return
	}

	if token.Id == "" {
		authSessionEnd(c)
		if c.Request.URL.Path == "/" {
			request.AbortRedirect(c, "/login")
		} else {
			request.AbortWithStatus(c, 401, "Token id invalid")
		}
		return
	}

	tokenTtl := time.Unix(token.Ttl, 0)
	tokenSince := time.Since(tokenTtl)

	if tokenSince < -730*time.Hour {
		authSessionEnd(c)
		if c.Request.URL.Path == "/" {
			request.AbortRedirect(c, "/login")
		} else {
			request.AbortWithStatus(c, 401, "Token session timestamp invalid")
		}
		return
	}

	if tokenSince > 0 {
		authSessionEnd(c)
		if c.Request.URL.Path == "/" {
			request.AbortRedirect(c, "/login")
		} else {
			request.AbortWithStatus(c, 401, "Token session expired")
		}
		return
	}

	c.Set("validated", true)
}

func Redirect(c *gin.Context) {
	if constants.ReverseProxyProtoHeader != "" &&
		strings.ToLower(c.Request.Header.Get(
			constants.ReverseProxyProtoHeader)) == "http" {

		u, err := url.Parse(c.Request.URL.String())
		if err != nil {
			err = errortypes.RequestError{
				errors.Wrap(err, "request: URL parse error"),
			}
			c.AbortWithError(418, err)
			return
		}

		u.Host = utils.StripPort(c.Request.Host)
		u.Scheme = "https"

		c.Redirect(http.StatusMovedPermanently, u.String())
		c.Abort()
		return
	}
}

func Register(engine *gin.Engine) {
	engine.Use(Limiter)
	engine.Use(Recovery)
	engine.Use(Redirect)

	openAuth := engine.Group("")
	openAuth.Use(Unauthorize)

	authGroup := engine.Group("")
	authGroup.Use(Authorize)

	authGroup.GET("/admin", adminGet)
	authGroup.GET("/admin/:admin_id", adminGet)
	authGroup.PUT("/admin/:admin_id", adminPut)
	authGroup.POST("/admin", adminPost)
	authGroup.DELETE("/admin/:admin_id", adminDelete)
	authGroup.GET("/admin/:admin_id/audit", adminAuditGet)

	openAuth.POST("/auth/session", authSessionPost)
	openAuth.DELETE("/auth/session", authSessionDelete)
	authGroup.GET("/state", authStateGet)

	authGroup.GET("/event", eventGet)
	authGroup.GET("/event/:cursor", eventGet)

	authGroup.GET("/device/unregistered", deviceUnregisteredGet)
	authGroup.PUT("/device/register/:org_id/:user_id/:device_id",
		deviceRegisterPut)
	authGroup.DELETE("/device/register/:org_id/:user_id/:device_id",
		deviceRegisterDelete)

	authGroup.GET("/host", hostGet)
	authGroup.GET("/host/:host_id", hostGet)
	authGroup.PUT("/host/:host_id", hostPut)
	authGroup.DELETE("/host/:host_id", hostDelete)
	authGroup.GET("/host/:host_id/usage/:period", hostUsageGet)

	authGroup.GET("/data/:org_id/:user_id", dataKeyGet)
	authGroup.GET("/data/:org_id/:user_id/:server_id", dataServerKeyGet)

	openAuth.GET("/key/:param1", keyGet)
	openAuth.GET("/key/:param1/:param2", keyGet)
	openAuth.GET("/key/:param1/:param2/:param3", keyGet)
	openAuth.GET("/key/:param1/:param2/:param3/:param4", keyGet)
	openAuth.GET("/key/:param1/:param2/:param3/:param4/:param5", keyGet)
	openAuth.POST("/key/duo", keyDuoPost)
	openAuth.POST("/key/yubico", keyYubicoPost)
	openAuth.PUT("/key_pin/:key_id", keyPinPut)
	openAuth.GET("/k/:short_code", keyShortGet)
	openAuth.DELETE("/k/:short_code", keyShortDelete)
	openAuth.GET("/ku/:short_code", keyApiShortGet)
	openAuth.POST("/key/wg/:org_id/:user_id/:server_id", keyWgPost)
	openAuth.PUT("/key/wg/:org_id/:user_id/:server_id", keyWgPut)
	openAuth.POST("/key/ovpn/:org_id/:user_id/:server_id", keyOvpnPost)
	openAuth.POST("/key/ovpn_wait/:org_id/:user_id/:server_id",
		keyOvpnWaitPost)
	openAuth.POST("/key/wg_wait/:org_id/:user_id/:server_id",
		keyWgWaitPost)
	openAuth.POST("/sso/authenticate", ssoAuthenticatePost)
	openAuth.GET("/sso/request", ssoRequestGet)
	openAuth.GET("/sso/callback", ssoCallbackGet)
	openAuth.POST("/sso/duo", ssoDuoPost)
	openAuth.POST("/sso/yubico", ssoYubicoPost)

	authGroup.GET("/link", linkGet)
	authGroup.POST("/link", linkPost)
	openAuth.PUT("/link/:link_id", linkPut)       // TODO
	openAuth.DELETE("/link/:link_id", linkDelete) // TODO
	authGroup.GET("/link/:link_id/location", linkLocationGet)
	authGroup.POST("/link/:link_id/location", linkLocationPost)
	authGroup.PUT("/link/:link_id/location/:location_id", linkLocationPut)
	authGroup.DELETE("/link/:link_id/location/:location_id",
		linkLocationDelete)
	authGroup.POST("/link/:link_id/location/:location_id/route",
		linkLocationRoutePost)
	authGroup.PUT("/link/:link_id/location/:location_id/route/:route_id",
		linkLocationRoutePut)
	authGroup.DELETE("/link/:link_id/location/:location_id/route/:route_id",
		linkLocationRouteDelete)
	authGroup.GET("/link/:link_id/location/:location_id/host/:host_id/uri",
		linkLocationHostUriGet)
	authGroup.GET("/link/:link_id/location/:location_id/host/:host_id/conf",
		linkLocationHostConfGet)
	authGroup.POST("/link/:link_id/location/:location_id/host",
		linkLocationHostPost)
	authGroup.PUT("/link/:link_id/location/:location_id/host/:host_id",
		linkLocationHostPut)
	authGroup.DELETE("/link/:link_id/location/:location_id/host/:host_id",
		linkLocationHostDelete)
	authGroup.POST("/link/:link_id/location/:location_id/peer",
		linkLocationPeerPost)
	authGroup.DELETE("/link/:link_id/location/:location_id/peer/:peer_id",
		linkLocationPeerDelete)
	authGroup.POST("/link/:link_id/location/:location_id/transit",
		linkLocationTransitPost)
	authGroup.DELETE(
		"/link/:link_id/location/:location_id/transit/:transit_id",
		linkLocationTransitDelete)

	authGroup.GET("/log", logGet)
	authGroup.GET("/logs", logsGet)

	authGroup.GET("/organization", orgGet)
	authGroup.GET("/organization/:org_id", orgGet)
	authGroup.POST("/organization", orgPost)
	authGroup.PUT("/organization/:org_id", orgPut)
	authGroup.DELETE("/organization/:org_id", orgDelete)

	openAuth.GET("/ping", pingGet)
	openAuth.GET("/check", checkGet)

	openAuth.GET("/robots.txt", robotsGet)

	authGroup.GET("/server", serverGet)
	authGroup.GET("/server/:server_id", serverGet)
	authGroup.POST("/server", serverPost)
	authGroup.PUT("/server/:server_id", serverPut)
	authGroup.DELETE("/server/:server_id", serverDelete)
	authGroup.GET("/server/:server_id/organization", serverOrgGet)
	authGroup.PUT("/server/:server_id/organization/:org_id", serverOrgPut)
	authGroup.DELETE("/server/:server_id/organization/:org_id",
		serverOrgDelete)
	authGroup.GET("/server/:server_id/route", serverRouteGet)
	authGroup.POST("/server/:server_id/route", serverRoutePost)
	authGroup.POST("/server/:server_id/routes", serverRoutesPost)
	authGroup.PUT("/server/:server_id/route/:route_net", serverRoutePut)
	authGroup.DELETE("/server/:server_id/route/:route_net", serverRouteDelete)
	authGroup.GET("/server/:server_id/host", serverHostGet)
	authGroup.PUT("/server/:server_id/host/:host_id", serverHostPut)
	authGroup.DELETE("/server/:server_id/host/:host_id", serverHostDelete)
	authGroup.GET("/server/:server_id/link", serverLinkGet)
	authGroup.PUT("/server/:server_id/link/:link_id", serverLinkPut)
	authGroup.DELETE("/server/:server_id/link/:link_id", serverLinkDelete)
	authGroup.PUT("/server/:server_id/operation/:operation",
		serverOperationPut)
	authGroup.GET("/server/:server_id/output", serverOutputGet)
	authGroup.DELETE("/server/:server_id/output", serverOutputDelete)
	authGroup.GET("/server/:server_id/link_output", serverLinkOutputGet)
	authGroup.DELETE("/server/:server_id/link_output", serverLinkOutputDelete)
	authGroup.GET("/server/:server_id/bandwidth/:period", serverBandwidthGet)

	authGroup.GET("/settings", settingsGet)
	authGroup.PUT("/settings", settingsPut)
	authGroup.GET("/settings/zones", settingsZonesGet)

	openAuth.GET("/setup", setupGet)
	openAuth.GET("/upgrade", upgradeGet)
	openAuth.GET("/setup/s/fredoka-one.eot", setupFredokaEotStaticGet)
	openAuth.GET("/setup/s/ubuntu-bold.eot", setupUbuntuEotStaticGet)
	openAuth.GET("/setup/s/fredoka-one.woff", setupFredokaWoffStaticGet)
	openAuth.GET("/setup/s/ubuntu-bold.woff", setupUbuntuWoffStaticGet)
	openAuth.PUT("/setup/mongodb", setupMongoPut)
	openAuth.GET("/setup/upgrade", setupUpgradeGet)
	openAuth.GET("/success", successGet)

	authGroup.GET("/s/*path", staticPathGet)
	openAuth.GET("/fredoka-one.eot", fredokaEotStaticGet)
	openAuth.GET("/ubuntu-bold.eot", ubuntuEotStaticGet)
	openAuth.GET("/fredoka-one.woff", fredokaWoffStaticGet)
	openAuth.GET("/ubuntu-bold.woff", ubuntuWoffStaticGet)
	openAuth.GET("/logo.png", logoStaticGet)
	authGroup.GET("/", rootGet)
	openAuth.GET("/login", loginGet)

	authGroup.GET("/status", statusGet)

	authGroup.GET("/subscription", subscriptionGet)
	authGroup.GET("/subscription/styles/:plan/:ver", subscriptionStylesGet)
	authGroup.POST("/subscription", subscriptionPost)
	authGroup.PUT("/subscription", subscriptionPut)
	authGroup.DELETE("/subscription", subscriptionDelete)

	authGroup.GET("/user/:org_id", usersGet)
	authGroup.GET("/user/:org_id/:user_id", userGet)
	authGroup.POST("/user/:org_id", userPost)
	authGroup.POST("/user/:org_id/multi", userMultiPost)
	authGroup.PUT("/user/:org_id/:user_id", userPut)
	authGroup.DELETE("/user/:org_id/:user_id", userDelete)
	authGroup.PUT("/user/:org_id/:user_id/otp_secret", userOtpSecretPut)
	authGroup.GET("/user/:org_id/:user_id/audit", userAuditGet)
	authGroup.PUT("/user/:org_id/:user_id/device/:device_id", userDevicePut)
	authGroup.DELETE("/user/:org_id/:user_id/device/:device_id",
		userDeviceDelete)
}
