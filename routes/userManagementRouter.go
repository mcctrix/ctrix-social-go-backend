package routes

import (
	"github.com/gofiber/fiber/v3"
	"github.com/mcctrix/ctrix-social-go-backend/controllers"
	"github.com/mcctrix/ctrix-social-go-backend/middleware"
)

func UserManagementRouter(router fiber.Router) {
	router.Use(middleware.AuthMiddleware())

	router.Get("/profile/:userid", controllers.GetUserProfileWithParam())
	router.Get("/profile", controllers.GetCurrentUserProfile())
	router.Patch("/profile", controllers.UpdateCurrentUserProfile())

	router.Get("/additional_info", controllers.GetAdditionalUserInfo())
	router.Patch("/additional_info", controllers.UpdateAdditionalUserInfo())

	router.Get("/user_settings", controllers.GetUserSettings())
	router.Patch("/user_settings", controllers.UpdateUserSettings())

	router.Get("/user_data", controllers.GetUserData())
	router.Patch("/user_data", controllers.UpdateUserData())
}
