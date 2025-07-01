package routes

import (
	"github.com/gofiber/fiber/v3"
	"github.com/mcctrix/ctrix-social-go-backend/internal/api/handlers/users"
	"github.com/mcctrix/ctrix-social-go-backend/internal/api/middleware"
	"github.com/mcctrix/ctrix-social-go-backend/internal/domain/services"
)

func UserRouter(router fiber.Router, services *services.Services) {
	router.Use(middleware.AuthMiddleware())

	profileHandler := users.NewProfileHandler(services.ProfileService)

	router.Get("/", profileHandler.GetCurrentUserProfile)
	router.Patch("/", profileHandler.UpdateCurrentUserProfile)

	router.Get("/user/:userid", profileHandler.GetUserProfileWithParam)
}
