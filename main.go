package main

import (
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/mcctrix/ctrix-social-go-backend/routes"
)

func main() {
	// Load the .env file in the current directory
	godotenv.Load()

	port := os.Getenv("PORT")

	if port == "" {
		port = "4000"
	}

	router := gin.New()
	router.SetTrustedProxies([]string{"localhost", "127.0.0.1", "0.0.0.0"})

	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000", "http://127.0.0.1:3000"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders:     []string{"Origin", "Content-type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))

	routes.AuthRouter(router)

	router.Run()
}
