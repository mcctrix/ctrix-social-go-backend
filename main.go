package main

import (
	"os"

	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/cors"
	"github.com/gofiber/fiber/v3/middleware/logger"
	recoverer "github.com/gofiber/fiber/v3/middleware/recover"
	"github.com/joho/godotenv"
	"github.com/mcctrix/ctrix-social-go-backend/routes"
)

func main() {
	// Load the .env file in the current directory
	godotenv.Load()

	port := os.Getenv("PORT")

	if port == "" {
		port = ":4000"
	}

	mainRouter := fiber.New(fiber.Config{
		TrustProxy: true,
		TrustProxyConfig: fiber.TrustProxyConfig{
			Proxies:  []string{"localhost", "127.0.0.1", "0.0.0.0"},
			Loopback: true,
		},
	})

	mainRouter.Use(recoverer.New())

	mainRouter.Use(logger.New(logger.Config{
		Format: "[${ip}]: ${port} ${status} - ${method} ${path}\n",
	}))

	corsMiddleware := cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000", "http://127.0.0.1:3000"}, // Allows all origins
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	})

	mainRouter.Use(corsMiddleware)

	mainRouter.Get("/", func(c fiber.Ctx) error {
		return c.SendString("Hello world fiber v3")
	})

	routes.AuthRouter(mainRouter.Group("/api/auth"))

	mainRouter.Listen(port)
}
