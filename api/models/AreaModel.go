package models

import (
	"fmt"

	"github.com/mdiazp/kmodel"
)

///////////////////////////////////////////////////////////////////////////////////

// AreaInfo ...
type AreaInfo struct {
	ID          int
	Name        string
	Description string
	Location    string
}

// Area ...
type Area struct {
	AreaInfo
	model Model

	locals *[]*Local
}

/////////////////////////////////////////////////////

// TableName ...
func (a *Area) TableName() string {
	return "area"
}

// AutoPKey ...
func (a *Area) AutoPKey() bool {
	return true
}

// PkeyName ...
func (a *Area) PkeyName() string {
	return "id"
}

// PkeyValue ...
func (a *Area) PkeyValue() interface{} {
	return a.ID
}

// PkeyPointer ...
func (a *Area) PkeyPointer() interface{} {
	return &a.ID
}

// ColumnNames ...
func (a *Area) ColumnNames() []string {
	return []string{
		"name",
		"description",
		"location",
	}
}

// ColumnValues ...
func (a *Area) ColumnValues() []interface{} {
	return []interface{}{
		a.Name,
		a.Description,
		a.Location,
	}
}

// ColumnPointers ...
func (a *Area) ColumnPointers() []interface{} {
	return []interface{}{
		&a.Name,
		&a.Description,
		&a.Location,
	}
}

/////////////////////////////////////////////////////

// Update ...
func (a *Area) Update() error {
	return a.model.Update(a)
}

// Load ...
func (a *Area) Load() error {
	return a.model.Retrieve(a)
}

// Locals ...
func (a *Area) Locals() (*[]*Local, error) {
	var e error
	if a.locals == nil {
		tmp := a.model.NewLocalCollection()
		hfilter := fmt.Sprintf("area_id=%d", a.ID)
		e = a.model.RetrieveCollection(&hfilter, nil, nil, nil, nil, tmp)
		if e == nil {
			a.locals = tmp.Locals
		}
	}
	return a.locals, e
}

///////////////////////////////////////////////////////////////////////////////////

// AreaCollection ...
type AreaCollection struct {
	model Model
	Areas *[]*Area
}

// NewObjectModel ...
func (c *AreaCollection) NewObjectModel() kmodel.ObjectModel {
	return c.model.NewArea()
}

// Add ...
func (c *AreaCollection) Add() kmodel.ObjectModel {
	a := c.model.NewArea()
	*(c.Areas) = append(*(c.Areas), a)
	return a
}

///////////////////////////////////////////////////////////////////////////////////

// AreaModel ...
type AreaModel interface {
	NewArea() *Area
	NewAreaCollection() *AreaCollection
	Areas(limit, offset *int, orderby *string,
		orderDesc *bool) (*AreaCollection, error)

	AreaCustomModel
}

/////////////////////////////////////////////////////

// NewArea ...
func (m *model) NewArea() *Area {
	a := &Area{
		model: m,
	}
	return a
}

// NewAreaCollection ...
func (m *model) NewAreaCollection() *AreaCollection {
	kk := make([]*Area, 0)
	return &AreaCollection{
		model: m,
		Areas: &kk,
	}
}

func (m *model) Areas(limit, offset *int, orderby *string,
	orderDesc *bool) (*AreaCollection, error) {

	collection := m.NewAreaCollection()
	e := m.RetrieveCollection(nil, limit, offset, orderby, orderDesc, collection)
	return collection, e
}
