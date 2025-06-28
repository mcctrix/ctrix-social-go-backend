package routes

import (
	"github.com/gofiber/fiber/v3"
	"github.com/mcctrix/ctrix-social-go-backend/internal/api/handlers/posts"
	"github.com/mcctrix/ctrix-social-go-backend/internal/api/middleware"
)

func PostManagementRouter(router fiber.Router) {
	router.Use(middleware.AuthMiddleware())

	// Post-related routes
	router.Get("/", posts.GetUserPosts())
	router.Post("/", posts.CreateUserPost())
	router.Get("/:postid", posts.GetPostByID())
	router.Get("/:postid/reacts", posts.GetPostReactions())
	router.Patch("/:postid", posts.UpdateUserPost())
	router.Delete("/:postid", posts.DeleteUserPost())

	// Comment-related routes
	router.Get("/comments/:postid", posts.GetPostComments())
	router.Post("/comments/:postid", posts.CreatePostComment())

	// Reactions
	router.Patch("/:postid/liketoggler", posts.PostLikeToggler())
	router.Patch("/comments/:commentid/liketoggler", posts.CommentLikeToggler())
}