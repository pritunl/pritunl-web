package constants

import (
	"os"
)

var (
	ReverseProxyHeader      = os.Getenv("REVERSE_PROXY_HEADER")
	ReverseProxyProtoHeader = os.Getenv("REVERSE_PROXY_PROTO_HEADER")
	RedirectServer          = os.Getenv("REDIRECT_SERVER")
	BindHost                = os.Getenv("BIND_HOST")
	BindPort                = os.Getenv("BIND_PORT")
	InternalHost            = os.Getenv("INTERNAL_ADDRESS")
	SslCert                 = os.Getenv("SSL_CERT")
	SslKey                  = os.Getenv("SSL_KEY")
	Ssl                     bool
	Scheme                  string
)

func init() {
	Ssl = SslCert != "" && SslKey != ""
	if Ssl {
		Scheme = "https"
	} else {
		Scheme = "http"
	}
}
