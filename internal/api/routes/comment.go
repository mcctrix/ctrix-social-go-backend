package routes

import (
	"github.com/gofiber/fiber/v3"
	"github.com/mcctrix/ctrix-social-go-backend/internal/api/handlers/comments"
	"github.com/mcctrix/ctrix-social-go-backend/internal/api/middleware"
	"github.com/mcctrix/ctrix-social-go-backend/internal/domain/services"
)

func CommentRouter(router fiber.Router, services *services.Services) {

	commentHandler := comments.NewCommentHandler(services.CommentService)

	router.Post("/post/:postid", middleware.AuthMiddleware(), commentHandler.CreateComment)
	router.Get("/:commentid", commentHandler.GetCommentByID)
	router.Get("/user/:userid", commentHandler.GetCommentByUserID)
	router.Get("/post/:postid", commentHandler.GetCommentsByPostID)
	router.Patch("/:commentid", middleware.AuthMiddleware(), commentHandler.UpdateCommentByID)
	router.Delete("/:commentid", middleware.AuthMiddleware(), commentHandler.DeleteCommentByID)
}
