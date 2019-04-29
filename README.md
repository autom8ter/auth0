# auth0
--
    import "github.com/autom8ter/auth0"


## Usage

#### type Client

```go
type Client struct {
	ClientID     string   `validate:"required"`
	ClientSecret string   `validate:"required"`
	Domain       string   `validate:"required"`
	Scopes       []string `validate:"required"`
}
```


#### func  NewClient

```go
func NewClient(clientID string, clientSecret string, domain string, managementToken string, cli *http.Client) *Client
```

#### func (*Client) OAuth2Config

```go
func (c *Client) OAuth2Config(redirect string) *oauth2.Config
```

#### func (*Client) Validate

```go
func (c *Client) Validate() error
```
