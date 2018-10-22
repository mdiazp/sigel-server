package models

import (
	"database/sql"
	"fmt"

	"github.com/astaxie/beego"
	"github.com/mdiazp/kmodel"

	// Postgresql Driver
	_ "github.com/lib/pq"
)

var (
	// ErrNoRows ...
	ErrNoRows = sql.ErrNoRows
)

// Model ...
type Model interface {
	kmodel.Model
	LocalModel
	LocalAdminModel
	AreaModel
	UserModel
	ReservationModel
}

// NewModel ...
func NewModel() Model {
	dbName := beego.AppConfig.String("DB_SOURCE_NAME")
	dbUser := beego.AppConfig.String("DB_USER")
	dbPassword := beego.AppConfig.String("DB_PASSWORD")

	dbDriver := "postgres"

	conn := fmt.Sprintf("user=%s dbname=%s password=%s sslmode=disable",
		dbUser, dbName, dbPassword)

	db, e := sql.Open(dbDriver, conn)
	if e != nil {
		panic(e)
	}

	return &model{
		Model: kmodel.NewModel(db),
	}
}

type model struct {
	kmodel.Model
}
