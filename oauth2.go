package auth0

import (
	"golang.org/x/oauth2"
	"net/http"
	"context"
)

type OAuth2Client struct {
	domain string
	cfg    *oauth2.Config
	ctx context.Context
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

