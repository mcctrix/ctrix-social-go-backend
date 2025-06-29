package config

import (
	"os"

	"github.com/mcctrix/ctrix-social-go-backend/internal/pkg/utils"
)

type DatabaseConfig struct {
	Host     string
	Port     int
	User     string
	Password string
	Name     string
	SSLMode  string
}

var DBConfig *DatabaseConfig

func LoadDBConfig() *DatabaseConfig {
	currentEnv := os.Getenv("APP_ENV")
	if currentEnv == "dev" {
		DBConfig.Host = utils.GetEnv("postgresHostDev", "localhost")
		DBConfig.Name = utils.GetEnv("postgresDBDev", "")
		DBConfig.User = utils.GetEnv("postgresUsernameDev", "")
		DBConfig.Password = utils.GetEnv("postgresPasswordDev", "")
	}

	if currentEnv == "production" {
		DBConfig.Host = utils.GetEnv("postgresHostProd", "")
		DBConfig.Name = utils.GetEnv("postgresDBProd", "")
		DBConfig.User = utils.GetEnv("postgresUsernameProd", "")
		DBConfig.Password = utils.GetEnv("postgresPasswordProd", "")
	}
	return DBConfig
}
