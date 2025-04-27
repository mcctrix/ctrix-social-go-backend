package routes

import (
	"github.com/gofiber/fiber/v3"
	"github.com/mcctrix/ctrix-social-go-backend/controllers"
	"github.com/mcctrix/ctrix-social-go-backend/middleware"
)

func PostManagementRouter(router fiber.Router) {
	router.Use(middleware.AuthMiddleware())

	// Post-related routes
	router.Get("/", controllers.GetUserPosts())
	router.Post("/", controllers.CreateUserPost())
	router.Get("/:postid", controllers.GetPostByID())
	router.Patch("/:postid", controllers.UpdateUserPost())
	router.Delete("/:postid", controllers.DeleteUserPost())
	router.Patch("/:postid/liketoggler", controllers.LikeToggler())

	// Comment-related routes
	router.Get("/:postid/comments", controllers.GetPostComments())
	router.Post("/:postid/comments", controllers.CreatePostComment())
	router.Get("/comments/:commentid", controllers.GetCommentByID())
	router.Patch("/comments/:commentid", controllers.UpdatePostComment())
	router.Delete("/comments/:commentid", controllers.DeletePostComment())
}
