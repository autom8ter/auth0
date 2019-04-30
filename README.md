# auth0
--
    import "github.com/autom8ter/auth0"


## Usage

#### type Auth0

```go
type Auth0 struct {
}
```


#### func  NewAuth0

```go
func NewAuth0(api driver.API) *Auth0
```

#### func (*Auth0) CreateUser

```go
func (a *Auth0) CreateUser(u *api.User) (*api.User, error)
```

#### func (*Auth0) DeleteUser

```go
func (a *Auth0) DeleteUser(id string) (*api.User, error)
```

#### func (*Auth0) GetUser

```go
func (a *Auth0) GetUser(id string, fields string) (*api.User, error)
```

#### func (*Auth0) GetUserByEmail

```go
func (a *Auth0) GetUserByEmail(email, fields string) (*api.User, error)
```

#### func (*Auth0) SearchLogs

```go
func (a *Auth0) SearchLogs(email, fields string) (*api.User, error)
```

#### func (*Auth0) SearchUsers

```go
func (a *Auth0) SearchUsers(perPage, pageNum int, sort, luceneQuery, fields string) ([]*api.User, error)
```

#### func (*Auth0) UpdateUser

```go
func (a *Auth0) UpdateUser(u *api.User) (*api.User, error)
```
