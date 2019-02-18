package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"time"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"

	"github.com/mdiazp/sigel-server/api/app"
	_ "github.com/mdiazp/sigel-server/api/routers"
)

// SirelConfig ...
type SirelConfig struct {
	SIREL_PASSWORD string `json:"SIREL_PASSWORD"`

	DB_HOST     string
	DB_PORT     string
	DB_NAME     string
	DB_USER     string `json:"DB_USER"`
	DB_PASSWORD string `json:"DB_PASSWORD"`

	AdAddress  string `json:"AdAddress"`
	AdSuff     string `json:"AdSuff"`
	AdBDN      string `json:"AdBDN"`
	AdUser     string `json:"ad_user"`
	AdPassword string `json:"ad_password"`

	MailSenderUser     string
	MailSenderPassword string
	MailSenderHost     string
	MailSenderPort     string

	LoggingFilePath string

	PublicStaticFilesPath string
}

func loadConfig() error {
	var configPath string
	flag.StringVar(&configPath, "configpath", "/home/kino/my_configs/sigel/config.json", "Path to config file.")
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

	beego.AppConfig.Set("DB_HOST", config.DB_HOST)
	beego.AppConfig.Set("DB_PORT", config.DB_PORT)
	beego.AppConfig.Set("DB_NAME", config.DB_NAME)
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

	beego.AppConfig.Set("LoggingFilePath", config.LoggingFilePath)

	beego.AppConfig.Set("PublicStaticFilesPath", config.PublicStaticFilesPath)

	return nil
}

func main() {
	e := loadConfig()
	if e != nil {
		return
	}

	app.InitApp()

	if beego.AppConfig.String("SERVE_SWAGGER") == "true" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}

	/************************************************************/
	// Logging File
	tim := time.Now()
	pln := fmt.Sprintf("%d-%02d-%02dT%02d-%02d-%02d",
		tim.Year(), tim.Month(), tim.Day(),
		tim.Hour(), tim.Minute(), tim.Second())
	beego.SetLogger(
		"file",
		fmt.Sprintf(
			`{"filename":"%s"}`,
			beego.AppConfig.String("LoggingFilePath")+"/"+pln+"-upr-sigel.log",
		),
	)
	beego.BConfig.Log.AccessLogs = true
	/*beego.BeeLogger.DelLogger("console")*/
	logs.EnableFuncCallDepth(true)

	beego.Run()
}
