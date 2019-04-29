# driver
--
    import "github.com/autom8ter/auth0/driver"


## Usage

#### type API

```go
type API interface {
	TokenURL() string
	UserInfoURL() string
	UsersURL() string
	AuthURL() string
	SearchUsersURL() string
	RolesURL() string
	LogsURL() string
	GrantsURL() string
	StatsURL() string
	ClientsURL() string
	JWKSURL() string
	EmailsURL() string
	DeviceCredentialsURL() string
	RulesURL() string
	CustomDomainsURL() string
	ConnectionsURL() string
	ClientGrantsURL() string
	EmailTemplatesURL() string
	TenantsURL() string
	Do(r *http.Request) (*http.Response, error)
	Post(uRL string, obj interface{}) (*http.Response, error)
	PostForm(uRL string, formValues url.Values) (*http.Response, error)
	Get(uRL string) (*http.Response, error)
	Delete(uRL string) (*http.Response, error)
}
```
