package routes

import (
	"github.com/gofiber/fiber/v3"
	"github.com/mcctrix/ctrix-social-go-backend/controllers"
)

func AuthRouter(router fiber.Router) {
	router.Post("/login", controllers.Login())
	router.Post("/signup", controllers.SignUp())
	router.Post("/logout", controllers.Logout())
	router.Post("/refresh-token", controllers.RefreshToken())
	router.Post("/forgot-password", controllers.ForgetPassword())
	router.Post("/reset-password", controllers.ResetPassword())
}
