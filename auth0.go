//go:generate godocdown -o README.md

package auth0

import (
	"github.com/autom8ter/auth0/driver"
)

type Auth0 struct {
	API driver.API
}

func NewAuth0(api driver.API) *Auth0 {
	return &Auth0{
		API: api,
	}
}
