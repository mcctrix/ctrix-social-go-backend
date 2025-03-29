package main

import (
	"os"

	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/cors"
	"github.com/gofiber/fiber/v3/middleware/logger"
	recoverer "github.com/gofiber/fiber/v3/middleware/recover"
	"github.com/joho/godotenv"
	"github.com/mcctrix/ctrix-social-go-backend/db"
	"github.com/mcctrix/ctrix-social-go-backend/routes"
	"github.com/mcctrix/ctrix-social-go-backend/utils"
)

func main() {

	loadEnvironment()

	port := os.Getenv("PORT")

	if port == "" {
		port = ":4000"
	}

	mainRouter := makeRouter()

	mainRouter.Get("/", func(c fiber.Ctx) error {
		return c.SendString("Hello world fiber v3")
	})

	routes.AuthRouter(mainRouter.Group("/api/auth"))
	routes.UserManagementRouter(mainRouter.Group("/api/users"))
  routes.PostManagementRouter(mainRouter.Group("/api/post"))

	mainRouter.Listen(port)
}

func loadEnvironment() {
	// Load the .env file in the current directory
	godotenv.Load()

	if _, err := os.Stat("./ecdsa_private_key.pem"); err == nil {
		// Pem File exist so do nothing
	} else {
		utils.GenerateEcdsaPrivateKey()
  }
  // db.ResetDB()
	db.CreateInitialDBStructure()
}

func makeRouter() *fiber.App {
	router := fiber.New(fiber.Config{
		TrustProxy: true,
		TrustProxyConfig: fiber.TrustProxyConfig{
			Proxies:  []string{"127.0.0.1", "0.0.0.0"},
			Loopback: true,
		},
	})
	router.Use(recoverer.New())

	router.Use(logger.New(logger.Config{
		Format: "[${ip}]: ${port} ${status} - ${method} ${path}\n",
	}))

	corsMiddleware := cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000", "http://127.0.0.1:3000"}, // Allows all origins
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization", "Set-Cookie"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	})

	router.Use(corsMiddleware)
	return router
}
