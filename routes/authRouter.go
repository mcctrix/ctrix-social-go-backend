package routes

import (
	"fmt"

	"github.com/gofiber/fiber/v3"
	"github.com/mcctrix/ctrix-social-go-backend/controllers"
	"github.com/mcctrix/ctrix-social-go-backend/middleware"
)

func AuthRouter(router fiber.Router) {
	router.Use(middleware.AuthMiddleware())
	fmt.Println("Reach in AuthRouter!")
	router.Get("/", controllers.Login())

}
