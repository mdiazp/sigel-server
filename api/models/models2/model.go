package models2

import (
	"database/sql"

	"github.com/astaxie/beego/orm"
	"github.com/mdiazp/kmodel"
)

var (
	// ErrNoRows ...
	ErrNoRows = sql.ErrNoRows
)

// Model ...
type Model interface {
	orm.Ormer
	kmodel.Model
	LocalModel
	LocalAdminModel
	AreaModel
	AreaAdminModel
	UserModel
}

// NewModel ...
func NewModel(o orm.Ormer) Model {
	m := new(model)
	m.Ormer = o
	db, e := orm.GetDB()
	if e != nil {
		panic(e)
	}
	m.Model = kmodel.NewModel(db)
	return m
}

type model struct {
	orm.Ormer
	kmodel.Model
}
