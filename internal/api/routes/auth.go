package routes

import (
	"github.com/gofiber/fiber/v3"
	"github.com/mcctrix/ctrix-social-go-backend/internal/api/handlers/auth"
	"github.com/mcctrix/ctrix-social-go-backend/internal/domain/services"
)

func AuthRouter(router fiber.Router, services *services.Services) {
	authHandler := auth.NewAuthService(services.UserService)

	router.Post("/login", authHandler.Login)
	router.Post("/signup", authHandler.SignUp)
	router.Post("/logout", authHandler.Logout)
	router.Post("/refresh-token", authHandler.RefreshToken)
	router.Post("/forgot-password", authHandler.ForgetPassword)
	router.Post("/reset-password", authHandler.ResetPassword)
}
