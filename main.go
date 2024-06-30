package main

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/mcctrix/ctrix-social-go-backend/routes"
)

func main() {
	port := os.Getenv("PORT")

	if port == "" {
		port = "4000"
	}

	router := gin.New()
	router.SetTrustedProxies([]string{"localhost"})

	router.Use(gin.Logger())
	router.Use(gin.Recovery())
	routes.AuthRouter(router)

	router.Run()
}
