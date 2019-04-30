//go:generate godocdown -o README.md

package endpoints

func TokenURL(domain string) string {
	return "https://" + domain + "/oauth/token"
}

func UserInfoURL(domain string) string {
	return "https://" + domain + "/userinfo"
}

func UsersURL(domain string) string {
	return "https://" + domain + "/api/v2/users"
}

func AuthURL(domain string) string {
	return "https://" + domain + "/authorize"
}

func UserByEmailURL(domain string) string {
	return "https://" + domain + "/api/v2/users-by-email"
}

func RolesURL(domain string) string {
	return "https://" + domain + "/api/v2/roles"
}

func LogsURL(domain string) string {
	return "https://" + domain + "/api/v2/logs"
}

func GrantsURL(domain string) string {
	return "https://" + domain + "/api/v2/grants"
}

func StatsURL(domain string) string {
	return "https://" + domain + "/api/v2/stats/active-users"
}

func ClientsURL(domain string) string {

	return "https://" + domain + "/api/v2/clients"
}

func ClientGrantsURL(domain string) string {
	return "https://" + domain + "/api/v2/client-grants"
}

func ConnectionsURL(domain string) string {
	return "https://" + domain + "/api/v2/connections"
}

func CustomDomainsURL(domain string) string {
	return "https://" + domain + "/api/v2/custom-domains"
}

func RulesURL(domain string) string {
	return "https://" + domain + "/api/v2/rules"
}

func TenantsURL(domain string) string {
	return "https://" + domain + "/api/v2/tenants/settings"
}

func EmailURL(domain string) string {
	return "https://" + domain + "/api/v2/emails/provider"
}

func EmailTemplateURL(domain string) string {
	return "https://" + domain + "/api/v2/email-templates"
}

func DeviceCredentials(domain string) string {
	return "https://" + domain + "/api/v2/device-credentials"
}

func JWKSURL(domain string) string {
	return "https://" + domain + "/.well-known/jwks.json"
}
