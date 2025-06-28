package routes

import (
	"github.com/gofiber/fiber/v3"
	"github.com/mcctrix/ctrix-social-go-backend/internal/api/handlers/comments"
	"github.com/mcctrix/ctrix-social-go-backend/internal/api/middleware"
)

func CommentRouter(router fiber.Router) {
	router.Use(middleware.AuthMiddleware())

	router.Get("/:commentid", comments.GetCommentByID())
	router.Patch("/:commentid", comments.UpdatePostComment())
	router.Delete("/:commentid", comments.DeletePostComment())
}