package main

import (
	"fmt"
	"os"

	"github.com/gofiber/fiber/v3"

	recoverer "github.com/gofiber/fiber/v3/middleware/recover"
	"github.com/joho/godotenv"
	"github.com/mcctrix/ctrix-social-go-backend/internal/api/middleware"
	"github.com/mcctrix/ctrix-social-go-backend/internal/api/routes"
	db "github.com/mcctrix/ctrix-social-go-backend/internal/infrastructure/database"
	"github.com/mcctrix/ctrix-social-go-backend/internal/pkg/auth"
	"github.com/mcctrix/ctrix-social-go-backend/internal/pkg/utils"
)

func main() {

	loadEnvironment()
	CheckArgs()

	port := os.Getenv("PORT")

	if port == "" {
		port = "4000"
	}
	mainRouter := makeRouter()

	mainRouter.Get("/", func(c fiber.Ctx) error {
		return c.SendString("This is backend of Ctrix Social App!")
	})

	// Middleware
	mainRouter.Use(middleware.LoggerMiddleware())
	mainRouter.Use(middleware.CORSMiddleware())

	routes.SetupRoutes("/api", mainRouter)

	err := mainRouter.Listen(":" + port)
	if err != nil {
		fmt.Println(err)
	}
}

func CheckArgs() {
	if len(os.Args) == 0 {
		return
	}
	if utils.ContainsString(os.Args, "reset") {
		db.ResetDB()
	}
	if utils.ContainsString(os.Args, "init-db") {
		db.CreateInitialDBStructure()
	}
	if utils.ContainsString(os.Args, "populate-db") {
		db.PopulateDB()
		os.Exit(0)
	}
}

func loadEnvironment() {
	// Load the .env file in the current directory
	godotenv.Load()

	if _, err := os.Stat("./ecdsa_private_key.pem"); err != nil {
		auth.GenerateEcdsaPrivateKey()
	}
	_, err := db.DBConnection()
	if err != nil {
		fmt.Println("Error connecting to db: ", err)
		os.Exit(1)
	}
}

func makeRouter() *fiber.App {
	router := fiber.New(fiber.Config{
		TrustProxy: true,
		TrustProxyConfig: fiber.TrustProxyConfig{
			Proxies:  []string{"127.0.0.1", "0.0.0.0"},
			Loopback: true,
		},
		BodyLimit: 25 * 1024 * 1024,
	})
	router.Use(recoverer.New())

	return router
}
