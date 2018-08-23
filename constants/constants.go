package constants

import (
	"os"
)

var (
	ReverseProxyHeader = os.Getenv("REVERSE_PROXY_HEADER")
	RedirectServer     = os.Getenv("REDIRECT_SERVER")
	BindHost           = os.Getenv("BIND_HOST")
	BindPort           = os.Getenv("BIND_PORT")
	InternalHost       = os.Getenv("INTERNAL_ADDRESS")
	ServerHost         = os.Getenv("SERVER_HOST")
	CertPath           = os.Getenv("CERT_PATH")
	KeyPath            = os.Getenv("KEY_PATH")
	Ssl                bool
	Scheme             string
)

func init() {
	Ssl = CertPath != "" && KeyPath != ""
	if Ssl {
		Scheme = "https"
	} else {
		Scheme = "http"
	}
}
