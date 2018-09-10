package models

type LocalServices interface {
	CreateLocal(l Local) (Local, error)
	GetLocalById(id int) (Local, error)
	UpdateLocal(l Local) (Local, error)
	DeleteLocal(id int) error
	GetLocalQuerySeter() LocalQuerySeter
}
