package config

import (
	"fmt"
	"os"
	"strconv"

	"github.com/mcctrix/ctrix-social-go-backend/internal/pkg/utils"
)

type DatabaseConfig struct {
	Host     string
	Port     int
	User     string
	Password string
	Name     string
}

var DBConfig *DatabaseConfig = &DatabaseConfig{}

func LoadDBConfig() *DatabaseConfig {
	currentEnv := os.Getenv("APP_ENV")
	if currentEnv == "dev" {
		DBConfig.Host = utils.GetEnv("postgresHostDev", "localhost")
		DBConfig.Name = utils.GetEnv("postgresDBDev", "")
		DBConfig.User = utils.GetEnv("postgresUsernameDev", "")
		DBConfig.Password = utils.GetEnv("postgresPasswordDev", "")
		port, err := strconv.Atoi(utils.GetEnv("postgresPortDev", "5432"))
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		DBConfig.Port = port
	}

	if currentEnv == "production" {
		DBConfig.Host = utils.GetEnv("postgresHostProd", "")
		DBConfig.Name = utils.GetEnv("postgresDBProd", "")
		DBConfig.User = utils.GetEnv("postgresUsernameProd", "")
		DBConfig.Password = utils.GetEnv("postgresPasswordProd", "")
		port, err := strconv.Atoi(utils.GetEnv("postgresPortProd", "5432"))
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		DBConfig.Port = port
	}
	fmt.Println(DBConfig)
	return DBConfig
}
