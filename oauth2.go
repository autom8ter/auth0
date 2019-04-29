package auth0

import (
	"github.com/autom8ter/auth0/endpoints"
	"golang.org/x/oauth2"
	"net/http"
	"context"
)

type OAuth2Client struct {
	domain string
	cfg    *oauth2.Config
	ctx context.Context
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

func NewOAuth2Client(ctx context.Context, domain, clientID, clientSecret, redirect string, scopes []string) *OAuth2Client {
	if ctx == nil {
		ctx  = context.TODO()
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
			RedirectURL:  redirect,
			Scopes:       scopes,
		},
		ctx: ctx,
	}
}

func (c *OAuth2Client) Context() context.Context {
	if c.ctx == nil {
		c.ctx  = context.TODO()
	}
	return c.ctx
}

func (c *OAuth2Client) HTTP(t *oauth2.Token) *http.Client {
	return c.cfg.Client(c.Context(), t)
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

