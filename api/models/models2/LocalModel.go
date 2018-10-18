package models2

import (
	"fmt"

	"github.com/mdiazp/kmodel"
)

///////////////////////////////////////////////////////////////////////////////////

// Local ...
type Local struct {
	ID                      int
	AreaID                  int
	Name                    string
	Description             string
	Location                string
	WorkingMonths           string
	WorkingWeekDays         string
	WorkingBeginTimeHours   int
	WorkingBeginTimeMinutes int
	WorkingEndTimeHours     int
	WorkingEndTimeMinutes   int
	EnableToReserve         bool
	model                   Model

	localAdmins *[]*LocalAdmin
	area        *Area
}

/////////////////////////////////////////////////////

// TableName ...
func (l *Local) TableName() string {
	return "local"
}

// AutoPKey ...
func (l *Local) AutoPKey() bool {
	return true
}

// PkeyName ...
func (l *Local) PkeyName() string {
	return "id"
}

// PkeyValue ...
func (l *Local) PkeyValue() interface{} {
	return l.ID
}

// PkeyPointer ...
func (l *Local) PkeyPointer() interface{} {
	return &l.ID
}

// ColumnNames ...
func (l *Local) ColumnNames() []string {
	return []string{
		"area_id",
		"name",
		"description",
		"location",
		"working_months",
		"working_week_days",
		"working_begin_time_hours",
		"working_begin_time_minutes",
		"working_end_time_hours",
		"working_end_time_minutes",
		"enable_to_reserve",
	}
}

// ColumnValues ...
func (l *Local) ColumnValues() []interface{} {
	return []interface{}{
		l.AreaID,
		l.Name,
		l.Description,
		l.Location,
		l.WorkingMonths,
		l.WorkingWeekDays,
		l.WorkingBeginTimeHours,
		l.WorkingBeginTimeMinutes,
		l.WorkingEndTimeHours,
		l.WorkingEndTimeMinutes,
		l.EnableToReserve,
	}
}

// ColumnPointers ...
func (l *Local) ColumnPointers() []interface{} {
	return []interface{}{
		&l.AreaID,
		&l.Name,
		&l.Description,
		&l.Location,
		&l.WorkingMonths,
		&l.WorkingWeekDays,
		&l.WorkingBeginTimeHours,
		&l.WorkingBeginTimeMinutes,
		&l.WorkingEndTimeHours,
		&l.WorkingEndTimeMinutes,
		&l.EnableToReserve,
	}
}

/////////////////////////////////////////////////////

// Update ...
func (l *Local) Update() error {
	return l.model.Update2(l)
}

// Load ...
func (l *Local) Load() error {
	return l.model.Retrieve(l)
}

// Area ...
func (l *Local) Area() (*Area, error) {
	var e error
	if l.area == nil {
		l.area = l.model.NewArea()
		l.area.ID = l.AreaID
		e = l.model.Retrieve(l.area)
	}
	return l.area, e
}

// LocalAdmins ...
func (l *Local) LocalAdmins() (*[]*LocalAdmin, error) {
	var e error
	if l.localAdmins == nil {
		tmp := l.model.NewLocalAdminCollection()
		hfilter := fmt.Sprintf("local_id=%d", l.ID)
		e = l.model.RetrieveCollection(&hfilter, nil, nil, nil, nil, tmp)
		if e == nil {
			l.localAdmins = tmp.LocalAdmins
		}
	}
	return l.localAdmins, e
}

///////////////////////////////////////////////////////////////////////////////////

// LocalCollection ...
type LocalCollection struct {
	model  Model
	Locals *[]*Local
}

// NewObjectModel ...
func (c *LocalCollection) NewObjectModel() kmodel.ObjectModel {
	return c.model.NewLocal()
}

// Add ...
func (c *LocalCollection) Add() kmodel.ObjectModel {
	l := c.model.NewLocal()
	*(c.Locals) = append(*(c.Locals), l)
	return l
}

///////////////////////////////////////////////////////////////////////////////////

// LocalModel ...
type LocalModel interface {
	NewLocal() *Local
	NewLocalCollection() *LocalCollection
	Locals(limit, offset *int, orderby *string,
		orderDesc *bool) (*LocalCollection, error)
}

/////////////////////////////////////////////////////

// NewLocal ...
func (m *model) NewLocal() *Local {
	l := &Local{
		model: m,
	}
	return l
}

// NewLocalCollection ...
func (m *model) NewLocalCollection() *LocalCollection {
	kk := make([]*Local, 0)
	return &LocalCollection{
		model:  m,
		Locals: &kk,
	}
}

func (m *model) Locals(limit, offset *int, orderby *string,
	orderDesc *bool) (*LocalCollection, error) {

	collection := m.NewLocalCollection()
	e := m.RetrieveCollection(nil, limit, offset, orderby, orderDesc, collection)
	return collection, e
}
