package routes

import (
	"github.com/gofiber/fiber/v3"
	"github.com/mcctrix/ctrix-social-go-backend/controllers"
	"github.com/mcctrix/ctrix-social-go-backend/middleware"
)

func BookmarkRouter(router fiber.Router) {
	router.Use(middleware.AuthMiddleware())

	router.Get("/", controllers.GetBookmark())
	router.Post("/:postID", controllers.CreateBookmark())
	router.Delete("/:postID", controllers.DeleteBookmark())
}
