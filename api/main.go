package main

import (
	"encoding/json"
	"flag"
	"os"

	"github.com/astaxie/beego"

	"github.com/mdiazp/sirel-server/api/app"
	_ "github.com/mdiazp/sirel-server/api/routers"
)

// SirelConfig ...
type SirelConfig struct {
	SIREL_PASSWORD string `json:"SIREL_PASSWORD"`

	DB_SOURCE_NAME string `json:"DB_SOURCE_NAME"`
	DB_USER        string `json:"DB_USER"`
	DB_PASSWORD    string `json:"DB_PASSWORD"`

	AdAddress  string `json:"AdAddress"`
	AdSuff     string `json:"AdSuff"`
	AdBDN      string `json:"AdBDN"`
	AdUser     string `json:"ad_user"`
	AdPassword string `json:"ad_password"`

	MailSenderUser     string
	MailSenderPassword string
	MailSenderHost     string
	MailSenderPort     string
}

func loadConfig() error {
	var configPath string
	flag.StringVar(&configPath, "configpath", "/home/kino/my_configs/sirel/config.json", "Path to config file.")
	flag.Parse()

	file, err := os.Open(configPath)
	if err != nil {
		beego.Critical("Cannot open config file: ", err)
		return err
	}
	defer file.Close()

	config := &SirelConfig{}

	//Parsing json file
	decoder := json.NewDecoder(file)
	err = decoder.Decode(config)
	if err != nil {
		beego.Critical("Cannot get configuration from file: ", err)
		return err
	}

	beego.AppConfig.Set("SIREL_PASSWORD", config.SIREL_PASSWORD)

	beego.AppConfig.Set("DB_SOURCE_NAME", config.DB_SOURCE_NAME)
	beego.AppConfig.Set("DB_USER", config.DB_USER)
	beego.AppConfig.Set("DB_PASSWORD", config.DB_PASSWORD)

	beego.AppConfig.Set("AdAddress", config.AdAddress)
	beego.AppConfig.Set("AdSuff", config.AdSuff)
	beego.AppConfig.Set("AdBDN", config.AdBDN)
	beego.AppConfig.Set("AdUser", config.AdUser)
	beego.AppConfig.Set("AdPassword", config.AdPassword)

	beego.AppConfig.Set("MailSenderUser", config.MailSenderUser)
	beego.AppConfig.Set("MailSenderPassword", config.MailSenderPassword)
	beego.AppConfig.Set("MailSenderHost", config.MailSenderHost)
	beego.AppConfig.Set("MailSenderPort", config.MailSenderPort)

	return nil
}

func main() {
	e := loadConfig()
	if e != nil {
		return
	}

	app.InitApp()

	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}

	beego.Run()
}
