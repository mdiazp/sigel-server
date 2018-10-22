package main

import (
	"encoding/json"
	"flag"
	"os"

	"github.com/astaxie/beego"

	"github.com/mdiazp/sirel-server/api/app"
	_ "github.com/mdiazp/sirel-server/api/routers"
)

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
}

func loadConfig() error {
	var config_path string
	flag.StringVar(&config_path, "configpath", "/home/kino/my_configs/sirel/config.json", "Path to config file.")
	flag.Parse()

	file, err := os.Open(config_path)
	if err != nil {
		beego.Critical("Cannot open config file: ", err)
		return err
	}
	defer file.Close()

	cred := &SirelConfig{}

	//Parsing json file
	decoder := json.NewDecoder(file)
	err = decoder.Decode(cred)
	if err != nil {
		beego.Critical("Cannot get configuration from file: ", err)
		return err
	}

	beego.AppConfig.Set("SIREL_PASSWORD", cred.SIREL_PASSWORD)

	beego.AppConfig.Set("DB_SOURCE_NAME", cred.DB_SOURCE_NAME)
	beego.AppConfig.Set("DB_USER", cred.DB_USER)
	beego.AppConfig.Set("DB_PASSWORD", cred.DB_PASSWORD)

	beego.AppConfig.Set("AdAddress", cred.AdAddress)
	beego.AppConfig.Set("AdSuff", cred.AdSuff)
	beego.AppConfig.Set("AdBDN", cred.AdBDN)
	beego.AppConfig.Set("AdUser", cred.AdUser)
	beego.AppConfig.Set("AdPassword", cred.AdPassword)

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
