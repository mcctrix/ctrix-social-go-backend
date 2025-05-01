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
	router.Get("/profile/:userid", controllers.GetUserProfileWithParam())

	router.Get("/additional_info", controllers.GetAdditionalUserInfo())
	router.Post("/additional_info", controllers.CreateAdditionalUserInfo())
	router.Patch("/additional_info", controllers.UpdateAdditionalUserInfo())

	router.Get("/user_settings", controllers.GetUserSettings())
	router.Post("/user_settings", controllers.CreateUserSettings())
	router.Patch("/user_settings", controllers.UpdateUserSettings())

	router.Get("/user_data", controllers.GetUserData())
	router.Post("/user_data", controllers.SetUserData())

}
