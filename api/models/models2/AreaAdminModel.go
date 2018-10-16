package models2

import "errors"

///////////////////////////////////////////////////////////////////////////////////

// AreaAdmin ...
type AreaAdmin struct {
	ID     int
	UserID int
	AreaID int

	CanCreate   bool
	CanRetrieve bool
	CanUpdate   bool
	CanDelete   bool

	model *model `json:"-"`

	area       *Area
	lzLoadArea bool `json:"-"`

	admin       *User
	lzLoadAdmin bool `json:"-"`
}

// Area ...
func (aadmin *AreaAdmin) Area() (*Area, error) {
	return nil, errors.New("Not implemented yet")
}

// Admin ...
func (aadmin *AreaAdmin) Admin() (*User, error) {
	return nil, errors.New("Not implemented yet")
}

// Update ...
func (aadmin *AreaAdmin) Update() error {
	return errors.New("Not implemented yet")
}

///////////////////////////////////////////////////////////////////////////////////

// AreaAdminModel ...
type AreaAdminModel interface {
	NewAreaAdmin() AreaAdmin
}

///////////////////////////////////////////////////////////////////////////////////

func (m *model) NewAreaAdmin() AreaAdmin {
	aadmin := AreaAdmin{
		model: m,
	}
	return aadmin
}

///////////////////////////////////////////////////////////////////////////////////
