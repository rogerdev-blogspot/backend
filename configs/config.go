package configs

import (
	"os"
	"strconv"
	"sync"

	"github.com/labstack/gommon/log"
)

type AppConfig struct {
	Port     int
	Database struct {
		Driver   string
		Name     string
		Address  string
		Port     string
		Username string
		Password string
	}
	GoogleClientID     string
	GoogleClientSecret string
}

var lock = &sync.Mutex{}
var appConfig *AppConfig

func GetConfig() *AppConfig {
	lock.Lock()
	defer lock.Unlock()

	if appConfig == nil {
		appConfig = initConfig()
	}

	return appConfig
}

func initConfig() *AppConfig {

	portDB := os.Getenv("DB_PORT")

	var defaultConfig AppConfig
	port, errPort := strconv.Atoi(os.Getenv("APP_PORT"))
	if errPort != nil {
		log.Warn(errPort.Error())
	}

	defaultConfig.Port = port
	defaultConfig.Database.Driver = os.Getenv("DB_DRIVER")
	defaultConfig.Database.Name = os.Getenv("DB_NAME")
	defaultConfig.Database.Address = os.Getenv("DB_ADDRESS")
	defaultConfig.Database.Port = portDB
	defaultConfig.Database.Username = os.Getenv("DB_USERNAME")
	defaultConfig.Database.Password = os.Getenv("DB_PASSWORD")

	defaultConfig.GoogleClientID = os.Getenv("GOOGLE_CLIENT_ID")
	defaultConfig.GoogleClientSecret = os.Getenv("GOOGLE_CLIENT_SECRET")

	return &defaultConfig
}
