package routes

import (
	"github.com/gofiber/fiber/v3"
	"github.com/mcctrix/ctrix-social-go-backend/internal/api/handlers/auth"
)

func AuthRouter(router fiber.Router) {
	router.Post("/login", auth.Login())
	router.Post("/signup", auth.SignUp())
	router.Post("/logout", auth.Logout())
	router.Post("/refresh-token", auth.RefreshToken())
	router.Post("/forgot-password", auth.ForgetPassword())
	router.Post("/reset-password", auth.ResetPassword())
}
