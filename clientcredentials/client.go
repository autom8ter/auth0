//go:generate godocdown -o README.md

package clientcredentials

import (
	"context"
	"github.com/autom8ter/auth0/endpoints"
	"github.com/autom8ter/auth0/util"
	"golang.org/x/oauth2/clientcredentials"
	"net/http"
	"net/url"
	"strings"
)

type Client struct {
	domain string
	cfg    *clientcredentials.Config
	ctx    context.Context
}

func (c *Client) AuthURL() string {
	return endpoints.AuthURL(c.domain)
}

func NewClient(ctx context.Context, clientID string, clientSecret string, domain string, scopes []string) *Client {
	return &Client{
		domain: domain,
		cfg: &clientcredentials.Config{
			ClientID:     clientID,
			ClientSecret: clientSecret,
			TokenURL:     "https://" + domain + "/oauth/token",
			Scopes:       scopes,
		},
		ctx: ctx,
	}
}

func (c *Client) Do(req *http.Request) (*http.Response, error) {
	return c.HTTP().Do(req)
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

func (c *Client) AddScopes(scopes ...string) {
	c.cfg.Scopes = append(c.cfg.Scopes, scopes...)
}

func (c *Client) SetContext(ctx context.Context) {
	c.ctx = ctx
}

func (c *Client) Context() context.Context {
	if c.ctx == nil {
		c.ctx = context.TODO()
	}
	return c.ctx
}

func (c *Client) HTTP() *http.Client {
	return c.cfg.Client(c.Context())
}

func (c *Client) ClientID() string {
	return c.cfg.ClientID
}

func (c *Client) ClientSecret() string {
	return c.cfg.ClientSecret
}

func (c *Client) TokenURL() string {
	return c.cfg.TokenURL
}

func (c *Client) Scopes() []string {
	return c.cfg.Scopes
}

func (c *Client) Domain() string {
	return c.domain
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
