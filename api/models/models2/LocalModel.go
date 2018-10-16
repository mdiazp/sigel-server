package models2

import "errors"

///////////////////////////////////////////////////////////////////////////////////

// Local ...
type Local struct {
	ID                      int    `json:"id"`
	AreaID                  int    `json:"area_id" valid:"Required"`
	Name                    string `json:"name" valid:"Required;MaxSize(100)"`
	Description             string `json:"description" valid:"Required;MaxSize(1024)"`
	Location                string `json:"location" valid:"Required;MaxSize(1024)"`
	WorkingMonths           string `json:"working_months" valid:"Required;MinSize(12);MaxSize(12)"`
	WorkingWeekDays         string `json:"working_week_days" valid:"Required;MinSize(7);MaxSize(7)"`
	WorkingBeginTimeHours   int    `json:"working_begin_time_hours" valid:"Min(0);Max(23)"`
	WorkingBeginTimeMinutes int    `json:"working_begin_time_minutes" valid:"Min(0);Max(59)"`
	WorkingEndTimeHours     int    `json:"working_end_time_hours" valid:"Min(0);Max(23)"`
	WorkingEndTimeMinutes   int    `json:"working_end_time_minutes" valid:"Min(0);Max(59)"`
	EnableToReserve         bool   `json:"enable_to_reserve"`

	model *model `json:"-"`

	area       *Area `json:"-"`
	lzLoadArea bool  `json:"-"`

	admins       []*AreaAdmin `json:"-"`
	lzLoadAdmins bool         `json:"-"`
}

// Update ...
func (lo *Local) Update() error {
	return errors.New("Not implemented yet")
}

// Area ...
func (lo *Local) Area() (*Area, error) {
	return nil, errors.New("Not implemented yet")
}

// Admins ...
func (lo *Local) Admins() ([]*LocalAdmin, error) {
	return nil, errors.New("Not implemented yet")
}

// AddAdmin ...
func (lo *Local) AddAdmin(admin *LocalAdmin) error {
	return errors.New("Not implemented yet")
}

// DeleteAdmin ...
func (lo *Local) DeleteAdmin(admin *LocalAdmin) error {
	return errors.New("Not implemented yet")
}

///////////////////////////////////////////////////////////////////////////////////

// LocalModel ...
type LocalModel interface {
	NewLocal() Local
	CreateLocal(*Local) error
	RetrieveLocal(*Local) error
	UpdateLocal(*Local) error
	DeleteLocal(*Local) error
}

///////////////////////////////////////////////////////////////////////////////////

// NewLocal ...
func (m *model) NewLocal() Local {
	lo := Local{
		model: m,
	}
	return lo
}

// CreateLocal ...
func (m *model) CreateLocal(lo *Local) error {
	return errors.New("Not implemented yet")
}

// RetrieveLocal ...
func (m *model) RetrieveLocal(lo *Local) error {
	return errors.New("Not implemented yet")
}

// UpdateLocal ...
func (m *model) UpdateLocal(lo *Local) error {
	return errors.New("Not implemented yet")
}

// DeleteLocal ...
func (m *model) DeleteLocal(lo *Local) error {
	return errors.New("Not implemented yet")
}

///////////////////////////////////////////////////////////////////////////////////
