package config

import (
	"fmt"
	"go_todo/utils"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gopkg.in/go-ini/ini.v1"
)

type ConfigList struct {
	Port      string
	SQLDriver string
	DbName    string
	LogFile   string
	Static    string
	AppEnv    string
}

var Config ConfigList

func init() {
	LoadConfig()
	utils.Loggingsettings(Config.LogFile)
}

func LoadConfig() {
	cfg, err := ini.Load("config.ini")
	if err != nil {
		log.Fatalln(err)
	}
	if os.Getenv("APP_ENV") != "production" {
		err = godotenv.Load(fmt.Sprintf("./%s.env", os.Getenv("GO_ENV")))
		if err != nil {
			log.Fatalln(err)
		}
	}
	Config = ConfigList{
		Port:      cfg.Section("web").Key("Port").MustString("8080"),
		SQLDriver: cfg.Section("db").Key("driver").String(),
		DbName:    cfg.Section("db").Key("name").String(),
		LogFile:   cfg.Section("web").Key("logfile").String(),
		Static:    cfg.Section("web").Key("static").String(),
		AppEnv:    os.Getenv("APP_ENV"),
	}
}
