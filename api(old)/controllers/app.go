package controllers

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"

	"github.com/astaxie/beego"
	"gitlab.com/manuel.diaz/sirel/server/api/models"
	"gitlab.com/manuel.diaz/sirel/server/api/models/bo"
)

var AppModel models.Model

func InitAppModel() {
	db_name := beego.AppConfig.String("DB_SOURCE_NAME")
	db_user := beego.AppConfig.String("DB_USER")
	db_password := beego.AppConfig.String("DB_PASSWORD")

	conn := fmt.Sprintf("user=%s dbname=%s password=%s sslmode=disable",
		db_user, db_name, db_password)
	db, e := sql.Open("postgres", conn)

	if e != nil {
		beego.Critical(fmt.Sprintf("Error opening database: %s", e.Error()))
		panic(e.Error())
	}

	AppModel, e = bo.NewModel(db)
	if e != nil {
		beego.Critical(fmt.Sprintf("Error creating model: %s", e.Error()))
		panic(e.Error())
	}
}
