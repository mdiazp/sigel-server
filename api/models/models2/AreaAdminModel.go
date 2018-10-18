package models2

import (
	"github.com/mdiazp/kmodel"
)

///////////////////////////////////////////////////////////////////////////////////

// AreaAdmin ...
type AreaAdmin struct {
	ID              int
	UserID          int
	AreaID          int
	PermissionsCRUD string
	model           Model

	area *Area
	user *User
}

/////////////////////////////////////////////////////

// TableName ...
func (aa *AreaAdmin) TableName() string {
	return "area_admin"
}

// AutoPKey ...
func (aa *AreaAdmin) AutoPKey() bool {
	return true
}

// PkeyName ...
func (aa *AreaAdmin) PkeyName() string {
	return "id"
}

// PkeyValue ...
func (aa *AreaAdmin) PkeyValue() interface{} {
	return aa.ID
}

// PkeyPointer ...
func (aa *AreaAdmin) PkeyPointer() interface{} {
	return &aa.ID
}

// ColumnNames ...
func (aa *AreaAdmin) ColumnNames() []string {
	return []string{
		"user_id",
		"area_id",
		"permissions_crud",
	}
}

// ColumnValues ...
func (aa *AreaAdmin) ColumnValues() []interface{} {
	return []interface{}{
		aa.UserID,
		aa.AreaID,
		aa.PermissionsCRUD,
	}
}

// ColumnPointers ...
func (aa *AreaAdmin) ColumnPointers() []interface{} {
	return []interface{}{
		&aa.UserID,
		&aa.AreaID,
		&aa.PermissionsCRUD,
	}
}

/////////////////////////////////////////////////////

// Update ...
func (aa *AreaAdmin) Update() error {
	return aa.model.Update2(aa)
}

// Load ...
func (aa *AreaAdmin) Load() error {
	return aa.model.Retrieve(aa)
}

// Area ...
func (aa *AreaAdmin) Area() (*Area, error) {
	var e error
	if aa.area == nil {
		aa.area = aa.model.NewArea()
		aa.area.ID = aa.AreaID
		e = aa.model.Retrieve(aa.area)
	}
	return aa.area, e
}

// User ...
func (aa *AreaAdmin) User() (*User, error) {
	var e error
	if aa.user == nil {
		aa.user = aa.model.NewUser()
		aa.user.ID = aa.UserID
		e = aa.model.Retrieve(aa.user)
	}
	return aa.user, e
}

///////////////////////////////////////////////////////////////////////////////////

// AreaAdminCollection ...
type AreaAdminCollection struct {
	model      Model
	AreaAdmins *[]*AreaAdmin
}

// NewObjectModel ...
func (c *AreaAdminCollection) NewObjectModel() kmodel.ObjectModel {
	return c.model.NewAreaAdmin()
}

// Add ...
func (c *AreaAdminCollection) Add() kmodel.ObjectModel {
	aa := c.model.NewAreaAdmin()
	*(c.AreaAdmins) = append(*(c.AreaAdmins), aa)
	return aa
}

///////////////////////////////////////////////////////////////////////////////////

// AreaAdminModel ...
type AreaAdminModel interface {
	NewAreaAdmin() *AreaAdmin
	NewAreaAdminCollection() *AreaAdminCollection
	AreaAdmins(limit, offset *int, orderby *string,
		orderDesc *bool) (*AreaAdminCollection, error)
}

/////////////////////////////////////////////////////

// NewAreaAdmin ...
func (m *model) NewAreaAdmin() *AreaAdmin {
	aa := &AreaAdmin{
		model: m,
	}
	return aa
}

// NewAreaAdminCollection ...
func (m *model) NewAreaAdminCollection() *AreaAdminCollection {
	kk := make([]*AreaAdmin, 0)
	return &AreaAdminCollection{
		model:      m,
		AreaAdmins: &kk,
	}
}

func (m *model) AreaAdmins(limit, offset *int, orderby *string,
	orderDesc *bool) (*AreaAdminCollection, error) {

	collection := m.NewAreaAdminCollection()
	e := m.RetrieveCollection(nil, limit, offset, orderby, orderDesc, collection)
	return collection, e
}
