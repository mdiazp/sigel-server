package models

import (
	"github.com/astaxie/beego/validation"
)

// UserCustomModel ...
type UserCustomModel interface {
	GetUser(username string) (*User, error)
	GetUsers(prefixFilter *string, limit, offset *int,
		orderby *string, desc *bool) (*[]*User, error)
}

// GetUser ...
func (m *model) GetUser(username string) (*User, error) {
	u := m.NewUser()
	e := m.RetrieveOne(u, "username=$1", username)
	return u, e
}

// GetUsers ...
func (m *model) GetUsers(prefixFilter *string,
	limit, offset *int, orderby *string, desc *bool) (*[]*User, error) {
	users := m.NewUserCollection()

	if prefixFilter != nil {
		*prefixFilter = "k_user.username like '" + *prefixFilter + "%'"
	}
	e := m.RetrieveCollection(prefixFilter, limit, offset, orderby, desc, users)
	return users.Users, e
}

// Valid ...
func (u *User) Valid(v *validation.Validation) {
	// Validation is only doed to post or put user
	// then only roles of Admin and Superadmin are valids,
	// because Superadmin never have to be created
	rp := rolPriority(u.Rol)
	if !(0 <= rp && rp <= 2) {
		// Set error messages of Name by SetError and HasErrors will return true
		v.SetError("Rol", "Rol contain invalid value")
	}
}

// HaveRol ...
func (u *User) HaveRol(rol string) bool {
	p1 := rolPriority(rol)
	p2 := rolPriority(u.Rol)

	return p1 <= p2
}

const (
	// RolSuperadmin ...
	RolSuperadmin string = "Superadmin"
	// RolAdmin ...
	RolAdmin string = "Admin"
	// RolUser ...
	RolUser string = "User"
)

var roltypes = []string{RolUser, RolAdmin, RolSuperadmin}

// RolPriority Return RolPrioity value or -1 if the rol is invalid
func rolPriority(rol string) int {
	for i, r := range roltypes {
		if rol == r {
			return i
		}
	}
	return -1
}

// UserPublicInfo ...
type UserPublicInfo struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Name     string `json:"name"`
}
