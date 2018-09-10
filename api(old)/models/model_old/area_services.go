package models

type AreaServices interface {
	CreateArea(a Area) (Area, error)
	GetAreaById(id int) (Area, error)
	UpdateArea(a Area) (Area, error)
	DeleteArea(id int) error
	GetAreaQuerySeter() AreaQuerySeter
}
