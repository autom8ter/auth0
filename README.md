# auth0
--
    import "github.com/autom8ter/auth0"


## Usage

#### type Client

```go
type Client struct {
	Domain string `validate:"required"`
}
```


#### func  NewClient

```go
func NewClient(clientID string, clientSecret string, domain string, scopes []string) *Client
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
