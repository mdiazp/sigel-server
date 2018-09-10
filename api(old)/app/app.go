package app

import (
	"github.com/astaxie/beego"
	"gitlab.com/manuel.diaz/sirel/server/api/models"
)

func InitApp() {
	db_name := beego.AppConfig.String("DB_SOURCE_NAME")
	db_user := beego.AppConfig.String("DB_USER")
	db_password := beego.AppConfig.String("DB_PASSWORD")
}

func Model() models.Model {
	return m
}
