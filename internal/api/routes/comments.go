package routes

import (
	"github.com/gofiber/fiber/v3"
	"github.com/mcctrix/ctrix-social-go-backend/controllers"
	"github.com/mcctrix/ctrix-social-go-backend/middleware"
)

func CommentRouter(router fiber.Router) {
	router.Use(middleware.AuthMiddleware())

	router.Get("/:commentid", controllers.GetCommentByID())
	router.Patch("/:commentid", controllers.UpdatePostComment())
	router.Delete("/:commentid", controllers.DeletePostComment())
}
