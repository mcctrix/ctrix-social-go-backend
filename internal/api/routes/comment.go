package routes

import (
	"github.com/gofiber/fiber/v3"
	comments "github.com/mcctrix/ctrix-social-go-backend/internal/api/handlers/comment"
	commentReaction "github.com/mcctrix/ctrix-social-go-backend/internal/api/handlers/comment_reaction"
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

	// Comment Reaction
	commentReactionHandler := commentReaction.NewCommentReactionHandler(services.CommentReactionService)
	router.Get("/:commentid/reaction", commentReactionHandler.GetCommentReactionCount)
	router.Post("/:commentid/reaction", middleware.AuthMiddleware(), commentReactionHandler.CreateCommentReaction)
	router.Delete("/:commentid/reaction", middleware.AuthMiddleware(), commentReactionHandler.DeleteCommentReaction)
}
