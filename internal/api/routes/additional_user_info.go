package routes

import (
	"github.com/gofiber/fiber/v3"
	addionalInfoUser "github.com/mcctrix/ctrix-social-go-backend/internal/api/handlers/addional_info_user"
	"github.com/mcctrix/ctrix-social-go-backend/internal/api/middleware"
	"github.com/mcctrix/ctrix-social-go-backend/internal/domain/services"
)

func AdditionalUserInfoRouter(router fiber.Router, services *services.Services) {
	router.Use(middleware.AuthMiddleware())

	additionalInfoUserHandler := addionalInfoUser.NewAdditionalInfoUserHandler(services.AdditionalUserInfoService)

	router.Get("/", additionalInfoUserHandler.GetAdditionalInfoUser)
	router.Patch("/", additionalInfoUserHandler.UpdateAdditionalInfoUser)

}
