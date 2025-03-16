package routes

import (
	"github.com/gofiber/fiber/v3"
	"github.com/mcctrix/ctrix-social-go-backend/controllers"
	"github.com/mcctrix/ctrix-social-go-backend/middleware"
)

func UserManagementRouter(router fiber.Router) {
	router.Use(middleware.AuthMiddleware())
	router.Get("/profile", controllers.GetCurrentUserProfile())
	router.Post("/profile", controllers.SetCurrentUserProfile())
}
