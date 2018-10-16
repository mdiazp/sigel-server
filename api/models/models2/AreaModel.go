package models2

import (
	"errors"
)

///////////////////////////////////////////////////////////////////////////////////

// Area ...
type Area struct {
	ID          int    `json:"id"`
	Name        string `json:"name" valid:"Required;MaxSize(100)"`
	Description string `json:"description" valid:"Required;MaxSize(1024)"`
	Location    string `json:"locatiion" valid:"Required;MaxSize(1024)"`

	model *model

	locals       *[]*Local
	lzLoadLocals bool

	admins       *[]*AreaAdmin
	lzLoadAdmins bool
}

///////////////////////////////////////////////////////////////////////////////////

func (a *Area) tableName() string {
	return "area"
}

func (a *Area) autoPKey() bool {
	return true
}

func (a *Area) pkeyName() string {
	return "id"
}

func (a *Area) pkeyValue() interface{} {
	return a.ID
}

func (a *Area) pkeyPointer() interface{} {
	return &a.ID
}

func (a *Area) columnNames(pk ...bool) []string {
	names := []string{"name", "description", "location"}
	if len(pk) > 0 && pk[0] {
		names = append(names, a.pkeyName())
	}
	return names
}

func (a *Area) columnValues(pk ...bool) []interface{} {
	values := []interface{}{a.Name, a.Description, a.Location}
	if len(pk) > 0 && pk[0] {
		values = append(values, a.pkeyValue())
	}
	return values
}

func (a *Area) columnPointers(pk ...bool) []interface{} {
	pointers := []interface{}{&a.Name, &a.Description, &a.Location}
	if len(pk) > 0 && pk[0] {
		pointers = append(pointers, a.pkeyPointer())
	}
	return pointers
}

///////////////////////////////////////////////////////////////////////////////////

// Update ...
func (a *Area) Update() error {
	return a.model.Update2(a)
}

// Locals ...
func (a *Area) Locals() (*[]*Local, error) {
	return nil, errors.New("Not implemented yet")
}

// AddLocal ...
func (a *Area) AddLocal(lo *Local) error {
	return errors.New("Not implemented yet")
}

// DeleteLocal ...
func (a *Area) DeleteLocal(lo *Local) error {
	return errors.New("Not implemented yet")
}

// Admins ...
func (a *Area) Admins() ([]*AreaAdmin, error) {
	return nil, errors.New("Not implemented yet")
}

// AddAdmin ...
func (a *Area) AddAdmin(aa *AreaAdmin) error {
	return errors.New("Not implemented yet")
}

// DeleteAdmin ...
func (a *Area) DeleteAdmin(aa *AreaAdmin) error {
	return errors.New("Not implemented yet")
}

///////////////////////////////////////////////////////////////////////////////////

// AreaModel ...
type AreaModel interface {
	NewArea() *Area
}

///////////////////////////////////////////////////////////////////////////////////

// NewArea ...
func (m *model) NewArea() *Area {
	a := Area{
		model: m,
	}

	return &a
}
