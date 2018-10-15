package models2

import "errors"

///////////////////////////////////////////////////////////////////////////////////

// Area ...
type Area struct {
	ID          int    `json:"id"`
	Name        string `json:"name" valid:"Required;MaxSize(100)"`
	Description string `json:"description" valid:"Required;MaxSize(1024)"`
	Location    string `json:"locatiion" valid:"Required;MaxSize(1024)"`

	areaActions
}

///////////////////////////////////////////////////////////////////////////////////

type areaActions interface {
	Update() error

	Locals() ([]*Local, error)
	AddLocal(*Local) error
	DeleteLocal(*Local) error

	Admins() ([]*AreaAdmin, error)
	AddAdmin(*AreaAdmin) error
	DeleteAdmin(*AreaAdmin) error
}

type area struct {
	*Area
	*model

	locals       []*Local
	lzLoadLocals bool

	admins       []*AreaAdmin
	lzLoadAdmins bool
}

func (a *area) Update() error {
	return errors.New("Not implemented yet")
}

func (a *area) Locals() ([]*Local, error) {
	return nil, errors.New("Not implemented yet")
}

func (a *area) AddLocal(lo *Local) error {
	return errors.New("Not implemented yet")
}

func (a *area) DeleteLocal(lo *Local) error {
	return errors.New("Not implemented yet")
}

func (a *area) Admins() ([]*AreaAdmin, error) {
	return nil, errors.New("Not implemented yet")
}

func (a *area) AddAdmin(aa *AreaAdmin) error {
	return errors.New("Not implemented yet")
}

func (a *area) DeleteAdmin(aa *AreaAdmin) error {
	return errors.New("Not implemented yet")
}

func (m *model) newArea() Area {
	a := Area{}
	a.areaActions = &area{
		Area: &a,
	}

	return a
}

///////////////////////////////////////////////////////////////////////////////////

// AreaModel ...
type AreaModel interface {
	InsertArea(Area) (Area, error)
	DeleteArea(Area) error
	UpdateArea(Area) (Area, error)
}

///////////////////////////////////////////////////////////////////////////////////

// InsertArea ...
func (m *model) InsertArea(a Area) (Area, error) {
	return Area{}, errors.New("Not implemented yet")
}

// DeleteArea ...
func (m *model) DeleteArea(a Area) error {
	return errors.New("Not implemented yet")
}

// UpdateArea ...
func (m *model) UpdateArea(a Area) (Area, error) {
	return Area{}, errors.New("Not implemented yet")
}
