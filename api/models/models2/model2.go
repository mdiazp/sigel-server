package models2

import (
	"database/sql"

	"github.com/astaxie/beego/orm"
)

// Model ...
type Model interface {
	orm.Ormer
	AreaModel
}

// NewModel ...
func NewModel(o orm.Ormer) Model {
	m := new(model)
	m.Ormer = o
	db, e := orm.GetDB()
	m.db = db

	if e != nil {
		panic(e)
	}

	return m
}

type model struct {
	orm.Ormer
	db *sql.DB
}
