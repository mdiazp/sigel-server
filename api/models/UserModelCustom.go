package models

import (
	"fmt"

	"github.com/astaxie/beego/validation"
)

// UserCustomModel ...
type UserCustomModel interface {
	GetUserByID(userID int) (*User, error)
	GetUser(username string) (*User, error)
	GetUsers(username *string, name *string, email *string,
		rol *string, enable *bool, limit, offset *int,
		orderby *string, desc *bool) (*[]*User, error)
}

// GetUserByID ...
func (m *model) GetUserByID(userID int) (*User, error) {
	u := m.NewUser()
	e := m.RetrieveOne(u, "id=$1", userID)
	return u, e
}

// GetUser ...
func (m *model) GetUser(username string) (*User, error) {
	u := m.NewUser()
	e := m.RetrieveOne(u, "username=$1", username)
	return u, e
}

// GetUsers ...
func (m *model) GetUsers(username *string, name *string, email *string,
	rol *string, enable *bool, limit, offset *int,
	orderby *string, desc *bool) (*[]*User, error) {
	users := m.NewUserCollection()

	where := ""

	if username != nil {
		if where != "" {
			where += " AND "
		}
		where += "k_user.username ilike '" + *username + "%'"
	}

	if name != nil {
		if where != "" {
			where += " AND "
		}
		where += "k_user.name ilike '%" + *name + "%'"
	}

	if email != nil {
		if where != "" {
			where += " AND "
		}
		where += "k_user.email ilike '%" + *email + "%'"
	}

	if rol != nil {
		if where != "" {
			where += " AND "
		}
		where += fmt.Sprintf("k_user.rol='%s'", *rol)
	}

	if enable != nil {
		if where != "" {
			where += " AND "
		}
		where += fmt.Sprintf("k_user.enable=%t", *enable)
	}

	hfilter := &where
	if where == "" {
		hfilter = nil
	}

	if orderby == nil {
		tmp := "username"
		orderby = &tmp
		tmp2 := false
		desc = &tmp2
	}

	e := m.RetrieveCollection(hfilter, limit, offset, orderby, desc, users)
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

// UserEdit ...
type UserEdit struct {
	Rol    string `json:"Rol"`
	Enable bool   `json:"Enable"`
}

// UserPublicInfo ...
type UserPublicInfo struct {
	ID       int    `json:"ID"`
	Username string `json:"Username"`
	Name     string `json:"Name"`
}

// UserProfile ...
type UserProfile struct {
	ID                       int    `json:"ID"`
	Username                 string `json:"Username"`
	Name                     string `json:"Name"`
	Email                    string `json:"Email"`
	SendNotificationsToEmail bool   `json:"SendNotificationsToEmail"`
}

// ProfileEdit ...
type ProfileEdit struct {
	Email                    string `json:"Email"`
	SendNotificationsToEmail bool   `json:"SendNotificationsToEmail"`
}
