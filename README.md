# auth0
--
    import "github.com/autom8ter/auth0"


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

#### func (*Client) ClientID

```go
func (c *Client) ClientID() string
```

#### func (*Client) ClientSecret

```go
func (c *Client) ClientSecret() string
```

#### func (*Client) Context

```go
func (c *Client) Context() context.Context
```

#### func (*Client) Delete

```go
func (c *Client) Delete(uRL string) (*http.Response, error)
```

#### func (*Client) Do

```go
func (c *Client) Do(req *http.Request) (*http.Response, error)
```

#### func (*Client) Domain

```go
func (c *Client) Domain() string
```

#### func (*Client) Get

```go
func (c *Client) Get(uRL string) (*http.Response, error)
```

#### func (*Client) HTTP

```go
func (c *Client) HTTP() *http.Client
```

#### func (*Client) Post

```go
func (c *Client) Post(uRL string, obj interface{}) (*http.Response, error)
```

#### func (*Client) PostForm

```go
func (c *Client) PostForm(uRL string, formValues url.Values) (*http.Response, error)
```

#### func (*Client) Scopes

```go
func (c *Client) Scopes() []string
```

#### func (*Client) SetContext

```go
func (c *Client) SetContext(ctx context.Context)
```

#### func (*Client) TokenURL

```go
func (c *Client) TokenURL() string
```

#### func (*Client) Validate

```go
func (c *Client) Validate() error
```

#### type OAuth2Client

```go
type OAuth2Client struct {
}
```


#### func  NewOAuth2Client

```go
func NewOAuth2Client(ctx context.Context, domain, clientID, clientSecret, redirect string, scopes []string) *OAuth2Client
```

#### func (*OAuth2Client) AuthURL

```go
func (c *OAuth2Client) AuthURL() string
```

#### func (*OAuth2Client) ClientID

```go
func (c *OAuth2Client) ClientID() string
```

#### func (*OAuth2Client) ClientSecret

```go
func (c *OAuth2Client) ClientSecret() string
```

#### func (*OAuth2Client) Context

```go
func (c *OAuth2Client) Context() context.Context
```

#### func (*OAuth2Client) Domain

```go
func (c *OAuth2Client) Domain() string
```

#### func (*OAuth2Client) HTTP

```go
func (c *OAuth2Client) HTTP(t *oauth2.Token) *http.Client
```

#### func (*OAuth2Client) Scopes

```go
func (c *OAuth2Client) Scopes() []string
```

#### func (*OAuth2Client) TokenURL

```go
func (c *OAuth2Client) TokenURL() string
```
