package routes

import "github.com/gofiber/fiber/v3"

func SetupRoutes(app *fiber.App) {
	AuthRouter(app.Group("/api/auth"))
	UserRouter(app.Group("/api/profile"))
	PostRouter(app.Group("/api/post"))
	CommentRouter(app.Group("/api/comments"))
	FeedRouter(app.Group("/api/feed"))
	BookmarkRouter(app.Group("/api/bookmark"))
}
