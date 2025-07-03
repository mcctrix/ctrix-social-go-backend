package routes

import (
	"github.com/gofiber/fiber/v3"
	"github.com/mcctrix/ctrix-social-go-backend/internal/api/handlers/posts"
	"github.com/mcctrix/ctrix-social-go-backend/internal/api/middleware"
	"github.com/mcctrix/ctrix-social-go-backend/internal/domain/services"
)

func PostRouter(router fiber.Router, services *services.Services) {
	router.Use(middleware.AuthMiddleware())

	handler := posts.NewPostHandler(services.PostService)

	// Post-related routes
	router.Get("/", handler.GetUserPosts)
	router.Post("/", handler.CreatePost)
	router.Get("/:postid", handler.GetPostByID)
	router.Patch("/:postid", handler.UpdatePost)
	router.Delete("/:postid", handler.DeletePost)

	// Reactions
	router.Get("/:postid/reaction", handler.GetPostReactions)
	router.Post("/:postid/reaction", handler.LikePost)
	router.Delete("/:postid/reaction", handler.DislikePost)
}
