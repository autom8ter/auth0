# oauth2
--
    import "github.com/autom8ter/auth0/oauth2"


## Usage

#### type OAuth2Client

```go
type OAuth2Client struct {
}
```


#### func  NewOAuth2Client

```go
func NewOAuth2Client(ctx context.Context, domain, clientID, clientSecret, redirect string, scopes []string, callbackRequest *http.Request) *OAuth2Client
```

#### func (*OAuth2Client) AuthURL

```go
func (c *OAuth2Client) AuthURL() string
```

#### func (*OAuth2Client) ClientGrantsURL

```go
func (c *OAuth2Client) ClientGrantsURL() string
```

#### func (*OAuth2Client) ClientID

```go
func (c *OAuth2Client) ClientID() string
```

#### func (*OAuth2Client) ClientSecret

```go
func (c *OAuth2Client) ClientSecret() string
```

#### func (*OAuth2Client) ClientsURL

```go
func (c *OAuth2Client) ClientsURL() string
```

#### func (*OAuth2Client) ConnectionsURL

```go
func (c *OAuth2Client) ConnectionsURL() string
```

#### func (*OAuth2Client) Context

```go
func (c *OAuth2Client) Context() context.Context
```

#### func (*OAuth2Client) CustomDomainsURL

```go
func (c *OAuth2Client) CustomDomainsURL() string
```

#### func (*OAuth2Client) Delete

```go
func (c *OAuth2Client) Delete(uRL string) (*http.Response, error)
```

#### func (*OAuth2Client) DeviceCredentialsURL

```go
func (c *OAuth2Client) DeviceCredentialsURL() string
```

#### func (*OAuth2Client) Do

```go
func (c *OAuth2Client) Do(apiReq *http.Request) (*http.Response, error)
```

#### func (*OAuth2Client) Domain

```go
func (c *OAuth2Client) Domain() string
```

#### func (*OAuth2Client) EmailTemplatesURL

```go
func (c *OAuth2Client) EmailTemplatesURL() string
```

#### func (*OAuth2Client) EmailsURL

```go
func (c *OAuth2Client) EmailsURL() string
```

#### func (*OAuth2Client) Get

```go
func (c *OAuth2Client) Get(uRL string) (*http.Response, error)
```

#### func (*OAuth2Client) GrantsURL

```go
func (c *OAuth2Client) GrantsURL() string
```

#### func (*OAuth2Client) HTTP

```go
func (c *OAuth2Client) HTTP() (*http.Client, error)
```

#### func (*OAuth2Client) JWKSURL

```go
func (c *OAuth2Client) JWKSURL() string
```

#### func (*OAuth2Client) LogsURL

```go
func (c *OAuth2Client) LogsURL() string
```

#### func (*OAuth2Client) Post

```go
func (c *OAuth2Client) Post(uRL string, obj interface{}) (*http.Response, error)
```

#### func (*OAuth2Client) PostForm

```go
func (c *OAuth2Client) PostForm(uRL string, formValues url.Values) (*http.Response, error)
```

#### func (*OAuth2Client) RolesURL

```go
func (c *OAuth2Client) RolesURL() string
```

#### func (*OAuth2Client) RulesURL

```go
func (c *OAuth2Client) RulesURL() string
```

#### func (*OAuth2Client) Scopes

```go
func (c *OAuth2Client) Scopes() []string
```

#### func (*OAuth2Client) SearchUsersURL

```go
func (c *OAuth2Client) SearchUsersURL() string
```

#### func (*OAuth2Client) StatsURL

```go
func (c *OAuth2Client) StatsURL() string
```

#### func (*OAuth2Client) TenantsURL

```go
func (c *OAuth2Client) TenantsURL() string
```

#### func (*OAuth2Client) TokenURL

```go
func (c *OAuth2Client) TokenURL() string
```

#### func (*OAuth2Client) UserInfoURL

```go
func (c *OAuth2Client) UserInfoURL() string
```

#### func (*OAuth2Client) UsersURL

```go
func (c *OAuth2Client) UsersURL() string
```
