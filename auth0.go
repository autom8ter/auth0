//go:generate godocdown -o README.md

package auth0

import (
	"context"
	"github.com/autom8ter/objectify"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/clientcredentials"
	"net/http"
	"net/url"
	"strings"
)

var util = objectify.Default()

type Client struct {
	domain string `validate:"required"`
	cfg *clientcredentials.Config
	ctx context.Context
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

func (c *Client) Validate() error {
	return util.Validate(c)
}

func (c *Client) Do(req *http.Request) (*http.Response, error){
	return c.HTTP().Do(req)
}

func (c *Client) Post(uRL string, obj interface{}) (*http.Response, error) {
	req, err := http.NewRequest("POST", uRL, strings.NewReader(string(util.MarshalJSON(obj))))
	if err != nil {
		return nil, err
	}
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

func (c *Client) AddScopes(scopes ...string){
	c.cfg.Scopes =append(c.cfg.Scopes, scopes...)
}

func (c *Client) SetContext(ctx context.Context){
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

func (c *OAuth2Client) TokenURL() string {
	return c.cfg.Endpoint.TokenURL
}

func (c *OAuth2Client) AuthURL() string {
	return c.cfg.Endpoint.AuthURL
}

func (c *OAuth2Client) Scopes() []string {
	return c.cfg.Scopes
}

func (c *OAuth2Client) Domain() string {
	return c.domain
}