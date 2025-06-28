package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"github.com/mcctrix/ctrix-social-go-backend/internal/infrastructure/database"
	"github.com/mcctrix/ctrix-social-go-backend/internal/pkg/auth"
)

func Load() {
	// Load the .env file in the current directory
	godotenv.Load()

	if _, err := os.Stat("./ecdsa_private_key.pem"); err != nil {
		auth.GenerateEcdsaPrivateKey()
	}
	_, err := database.DBConnection()
	if err != nil {
		fmt.Println("Error connecting to db: ", err)
		os.Exit(1)
	}
}
