//go:generate godocdown -o README.md

package auth0

import (
	"fmt"
	"github.com/autom8ter/auth0/endpoints"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"github.com/autom8ter/objectify"
	"golang.org/x/oauth2"
)

var util = objectify.Default()

type Client struct {
	ClientID string	`validate:"required"`
	ClientSecret string `validate:"required"`
	Domain string `validate:"required"`
	Scopes []string `validate:"required"`
	cli *http.Client
}

func NewClient(clientID string, clientSecret string, domain string, managementToken string, cli *http.Client) *Client {
	if cli == nil {
		cli = http.DefaultClient
	}
	return &Client{ClientID: clientID, ClientSecret: clientSecret, Domain: domain, ManagementToken: managementToken, cli: cli}
}

func (c *Client) Validate() error {
	return util.Validate(c)
}

func (c *Client) do(req *http.Request, contentType string) (*http.Response, error){
	req.Header.Add("content-type", contentType)
	return c.cli.Do(req)
}

func (c *Client) post(u string, obj interface{}) (*http.Response, error) {
	req, err := http.NewRequest("POST", u, strings.NewReader(string(util.MarshalJSON(obj))))
	if err != nil {
		return nil, err
	}
	return c.do(req, "application/json")
}

func (c *Client) postForm(formValues url.Values, u string) (*http.Response, error) {
	req, err := http.NewRequest("POST", u, strings.NewReader(formValues.Encode()))
	if err != nil {
		return nil, err
	}
	return c.do(req, "application/x-www-form-urlencoded")
}

func (c *Client) get(u string) (*http.Response, error) {
	req, err := http.NewRequest("GET", u, nil)
	if err != nil {
		return nil, err
	}
	return c.do(req, "application/json")
}

func (c *Client) delete(u string) (*http.Response, error) {
	req, err := http.NewRequest("DELETE", u, nil)
	if err != nil {
		return nil, err
	}
	return c.do(req, "application/json")
}

func (c *Client) OAuth2Config(redirect string) *oauth2.Config{
	return &oauth2.Config{
		ClientID:     c.ClientID,
		ClientSecret: c.ClientSecret,
		RedirectURL:  redirect,
		Scopes:       c.Scopes,
		Endpoint: oauth2.Endpoint{
			AuthURL:  "https://" + c.Domain + "/authorize",
			TokenURL: "https://" + c.Domain + "/oauth/token",
		},
	}
}