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

type OAuth2Client struct {
	domain          string
	cfg             *oauth2.Config
	callbackRequest *http.Request
	ctx             context.Context
}

func (c *OAuth2Client) UserInfoURL() string {
	return endpoints.UserInfoURL(c.domain)
}

func (c *OAuth2Client) UsersURL() string {
	return endpoints.UsersURL(c.domain)
}

func (c *OAuth2Client) SearchUsersURL() string {
	return endpoints.SearchUsersURL(c.domain)
}

func (c *OAuth2Client) RolesURL() string {
	return endpoints.RolesURL(c.domain)
}

func (c *OAuth2Client) LogsURL() string {
	return endpoints.LogsURL(c.domain)
}

func (c *OAuth2Client) GrantsURL() string {
	return endpoints.GrantsURL(c.domain)
}

func (c *OAuth2Client) StatsURL() string {
	return endpoints.StatsURL(c.domain)
}

func (c *OAuth2Client) ClientsURL() string {
	return endpoints.ClientsURL(c.domain)
}

func (c *OAuth2Client) JWKSURL() string {
	return endpoints.JWKSURL(c.domain)
}

func (c *OAuth2Client) EmailsURL() string {
	return endpoints.EmailURL(c.domain)
}

func (c *OAuth2Client) DeviceCredentialsURL() string {
	return endpoints.DeviceCredentials(c.domain)
}

func (c *OAuth2Client) RulesURL() string {
	return endpoints.RulesURL(c.domain)
}

func (c *OAuth2Client) CustomDomainsURL() string {
	return endpoints.CustomDomainsURL(c.domain)
}

func (c *OAuth2Client) ConnectionsURL() string {
	return endpoints.ConnectionsURL(c.domain)
}

func (c *OAuth2Client) ClientGrantsURL() string {
	return endpoints.ClientGrantsURL(c.domain)
}

func (c *OAuth2Client) EmailTemplatesURL() string {
	return endpoints.EmailTemplateURL(c.domain)
}

func (c *OAuth2Client) TenantsURL() string {
	return endpoints.TenantsURL(c.domain)
}

func NewOAuth2Client(ctx context.Context, domain, clientID, clientSecret, redirect string, scopes []string, callbackRequest *http.Request) *OAuth2Client {
	if ctx == nil {
		ctx = context.TODO()
	}
	return &OAuth2Client{
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

func (c *OAuth2Client) Context() context.Context {
	if c.ctx == nil {
		c.ctx = context.TODO()
	}
	return c.ctx
}

func (c *OAuth2Client) HTTP() (*http.Client, error) {
	code := c.callbackRequest.URL.Query().Get("code")

	t, err := c.cfg.Exchange(context.TODO(), code)
	if err != nil {
		return nil, err
	}
	return c.cfg.Client(c.Context(), t), nil
}

func (c *OAuth2Client) ClientID() string {
	return c.cfg.ClientID
}

func (c *OAuth2Client) ClientSecret() string {
	return c.cfg.ClientSecret
}

func (c *OAuth2Client) Scopes() []string {
	return c.cfg.Scopes
}

func (c *OAuth2Client) Domain() string {
	return c.domain
}

func (c *OAuth2Client) TokenURL() string {
	return c.cfg.Endpoint.TokenURL
}

func (c *OAuth2Client) AuthURL() string {
	return c.cfg.Endpoint.AuthURL
}

func (c *OAuth2Client) Do(apiReq *http.Request) (*http.Response, error) {
	cli, err := c.HTTP()
	if err != nil {
		return nil, err
	}
	return cli.Do(apiReq)
}

func (c *OAuth2Client) Post(uRL string, obj interface{}) (*http.Response, error) {
	req, err := http.NewRequest("POST", uRL, strings.NewReader(string(util.Util.MarshalJSON(obj))))
	if err != nil {
		return nil, err
	}
	return c.Do(req)
}

func (c *OAuth2Client) PostForm(uRL string, formValues url.Values) (*http.Response, error) {
	req, err := http.NewRequest("POST", uRL, strings.NewReader(formValues.Encode()))
	if err != nil {
		return nil, err
	}
	return c.Do(req)
}

func (c *OAuth2Client) Get(uRL string) (*http.Response, error) {
	req, err := http.NewRequest("GET", uRL, nil)
	if err != nil {
		return nil, err
	}
	return c.Do(req)
}

func (c *OAuth2Client) Delete(uRL string) (*http.Response, error) {
	req, err := http.NewRequest("DELETE", uRL, nil)
	if err != nil {
		return nil, err
	}
	return c.Do(req)
}
