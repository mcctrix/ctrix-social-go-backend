package routes

import (
	"github.com/gofiber/fiber/v3"
	userSetting "github.com/mcctrix/ctrix-social-go-backend/internal/api/handlers/user_settings"
	"github.com/mcctrix/ctrix-social-go-backend/internal/api/middleware"
	"github.com/mcctrix/ctrix-social-go-backend/internal/domain/services"
)

func UserSettingsRouter(router fiber.Router, services *services.Services) {
	router.Use(middleware.AuthMiddleware())

	userSettingsHandler := userSetting.NewUserSettingsHandler(services.UserSettingsService)

	router.Get("/", userSettingsHandler.GetUserSettings)
	router.Patch("/", userSettingsHandler.UpdateUserSettings)
}
