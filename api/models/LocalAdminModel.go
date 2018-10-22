package models

import (
	"github.com/mdiazp/kmodel"
)

///////////////////////////////////////////////////////////////////////////////////

// LocalAdminInfo ...
type LocalAdminInfo struct {
	ID      int
	UserID  int
	LocalID int
}

// LocalAdmin ...
type LocalAdmin struct {
	LocalAdminInfo
	model Model

	local *Local
	user  *User
}

/////////////////////////////////////////////////////

// TableName ...
func (la *LocalAdmin) TableName() string {
	return "local_admin"
}

// AutoPKey ...
func (la *LocalAdmin) AutoPKey() bool {
	return true
}

// PkeyName ...
func (la *LocalAdmin) PkeyName() string {
	return "id"
}

// PkeyValue ...
func (la *LocalAdmin) PkeyValue() interface{} {
	return la.ID
}

// PkeyPointer ...
func (la *LocalAdmin) PkeyPointer() interface{} {
	return &la.ID
}

// ColumnNames ...
func (la *LocalAdmin) ColumnNames() []string {
	return []string{
		"user_id",
		"local_id",
	}
}

// ColumnValues ...
func (la *LocalAdmin) ColumnValues() []interface{} {
	return []interface{}{
		la.UserID,
		la.LocalID,
	}
}

// ColumnPointers ...
func (la *LocalAdmin) ColumnPointers() []interface{} {
	return []interface{}{
		&la.UserID,
		&la.LocalID,
	}
}

/////////////////////////////////////////////////////

// Update ...
func (la *LocalAdmin) Update() error {
	return la.model.Update(la)
}

// Load ...
func (la *LocalAdmin) Load() error {
	return la.model.Retrieve(la)
}

// Local ...
func (la *LocalAdmin) Local() (*Local, error) {
	var e error
	if la.local == nil {
		la.local = la.model.NewLocal()
		la.local.ID = la.ID
		e = la.model.Retrieve(la.local)
	}
	return la.local, e
}

// User ...
func (la *LocalAdmin) User() (*User, error) {
	var e error
	if la.user == nil {
		la.user = la.model.NewUser()
		la.user.ID = la.UserID
		e = la.model.Retrieve(la.user)
	}
	return la.user, e
}

///////////////////////////////////////////////////////////////////////////////////

// LocalAdminCollection ...
type LocalAdminCollection struct {
	model       Model
	LocalAdmins *[]*LocalAdmin
}

// NewObjectModel ...
func (c *LocalAdminCollection) NewObjectModel() kmodel.ObjectModel {
	return c.model.NewLocalAdmin()
}

// Add ...
func (c *LocalAdminCollection) Add() kmodel.ObjectModel {
	la := c.model.NewLocalAdmin()
	*(c.LocalAdmins) = append(*(c.LocalAdmins), la)
	return la
}

///////////////////////////////////////////////////////////////////////////////////

// LocalAdminModel ...
type LocalAdminModel interface {
	NewLocalAdmin() *LocalAdmin
	NewLocalAdminCollection() *LocalAdminCollection
	LocalAdmins(limit, offset *int, orderby *string,
		orderDesc *bool) (*LocalAdminCollection, error)

	LocalAdminCustomModel
}

/////////////////////////////////////////////////////

// NewLocalAdmin ...
func (m *model) NewLocalAdmin() *LocalAdmin {
	la := &LocalAdmin{
		model: m,
	}
	return la
}

// NewLocalAdminCollection ...
func (m *model) NewLocalAdminCollection() *LocalAdminCollection {
	kk := make([]*LocalAdmin, 0)
	return &LocalAdminCollection{
		model:       m,
		LocalAdmins: &kk,
	}
}

func (m *model) LocalAdmins(limit, offset *int, orderby *string,
	orderDesc *bool) (*LocalAdminCollection, error) {

	collection := m.NewLocalAdminCollection()
	e := m.RetrieveCollection(nil, limit, offset, orderby, orderDesc, collection)
	return collection, e
}
