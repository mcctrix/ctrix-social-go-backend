package server

import (
	"github.com/gofiber/fiber/v3"
	"github.com/mcctrix/ctrix-social-go-backend/internal/api/middleware"
	"github.com/mcctrix/ctrix-social-go-backend/internal/api/routes"
)

// NewServer creates and configures a new Fiber application.
func NewServer() *fiber.App {
	router := fiber.New(fiber.Config{
		TrustProxy: true,
		TrustProxyConfig: fiber.TrustProxyConfig{
			Proxies:  []string{"127.0.0.1", "0.0.0.0"},
			Loopback: true,
		},
		BodyLimit: 25 * 1024 * 1024,
	})

	// Global Middleware
	router.Use(middleware.RecovererMiddleware())
	router.Use(middleware.LoggerMiddleware())
	router.Use(middleware.SecurityMiddleware())
	router.Use(middleware.CORSMiddleware())

	// Base route
	router.Get("/", func(c fiber.Ctx) error {
		return c.SendString("This is backend of Ctrix Social App!")
	})

	// API Routes
	routes.SetupRoutes("/api", router)

	return router
}
