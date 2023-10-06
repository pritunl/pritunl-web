package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/pritunl/pritunl-web/request"
)

func settingsGet(c *gin.Context) {
	req := &request.Request{
		Method: "GET",
		Path:   "/settings",
	}

	req.Do(c)
}

type settingsPutData struct {
	Username              string   `json:"username"`
	Password              string   `json:"password"`
	ServerCert            string   `json:"server_cert"`
	ServerKey             string   `json:"server_key"`
	ServerPort            int      `json:"server_port"`
	AcmeDomain            string   `json:"acme_domain"`
	Auditing              string   `json:"auditing"`
	Monitoring            string   `json:"monitoring"`
	InfluxdbUrl           string   `json:"influxdb_url"`
	InfluxdbOrg           string   `json:"influxdb_org"`
	InfluxdbBucket        string   `json:"influxdb_bucket"`
	InfluxdbToken         string   `json:"influxdb_token"`
	EmailFrom             string   `json:"email_from"`
	EmailServer           string   `json:"email_server"`
	EmailUsername         string   `json:"email_username"`
	EmailPassword         string   `json:"email_password"`
	EmailTls              bool     `json:"email_tls"`
	PinMode               string   `json:"pin_mode"`
	Sso                   string   `json:"sso"`
	SsoMatch              []string `json:"sso_match"`
	SsoAzureDirectoryId   string   `json:"sso_azure_directory_id"`
	SsoAzureAppId         string   `json:"sso_azure_app_id"`
	SsoAzureAppSecret     string   `json:"sso_azure_app_secret"`
	SsoAzureRegion        string   `json:"sso_azure_region"`
	SsoAzureVersion       int      `json:"sso_azure_version"`
	SsoAuthZeroDomain     string   `json:"sso_authzero_domain"`
	SsoAuthZeroAppId      string   `json:"sso_authzero_app_id"`
	SsoAuthZeroAppSecret  string   `json:"sso_authzero_app_secret"`
	SsoGoogleKey          string   `json:"sso_google_key"`
	SsoGoogleEmail        string   `json:"sso_google_email"`
	SsoDuoToken           string   `json:"sso_duo_token"`
	SsoDuoSecret          string   `json:"sso_duo_secret"`
	SsoDuoHost            string   `json:"sso_duo_host"`
	SsoDuoMode            string   `json:"sso_duo_mode"`
	SsoYubicoClient       string   `json:"sso_yubico_client"`
	SsoYubicoSecret       string   `json:"sso_yubico_secret"`
	SsoRadiusSecret       string   `json:"sso_radius_secret"`
	SsoRadiusHost         string   `json:"sso_radius_host"`
	SsoOrg                string   `json:"sso_org"`
	SsoSamlUrl            string   `json:"sso_saml_url"`
	SsoSamlIssuerUrl      string   `json:"sso_saml_issuer_url"`
	SsoSamlCert           string   `json:"sso_saml_cert"`
	SsoOktaAppId          string   `json:"sso_okta_app_id"`
	SsoOktaMode           string   `json:"sso_okta_mode"`
	SsoOktaToken          string   `json:"sso_okta_token"`
	SsoOneloginAppId      string   `json:"sso_onelogin_app_id"`
	SsoOneloginId         string   `json:"sso_onelogin_id"`
	SsoOneloginSecret     string   `json:"sso_onelogin_secret"`
	SsoOneloginMode       string   `json:"sso_onelogin_mode"`
	SsoJumpCloudAppId     string   `json:"sso_jumpcloud_app_id"`
	SsoJumpCloudSecret    string   `json:"sso_jumpcloud_secret"`
	ServerSsoUrl          string   `json:"server_sso_url"`
	Ipv6                  bool     `json:"ipv6"`
	SsoCache              bool     `json:"sso_cache"`
	SsoClientCache        bool     `json:"sso_client_cache"`
	RestrictImport        bool     `json:"restrict_import"`
	RestrictClient        bool     `json:"restrict_client"`
	ClientReconnect       bool     `json:"client_reconnect"`
	DropPermissions       bool     `json:"drop_permissions"`
	Theme                 string   `json:"theme"`
	PublicAddress         string   `json:"public_address"`
	PublicAddress6        string   `json:"public_address6"`
	RoutedSubnet6         string   `json:"routed_subnet6"`
	RoutedSubnet6Wg       string   `json:"routed_subnet6_wg"`
	ReverseProxy          bool     `json:"reverse_proxy"`
	CloudProvider         string   `json:"cloud_provider"`
	Route53Region         string   `json:"route53_region"`
	Route53Zone           string   `json:"route53_zone"`
	OracleUserOcid        string   `json:"oracle_user_ocid"`
	OraclePublicKey       string   `json:"oracle_public_key"`
	UsEast1AccessKey      string   `json:"us_east_1_access_key"`
	UsEast1SecretKey      string   `json:"us_east_1_secret_key"`
	UsEast2AccessKey      string   `json:"us_east_2_access_key"`
	UsEast2SecretKey      string   `json:"us_east_2_secret_key"`
	UsWest1AccessKey      string   `json:"us_west_1_access_key"`
	UsWest1SecretKey      string   `json:"us_west_1_secret_key"`
	UsWest2AccessKey      string   `json:"us_west_2_access_key"`
	UsWest2SecretKey      string   `json:"us_west_2_secret_key"`
	UsEastGov1AccessKey   string   `json:"us_gov_east_1_access_key"`
	UsEastGov1SecretKey   string   `json:"us_gov_east_1_secret_key"`
	UsWestGov1AccessKey   string   `json:"us_gov_west_1_access_key"`
	UsWestGov1SecretKey   string   `json:"us_gov_west_1_secret_key"`
	EuNorth1AccessKey     string   `json:"eu_north_1_access_key"`
	EuNorth1SecretKey     string   `json:"eu_north_1_secret_key"`
	EuWest1AccessKey      string   `json:"eu_west_1_access_key"`
	EuWest1SecretKey      string   `json:"eu_west_1_secret_key"`
	EuWest2AccessKey      string   `json:"eu_west_2_access_key"`
	EuWest2SecretKey      string   `json:"eu_west_2_secret_key"`
	EuWest3AccessKey      string   `json:"eu_west_3_access_key"`
	EuWest3SecretKey      string   `json:"eu_west_3_secret_key"`
	EuCentral1AccessKey   string   `json:"eu_central_1_access_key"`
	EuCentral1SecretKey   string   `json:"eu_central_1_secret_key"`
	CaCentral1AccessKey   string   `json:"ca_central_1_access_key"`
	CaCentral1SecretKey   string   `json:"ca_central_1_secret_key"`
	CnNorth1AccessKey     string   `json:"cn_north_1_access_key"`
	CnNorth1SecretKey     string   `json:"cn_north_1_secret_key"`
	CnNorthwest1AccessKey string   `json:"cn_northwest_1_access_key"`
	CnNorthwest1SecretKey string   `json:"cn_northwest_1_secret_key"`
	ApNortheast1AccessKey string   `json:"ap_northeast_1_access_key"`
	ApNortheast1SecretKey string   `json:"ap_northeast_1_secret_key"`
	ApNortheast2AccessKey string   `json:"ap_northeast_2_access_key"`
	ApNortheast2SecretKey string   `json:"ap_northeast_2_secret_key"`
	ApSoutheast1AccessKey string   `json:"ap_southeast_1_access_key"`
	ApSoutheast1SecretKey string   `json:"ap_southeast_1_secret_key"`
	ApSoutheast2AccessKey string   `json:"ap_southeast_2_access_key"`
	ApSoutheast2SecretKey string   `json:"ap_southeast_2_secret_key"`
	ApEast1AccessKey      string   `json:"ap_east_1_access_key"`
	ApEast1SecretKey      string   `json:"ap_east_1_secret_key"`
	ApSouth1AccessKey     string   `json:"ap_south_1_access_key"`
	ApSouth1SecretKey     string   `json:"ap_south_1_secret_key"`
	SaEast1AccessKey      string   `json:"sa_east_1_access_key"`
	SaEast1SecretKey      string   `json:"sa_east_1_secret_key"`
}

func settingsPut(c *gin.Context) {
	data := &settingsPutData{}

	req := &request.Request{
		Method: "PUT",
		Path:   "/settings",
		Json:   data,
	}

	req.Do(c)
}

func settingsZonesGet(c *gin.Context) {
	req := &request.Request{
		Method: "GET",
		Path:   "/settings/zones",
	}

	req.Do(c)
}
