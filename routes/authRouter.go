package routes

import (
	"github.com/gofiber/fiber/v3"
	"github.com/mcctrix/ctrix-social-go-backend/controllers"
)

func AuthRouter(router fiber.Router) {
	router.Post("/login", controllers.Login())
	router.Post("/signup", controllers.SignUp())
	router.Get("/logout", controllers.Logout())

}
