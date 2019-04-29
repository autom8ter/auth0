# clientcredentials
--
    import "github.com/autom8ter/auth0/clientcredentials"


## Usage

#### type Client

```go
type Client struct {
}
```


#### func  NewClient

```go
func NewClient(ctx context.Context, clientID string, clientSecret string, domain string, scopes []string) *Client
```

#### func (*Client) AddScopes

```go
func (c *Client) AddScopes(scopes ...string)
```

#### func (*Client) ClientGrantsURL

```go
func (c *Client) ClientGrantsURL() string
```

#### func (*Client) ClientID

```go
func (c *Client) ClientID() string
```

#### func (*Client) ClientSecret

```go
func (c *Client) ClientSecret() string
```

#### func (*Client) ClientsURL

```go
func (c *Client) ClientsURL() string
```

#### func (*Client) ConnectionsURL

```go
func (c *Client) ConnectionsURL() string
```

#### func (*Client) Context

```go
func (c *Client) Context() context.Context
```

#### func (*Client) CustomDomainsURL

```go
func (c *Client) CustomDomainsURL() string
```

#### func (*Client) Delete

```go
func (c *Client) Delete(uRL string) (*http.Response, error)
```

#### func (*Client) DeviceCredentialsURL

```go
func (c *Client) DeviceCredentialsURL() string
```

#### func (*Client) Do

```go
func (c *Client) Do(req *http.Request) (*http.Response, error)
```

#### func (*Client) Domain

```go
func (c *Client) Domain() string
```

#### func (*Client) EmailTemplatesURL

```go
func (c *Client) EmailTemplatesURL() string
```

#### func (*Client) EmailsURL

```go
func (c *Client) EmailsURL() string
```

#### func (*Client) Get

```go
func (c *Client) Get(uRL string) (*http.Response, error)
```

#### func (*Client) GrantsURL

```go
func (c *Client) GrantsURL() string
```

#### func (*Client) HTTP

```go
func (c *Client) HTTP() *http.Client
```

#### func (*Client) JWKSURL

```go
func (c *Client) JWKSURL() string
```

#### func (*Client) LogsURL

```go
func (c *Client) LogsURL() string
```

#### func (*Client) Post

```go
func (c *Client) Post(uRL string, obj interface{}) (*http.Response, error)
```

#### func (*Client) PostForm

```go
func (c *Client) PostForm(uRL string, formValues url.Values) (*http.Response, error)
```

#### func (*Client) RolesURL

```go
func (c *Client) RolesURL() string
```

#### func (*Client) RulesURL

```go
func (c *Client) RulesURL() string
```

#### func (*Client) Scopes

```go
func (c *Client) Scopes() []string
```

#### func (*Client) SearchUsersURL

```go
func (c *Client) SearchUsersURL() string
```

#### func (*Client) SetContext

```go
func (c *Client) SetContext(ctx context.Context)
```

#### func (*Client) StatsURL

```go
func (c *Client) StatsURL() string
```

#### func (*Client) TenantsURL

```go
func (c *Client) TenantsURL() string
```

#### func (*Client) TokenURL

```go
func (c *Client) TokenURL() string
```

#### func (*Client) UserInfoURL

```go
func (c *Client) UserInfoURL() string
```

#### func (*Client) UsersURL

```go
func (c *Client) UsersURL() string
```
