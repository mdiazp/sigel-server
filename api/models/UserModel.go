package models

import (
	"fmt"

	"github.com/mdiazp/kmodel"
)

///////////////////////////////////////////////////////////////////////////////////

// UserInfo ...
type UserInfo struct {
	ID                       int
	Username                 string
	Name                     string
	Email                    string
	SendNotificationsToEmail bool
	Rol                      string
	Enable                   bool
}

// User ...
type User struct {
	UserInfo
	model Model

	localAdmins *[]*LocalAdmin
}

/////////////////////////////////////////////////////

// TableName ...
func (u *User) TableName() string {
	return "k_user"
}

// AutoPKey ...
func (u *User) AutoPKey() bool {
	return true
}

// PkeyName ...
func (u *User) PkeyName() string {
	return "id"
}

// PkeyValue ...
func (u *User) PkeyValue() interface{} {
	return u.ID
}

// PkeyPointer ...
func (u *User) PkeyPointer() interface{} {
	return &u.ID
}

// ColumnNames ...
func (u *User) ColumnNames() []string {
	return []string{
		"username",
		"name",
		"email",
		"send_notifications_to_email",
		"rol",
		"enable",
	}
}

// ColumnValues ...
func (u *User) ColumnValues() []interface{} {
	return []interface{}{
		u.Username,
		u.Name,
		u.Email,
		u.SendNotificationsToEmail,
		u.Rol,
		u.Enable,
	}
}

// ColumnPointers ...
func (u *User) ColumnPointers() []interface{} {
	return []interface{}{
		&u.Username,
		&u.Name,
		&u.Email,
		&u.SendNotificationsToEmail,
		&u.Rol,
		&u.Enable,
	}
}

/////////////////////////////////////////////////////

// Update ...
func (u *User) Update() error {
	return u.model.Update(u)
}

// Load ...
func (u *User) Load() error {
	return u.model.Retrieve(u)
}

// LocalAdmins ...
func (u *User) LocalAdmins() (*[]*LocalAdmin, error) {
	var e error
	if u.localAdmins == nil {
		tmp := u.model.NewLocalAdminCollection()
		hfilter := fmt.Sprintf("user_id=%d", u.ID)
		e = u.model.RetrieveCollection(&hfilter, nil, nil, nil, nil, tmp)
		if e == nil {
			u.localAdmins = tmp.LocalAdmins
		}
	}
	return u.localAdmins, e
}

///////////////////////////////////////////////////////////////////////////////////

// UserCollection ...
type UserCollection struct {
	model Model
	Users *[]*User
}

// NewObjectModel ...
func (c *UserCollection) NewObjectModel() kmodel.ObjectModel {
	return c.model.NewUser()
}

// Add ...
func (c *UserCollection) Add() kmodel.ObjectModel {
	u := c.model.NewUser()
	*(c.Users) = append(*(c.Users), u)
	return u
}

///////////////////////////////////////////////////////////////////////////////////

// UserModel ...
type UserModel interface {
	NewUser() *User
	NewUserCollection() *UserCollection
	Users(limit, offset *int, orderby *string,
		orderDesc *bool) (*UserCollection, error)

	UserCustomModel
}

/////////////////////////////////////////////////////

// NewUser ...
func (m *model) NewUser() *User {
	u := &User{
		model: m,
	}
	return u
}

// NewUserCollection ...
func (m *model) NewUserCollection() *UserCollection {
	kk := make([]*User, 0)
	return &UserCollection{
		model: m,
		Users: &kk,
	}
}

func (m *model) Users(limit, offset *int, orderby *string,
	orderDesc *bool) (*UserCollection, error) {

	collection := m.NewUserCollection()
	e := m.RetrieveCollection(nil, limit, offset, orderby, orderDesc, collection)
	return collection, e
}
