package config

import (
	"os"

	"github.com/joho/godotenv"
	"github.com/mcctrix/ctrix-social-go-backend/internal/pkg/auth"
	"github.com/mcctrix/ctrix-social-go-backend/internal/pkg/utils"
)

type loadConfigs struct {
	Port     string
	Host     string
	DBConfig *DatabaseConfig
}

var configs *loadConfigs

func LoadConfig() *loadConfigs {
	// Load the .env file in the current directory
	godotenv.Load()

	if _, err := os.Stat("./ecdsa_private_key.pem"); err != nil {
		auth.GenerateEcdsaPrivateKey()
	}
	configs.DBConfig = LoadDBConfig()
	configs.Port = utils.GetEnv("Server_Port", "4000")
	configs.Host = utils.GetEnv("Server_Host", "localhost")
	return configs
}
