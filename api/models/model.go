package models

import (
	"fmt"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/lib/pq"
)

var (
	ErrResultNotFound = orm.ErrNoRows
)

var Model orm.Ormer

func init() {
	orm.Debug = true

	orm.RegisterModel(
		new(User),
		new(Area),
		new(Local),
		new(Reservation),
		new(Notification),
	)
}

func NewModel() orm.Ormer {
	db_name := beego.AppConfig.String("DB_SOURCE_NAME")
	db_user := beego.AppConfig.String("DB_USER")
	db_password := beego.AppConfig.String("DB_PASSWORD")

	db_driver := "postgres"
	db_alias := "default"

	conn := fmt.Sprintf("user=%s dbname=%s password=%s sslmode=disable",
		db_user, db_name, db_password)

	orm.RegisterDriver(db_driver, orm.DRPostgres)
	orm.RegisterDataBase(db_alias, db_driver, conn)

	o := orm.NewOrm()
	return o
}
