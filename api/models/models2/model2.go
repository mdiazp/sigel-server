package models2

import (
	"database/sql"
	"strconv"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

///////////////////////////////////////////////////////////////////////////////////

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

///////////////////////////////////////////////////////////////////////////////////

// Model ...
type Model interface {
	orm.Ormer
	AreaModel
	LocalModel
	CRUD
}

// CRUD ...
type CRUD interface {
	Create2(o objectModel) error
	Retrieve2(o objectModel) error
	Update2(o objectModel) error
	Delete2(o objectModel) error
}

///////////////////////////////////////////////////////////////////////////////////

type objectModel interface {
	tableName() string

	autoPKey() bool
	pkeyName() string
	pkeyValue() interface{}
	pkeyPointer() interface{}

	columnNames(pk ...bool) []string
	columnValues(pk ...bool) []interface{}
	columnPointers(pk ...bool) []interface{}
}

///////////////////////////////////////////////////////////////////////////////////

type model struct {
	orm.Ormer
	db *sql.DB
}

// Create2 ...
func (m *model) Create2(o objectModel) error {
	cnames := o.columnNames(!o.autoPKey())
	qnums := ""
	ln := len(cnames)

	q := "INSERT INTO " + o.tableName() + " ("
	paramID := 1

	for i := 0; i < ln; i++ {
		q += cnames[i]
		qnums += "$" + strconv.Itoa(paramID)
		paramID++

		if i+1 < ln {
			q += ", "
			qnums += ", "
		}
	}

	q += ") VALUES (" + qnums + ")"
	q += " RETURNING " + fComa(o.columnNames(true)...)

	beego.Debug("q = ", q)

	stmt, e := m.db.Prepare(q)
	if e == nil {
		defer stmt.Close()
		e = stmt.QueryRow(o.columnValues(!o.autoPKey())...).
			Scan(o.columnPointers(true)...)
	}
	return e
}

// RetrieveByPK ...
func (m *model) Retrieve2(o objectModel) error {
	q := "SELECT " + fComa(o.columnNames(true)...)
	q += " FROM " + o.tableName() + " WHERE " + o.pkeyName() + "=$1"

	beego.Debug("q = ", q)

	stmt, e := m.db.Prepare(q)
	if e == nil {
		defer stmt.Close()
		e = stmt.QueryRow(o.pkeyValue()).
			Scan(o.columnPointers(true)...)
	}
	return e
}

// Update ...
func (m *model) Update2(o objectModel) error {
	q := "UPDATE " + o.tableName() + " SET "

	cnames := o.columnNames()
	ln := len(cnames)
	for i := 0; i < ln; i++ {
		q += cnames[i] + "=$" + strconv.Itoa(i+1)
		if i+1 < ln {
			q += ", "
		}
	}
	q += " WHERE " + o.pkeyName() + "=$" + strconv.Itoa(ln+1)
	q += " RETURNING " + fComa(o.columnNames(true)...)

	beego.Debug("q = ", q)

	stmt, e := m.db.Prepare(q)
	if e == nil {
		defer stmt.Close()
		e = stmt.QueryRow(o.columnValues(true)...).
			Scan(o.columnPointers(true)...)
	}
	return e
}

// Delete ...
func (m *model) Delete2(o objectModel) error {
	q := "DELETE FROM " + o.tableName() +
		" WHERE " + o.pkeyName() + "=$1"

	beego.Debug(q)

	stmt, e := m.db.Prepare(q)
	if e == nil {
		defer stmt.Close()
		_, e = stmt.Exec(o.pkeyValue())
	}
	return e
}

func fComa(x ...string) string {
	s := ""
	ln := len(x)
	for i := 0; i < ln; i++ {
		s += x[i]
		if i+1 < ln {
			s += ", "
		}
	}
	return s
}
