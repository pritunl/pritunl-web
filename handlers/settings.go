package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/pritunl/pritunl-web/request"
	"github.com/pritunl/pritunl-web/utils"
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
	PritunlCloudHost      string   `json:"pritunl_cloud_host"`
	PritunlCloudToken     string   `json:"pritunl_cloud_token"`
	PritunlCloudSecret    string   `json:"pritunl_cloud_secret"`
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
	ApSoutheast3AccessKey string   `json:"ap_southeast_3_access_key"`
	ApSoutheast3SecretKey string   `json:"ap_southeast_3_secret_key"`
	ApEast1AccessKey      string   `json:"ap_east_1_access_key"`
	ApEast1SecretKey      string   `json:"ap_east_1_secret_key"`
	ApSouth1AccessKey     string   `json:"ap_south_1_access_key"`
	ApSouth1SecretKey     string   `json:"ap_south_1_secret_key"`
	SaEast1AccessKey      string   `json:"sa_east_1_access_key"`
	SaEast1SecretKey      string   `json:"sa_east_1_secret_key"`
}

func (d *settingsPutData) Filter() {
	d.Username = utils.FilterStr(d.Username, 1024)
	d.Password = utils.FilterText(d.Password, 512)
	d.ServerCert = utils.FilterPem(d.ServerCert)
	d.ServerKey = utils.FilterPem(d.ServerKey)
	d.AcmeDomain = utils.FilterDomain(d.AcmeDomain)
	d.Auditing = utils.FilterId(d.Auditing)
	d.Monitoring = utils.FilterId(d.Monitoring)
	d.InfluxdbUrl = utils.FilterStr(d.InfluxdbUrl, 4096)
	d.InfluxdbOrg = utils.FilterText(d.InfluxdbOrg, 1024)
	d.InfluxdbBucket = utils.FilterText(d.InfluxdbBucket, 1024)
	d.InfluxdbToken = utils.FilterStr(d.InfluxdbToken, 4096)
	d.EmailFrom = utils.FilterStr(d.EmailFrom, 1024)
	d.EmailServer = utils.FilterDomain(d.EmailServer)
	d.EmailUsername = utils.FilterStr(d.EmailUsername, 1024)
	d.EmailPassword = utils.FilterText(d.EmailPassword, 512)
	d.PinMode = utils.FilterId(d.PinMode)
	d.Sso = utils.FilterId(d.Sso)
	d.SsoAzureDirectoryId = utils.FilterId(d.SsoAzureDirectoryId)
	d.SsoAzureAppId = utils.FilterId(d.SsoAzureAppId)
	d.SsoAzureAppSecret = utils.FilterStr(d.SsoAzureAppSecret, 512)
	d.SsoAzureRegion = utils.FilterId(d.SsoAzureRegion)
	d.SsoAuthZeroDomain = utils.FilterDomain(d.SsoAuthZeroDomain)
	d.SsoAuthZeroAppId = utils.FilterId(d.SsoAuthZeroAppId)
	d.SsoAuthZeroAppSecret = utils.FilterStr(d.SsoAuthZeroAppSecret, 512)
	d.SsoGoogleKey = utils.FilterText(d.SsoGoogleKey, 32768)
	d.SsoGoogleEmail = utils.FilterStr(d.SsoGoogleEmail, 1024)
	d.SsoDuoToken = utils.FilterId(d.SsoDuoToken)
	d.SsoDuoSecret = utils.FilterStr(d.SsoDuoSecret, 512)
	d.SsoDuoHost = utils.FilterDomain(d.SsoDuoHost)
	d.SsoDuoMode = utils.FilterId(d.SsoDuoMode)
	d.SsoYubicoClient = utils.FilterId(d.SsoYubicoClient)
	d.SsoYubicoSecret = utils.FilterStr(d.SsoYubicoSecret, 512)
	d.SsoRadiusSecret = utils.FilterText(d.SsoRadiusSecret, 1024)
	d.SsoRadiusHost = utils.FilterDomain(d.SsoRadiusHost)
	d.SsoOrg = utils.FilterId(d.SsoOrg)
	d.SsoSamlUrl = utils.FilterStr(d.SsoSamlUrl, 4096)
	d.SsoSamlIssuerUrl = utils.FilterStr(d.SsoSamlIssuerUrl, 4096)
	d.SsoSamlCert = utils.FilterPem(d.SsoSamlCert)
	d.SsoOktaAppId = utils.FilterId(d.SsoOktaAppId)
	d.SsoOktaMode = utils.FilterId(d.SsoOktaMode)
	d.SsoOktaToken = utils.FilterStr(d.SsoOktaToken, 512)
	d.SsoOneloginAppId = utils.FilterId(d.SsoOneloginAppId)
	d.SsoOneloginId = utils.FilterId(d.SsoOneloginId)
	d.SsoOneloginSecret = utils.FilterStr(d.SsoOneloginSecret, 512)
	d.SsoOneloginMode = utils.FilterId(d.SsoOneloginMode)
	d.SsoJumpCloudAppId = utils.FilterId(d.SsoJumpCloudAppId)
	d.SsoJumpCloudSecret = utils.FilterStr(d.SsoJumpCloudSecret, 512)
	d.ServerSsoUrl = utils.FilterStr(d.ServerSsoUrl, 4096)
	d.Theme = utils.FilterId(d.Theme)
	d.PublicAddress = utils.FilterDomain(d.PublicAddress)
	d.PublicAddress6 = utils.FilterDomain(d.PublicAddress6)
	d.RoutedSubnet6 = utils.FilterStr(d.RoutedSubnet6, 1024)
	d.RoutedSubnet6Wg = utils.FilterStr(d.RoutedSubnet6Wg, 1024)
	d.CloudProvider = utils.FilterId(d.CloudProvider)
	d.Route53Region = utils.FilterId(d.Route53Region)
	d.Route53Zone = utils.FilterStr(d.Route53Zone, 1024)
	d.OracleUserOcid = utils.FilterStr(d.OracleUserOcid, 1024)
	d.OraclePublicKey = utils.FilterPem(d.OraclePublicKey)
	d.PritunlCloudHost = utils.FilterDomain(d.PritunlCloudHost)
	d.PritunlCloudToken = utils.FilterStr(d.PritunlCloudToken, 4096)
	d.PritunlCloudSecret = utils.FilterStr(d.PritunlCloudSecret, 4096)
	d.UsEast1AccessKey = utils.FilterStr(d.UsEast1AccessKey, 4096)
	d.UsEast1SecretKey = utils.FilterStr(d.UsEast1SecretKey, 4096)
	d.UsEast2AccessKey = utils.FilterStr(d.UsEast2AccessKey, 4096)
	d.UsEast2SecretKey = utils.FilterStr(d.UsEast2SecretKey, 4096)
	d.UsWest1AccessKey = utils.FilterStr(d.UsWest1AccessKey, 4096)
	d.UsWest1SecretKey = utils.FilterStr(d.UsWest1SecretKey, 4096)
	d.UsWest2AccessKey = utils.FilterStr(d.UsWest2AccessKey, 4096)
	d.UsWest2SecretKey = utils.FilterStr(d.UsWest2SecretKey, 4096)
	d.UsEastGov1AccessKey = utils.FilterStr(d.UsEastGov1AccessKey, 4096)
	d.UsEastGov1SecretKey = utils.FilterStr(d.UsEastGov1SecretKey, 4096)
	d.UsWestGov1AccessKey = utils.FilterStr(d.UsWestGov1AccessKey, 4096)
	d.UsWestGov1SecretKey = utils.FilterStr(d.UsWestGov1SecretKey, 4096)
	d.EuNorth1AccessKey = utils.FilterStr(d.EuNorth1AccessKey, 4096)
	d.EuNorth1SecretKey = utils.FilterStr(d.EuNorth1SecretKey, 4096)
	d.EuWest1AccessKey = utils.FilterStr(d.EuWest1AccessKey, 4096)
	d.EuWest1SecretKey = utils.FilterStr(d.EuWest1SecretKey, 4096)
	d.EuWest2AccessKey = utils.FilterStr(d.EuWest2AccessKey, 4096)
	d.EuWest2SecretKey = utils.FilterStr(d.EuWest2SecretKey, 4096)
	d.EuWest3AccessKey = utils.FilterStr(d.EuWest3AccessKey, 4096)
	d.EuWest3SecretKey = utils.FilterStr(d.EuWest3SecretKey, 4096)
	d.EuCentral1AccessKey = utils.FilterStr(d.EuCentral1AccessKey, 4096)
	d.EuCentral1SecretKey = utils.FilterStr(d.EuCentral1SecretKey, 4096)
	d.CaCentral1AccessKey = utils.FilterStr(d.CaCentral1AccessKey, 4096)
	d.CaCentral1SecretKey = utils.FilterStr(d.CaCentral1SecretKey, 4096)
	d.CnNorth1AccessKey = utils.FilterStr(d.CnNorth1AccessKey, 4096)
	d.CnNorth1SecretKey = utils.FilterStr(d.CnNorth1SecretKey, 4096)
	d.CnNorthwest1AccessKey = utils.FilterStr(d.CnNorthwest1AccessKey, 4096)
	d.CnNorthwest1SecretKey = utils.FilterStr(d.CnNorthwest1SecretKey, 4096)
	d.ApNortheast1AccessKey = utils.FilterStr(d.ApNortheast1AccessKey, 4096)
	d.ApNortheast1SecretKey = utils.FilterStr(d.ApNortheast1SecretKey, 4096)
	d.ApNortheast2AccessKey = utils.FilterStr(d.ApNortheast2AccessKey, 4096)
	d.ApNortheast2SecretKey = utils.FilterStr(d.ApNortheast2SecretKey, 4096)
	d.ApSoutheast1AccessKey = utils.FilterStr(d.ApSoutheast1AccessKey, 4096)
	d.ApSoutheast1SecretKey = utils.FilterStr(d.ApSoutheast1SecretKey, 4096)
	d.ApSoutheast2AccessKey = utils.FilterStr(d.ApSoutheast2AccessKey, 4096)
	d.ApSoutheast2SecretKey = utils.FilterStr(d.ApSoutheast2SecretKey, 4096)
	d.ApSoutheast3AccessKey = utils.FilterStr(d.ApSoutheast3AccessKey, 4096)
	d.ApSoutheast3SecretKey = utils.FilterStr(d.ApSoutheast3SecretKey, 4096)
	d.ApEast1AccessKey = utils.FilterStr(d.ApEast1AccessKey, 4096)
	d.ApEast1SecretKey = utils.FilterStr(d.ApEast1SecretKey, 4096)
	d.ApSouth1AccessKey = utils.FilterStr(d.ApSouth1AccessKey, 4096)
	d.ApSouth1SecretKey = utils.FilterStr(d.ApSouth1SecretKey, 4096)
	d.SaEast1AccessKey = utils.FilterStr(d.SaEast1AccessKey, 4096)
	d.SaEast1SecretKey = utils.FilterStr(d.SaEast1SecretKey, 4096)

	for i, match := range d.SsoMatch {
		d.SsoMatch[i] = utils.FilterStr(match, 1024)
	}
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
