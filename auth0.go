//go:generate godocdown -o README.md

package auth0

import (
	"encoding/json"
	"github.com/autom8ter/api"
	"github.com/autom8ter/auth0/driver"
	"io/ioutil"
	"net/url"
)

type Auth0 struct {
	api driver.API
}

func NewAuth0(api driver.API) *Auth0 {
	return &Auth0{
		api: api,
	}
}

func (a *Auth0) SearchUsers(perPage, pageNum int, sort, luceneQuery, fields string) ([]*api.User, error) {
	formValues := url.Values{}
	formValues.Set("per_page", string(perPage))
	formValues.Set("page_num", string(pageNum))
	formValues.Set("sort", sort)
	formValues.Set("a", luceneQuery)
	formValues.Set("fields", fields)

	resp, err := a.api.PostForm(a.api.UsersURL(), formValues)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	bits, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	users := []*api.User{}
	err = json.Unmarshal(bits, users)
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (a *Auth0) CreateUser(u *api.User) (*api.User, error) {
	resp, err := a.api.Post(a.api.UsersURL(), u)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	bits, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	users := &api.User{}
	err = json.Unmarshal(bits, users)
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (a *Auth0) GetUser(id string, fields string) (*api.User, error) {
	formValues := url.Values{}
	formValues.Set("fields", fields)
	formValues.Set("id", id)
	resp, err := a.api.PostForm(a.api.UsersURL()+"/"+id, formValues)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	bits, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	users := &api.User{}
	err = json.Unmarshal(bits, users)
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (a *Auth0) DeleteUser(id string) (*api.User, error) {
	formValues := url.Values{}
	formValues.Set("id", id)
	resp, err := a.api.PostForm(a.api.UsersURL()+"/"+id, formValues)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	bits, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	users := &api.User{}
	err = json.Unmarshal(bits, users)
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (a *Auth0) UpdateUser(u *api.User) (*api.User, error) {
	formValues := url.Values{}
	formValues.Set("id", u.UserId.Id.Text)
	resp, err := a.api.PostForm(a.api.UsersURL()+"/"+u.UserId.Id.Text, formValues)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	bits, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	users := &api.User{}
	err = json.Unmarshal(bits, users)
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (a *Auth0) GetUserByEmail(email, fields string) (*api.User, error) {
	formValues := url.Values{}
	formValues.Set("email", email)
	formValues.Set("fields", fields)

	resp, err := a.api.PostForm(a.api.UsersURL(), formValues)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	bits, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	users := &api.User{}
	err = json.Unmarshal(bits, users)
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (a *Auth0) SearchLogs(email, fields string) (*api.User, error) {
	formValues := url.Values{}
	formValues.Set("email", email)
	formValues.Set("fields", fields)

	resp, err := a.api.PostForm(a.api.UsersURL(), formValues)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	bits, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	users := &api.User{}
	err = json.Unmarshal(bits, users)
	if err != nil {
		return nil, err
	}
	return users, nil
}
