package models2

import "errors"

///////////////////////////////////////////////////////////////////////////////////

// LocalAdmin ...
type LocalAdmin struct {
	ID      int
	UserID  int
	LocalID int

	CanCreate   bool
	CanRetrieve bool
	CanUpdate   bool
	CanDelete   bool

	model *model `json:"-"`

	local       *Local
	lzLoadLocal bool `json:"-"`

	admin       *User
	lzLoadAdmin bool `json:"-"`
}

// Local ...
func (loadmin *LocalAdmin) Local() (*Local, error) {
	return nil, errors.New("Not implemented yet")
}

// Admin ...
func (loadmin *LocalAdmin) Admin() (*User, error) {
	return nil, errors.New("Not implemented yet")
}

// Update ...
func (loadmin *LocalAdmin) Update() error {
	return errors.New("Not implemented yet")
}

///////////////////////////////////////////////////////////////////////////////////

// LocalAdminModel ...
type LocalAdminModel interface {
	NewLocalAdmin() LocalAdmin
}

///////////////////////////////////////////////////////////////////////////////////

func (m *model) NewLocalAdmin() LocalAdmin {
	loadmin := LocalAdmin{
		model: m,
	}
	return loadmin
}
