package routes

import (
	"github.com/gofiber/fiber/v3"
	"github.com/mcctrix/ctrix-social-go-backend/internal/api/handlers/users"
	"github.com/mcctrix/ctrix-social-go-backend/internal/api/middleware"
	"github.com/mcctrix/ctrix-social-go-backend/internal/domain/services"
)

func UserRouter(router fiber.Router, services *services.Services) {
	router.Use(middleware.AuthMiddleware())

	router.Get("/", users.GetCurrentUserProfile())
	router.Patch("/", users.UpdateCurrentUserProfile())

	router.Get("/user/:userid", users.GetUserProfileWithParam())

	router.Get("/additional_info", users.GetAdditionalUserInfo())
	router.Patch("/additional_info", users.UpdateAdditionalUserInfo())

	router.Get("/user_settings", users.GetUserSettings())
	router.Patch("/user_settings", users.UpdateUserSettings())

	router.Post("/profile-setup", users.ProfileSetup())
}
