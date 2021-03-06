package config

import (
	"os"
	"sync"

	"github.com/furqonzt99/refresh-token/constants"
	"github.com/joho/godotenv"
	"github.com/labstack/gommon/log"
)

type AppConfig struct {
	Port     string
	Database struct {
		Driver   string
		Name     string
		Host     string
		Port     string
		Username string
		Password string
	}
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

	err := godotenv.Load()

	if err != nil {
		log.Info("Error loading .env file")
	}

	var defaultConfig AppConfig
	defaultConfig.Port = os.Getenv("APP_PORT")
	defaultConfig.Database.Driver = os.Getenv("DB_DRIVER")
	defaultConfig.Database.Name = os.Getenv("DB_NAME")
	defaultConfig.Database.Host = os.Getenv("DB_HOST")
	defaultConfig.Database.Port = os.Getenv("DB_PORT")
	defaultConfig.Database.Username = os.Getenv("DB_USERNAME")
	defaultConfig.Database.Password = os.Getenv("DB_PASSWORD")

	constants.JWT_ACCESS_KEY = os.Getenv("JWT_ACCESS_KEY")
	constants.JWT_REFRESH_KEY = os.Getenv("JWT_REFRESH_KEY")

	constants.ADMIN_EMAIL = os.Getenv("ADMIN_EMAIL")
	constants.ADMIN_PASSWORD = os.Getenv("ADMIN_PASSWORD")

	return &defaultConfig
}
