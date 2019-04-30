//go:generate godocdown -o README.md

package oauth2

import (
	"context"
	"github.com/autom8ter/auth0/endpoints"
	"github.com/autom8ter/auth0/util"
	"golang.org/x/oauth2"
	"net/http"
	"net/url"
	"strings"
)

type Client struct {
	domain          string
	cfg             *oauth2.Config
	callbackRequest *http.Request
	ctx             context.Context
}

func (c *Client) UserInfoURL() string {
	return endpoints.UserInfoURL(c.domain)
}

func (c *Client) UsersURL() string {
	return endpoints.UsersURL(c.domain)
}

func (c *Client) UserByEmailURL() string {
	return endpoints.UserByEmailURL(c.domain)
}

func (c *Client) RolesURL() string {
	return endpoints.RolesURL(c.domain)
}

func (c *Client) LogsURL() string {
	return endpoints.LogsURL(c.domain)
}

func (c *Client) GrantsURL() string {
	return endpoints.GrantsURL(c.domain)
}

func (c *Client) StatsURL() string {
	return endpoints.StatsURL(c.domain)
}

func (c *Client) ClientsURL() string {
	return endpoints.ClientsURL(c.domain)
}

func (c *Client) JWKSURL() string {
	return endpoints.JWKSURL(c.domain)
}

func (c *Client) EmailsURL() string {
	return endpoints.EmailURL(c.domain)
}

func (c *Client) DeviceCredentialsURL() string {
	return endpoints.DeviceCredentials(c.domain)
}

func (c *Client) RulesURL() string {
	return endpoints.RulesURL(c.domain)
}

func (c *Client) CustomDomainsURL() string {
	return endpoints.CustomDomainsURL(c.domain)
}

func (c *Client) ConnectionsURL() string {
	return endpoints.ConnectionsURL(c.domain)
}

func (c *Client) ClientGrantsURL() string {
	return endpoints.ClientGrantsURL(c.domain)
}

func (c *Client) EmailTemplatesURL() string {
	return endpoints.EmailTemplateURL(c.domain)
}

func (c *Client) TenantsURL() string {
	return endpoints.TenantsURL(c.domain)
}

func NewClient(ctx context.Context, domain, clientID, clientSecret, redirect string, scopes []string, callbackRequest *http.Request) *Client {
	if ctx == nil {
		ctx = context.TODO()
	}
	return &Client{
		domain: domain,
		cfg: &oauth2.Config{
			ClientID:     clientID,
			ClientSecret: clientSecret,
			Endpoint: oauth2.Endpoint{
				AuthURL:  "https://" + domain + "/authorize",
				TokenURL: "https://" + domain + "/oauth/token",
			},
			RedirectURL: redirect,
			Scopes:      scopes,
		},
		callbackRequest: callbackRequest,
		ctx:             ctx,
	}
}

func (c *Client) Context() context.Context {
	if c.ctx == nil {
		c.ctx = context.TODO()
	}
	return c.ctx
}

func (c *Client) HTTP() (*http.Client, error) {
	code := c.callbackRequest.URL.Query().Get("code")

	t, err := c.cfg.Exchange(context.TODO(), code)
	if err != nil {
		return nil, err
	}
	return c.cfg.Client(c.Context(), t), nil
}

func (c *Client) ClientID() string {
	return c.cfg.ClientID
}

func (c *Client) ClientSecret() string {
	return c.cfg.ClientSecret
}

func (c *Client) Scopes() []string {
	return c.cfg.Scopes
}

func (c *Client) Domain() string {
	return c.domain
}

func (c *Client) TokenURL() string {
	return c.cfg.Endpoint.TokenURL
}

func (c *Client) AuthURL() string {
	return c.cfg.Endpoint.AuthURL
}

func (c *Client) Do(apiReq *http.Request) (*http.Response, error) {
	cli, err := c.HTTP()
	if err != nil {
		return nil, err
	}
	return cli.Do(apiReq)
}

func (c *Client) Post(uRL string, obj interface{}) (*http.Response, error) {
	req, err := http.NewRequest("POST", uRL, strings.NewReader(string(util.Util.MarshalJSON(obj))))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	return c.Do(req)
}

func (c *Client) PostForm(uRL string, formValues url.Values) (*http.Response, error) {
	req, err := http.NewRequest("POST", uRL, strings.NewReader(formValues.Encode()))
	if err != nil {
		return nil, err
	}
	return c.Do(req)
}

func (c *Client) Get(uRL string) (*http.Response, error) {
	req, err := http.NewRequest("GET", uRL, nil)
	if err != nil {
		return nil, err
	}
	return c.Do(req)
}

func (c *Client) Delete(uRL string) (*http.Response, error) {
	req, err := http.NewRequest("DELETE", uRL, nil)
	if err != nil {
		return nil, err
	}
	return c.Do(req)
}
