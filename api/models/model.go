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
	NotificationModel
}

// NewModel ...
func NewModel() Model {
	dbHost := beego.AppConfig.String("DB_HOST")
	dbPort := beego.AppConfig.String("DB_PORT")
	dbName := beego.AppConfig.String("DB_NAME")
	dbUser := beego.AppConfig.String("DB_USER")
	dbPassword := beego.AppConfig.String("DB_PASSWORD")

	dbDriver := "postgres"

	conn := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable",
		dbHost, dbPort, dbUser, dbName, dbPassword)

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
