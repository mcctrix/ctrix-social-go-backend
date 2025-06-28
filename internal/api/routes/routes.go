package routes

import "github.com/gofiber/fiber/v3"

func SetupRoutes(endpoint string, app *fiber.App) {

	routeGroup := app.Group(endpoint)

	AuthRouter(routeGroup.Group("/auth"))
	UserRouter(routeGroup.Group("/profile"))
	PostRouter(routeGroup.Group("/post"))
	CommentRouter(routeGroup.Group("/comments"))
	FeedRouter(routeGroup.Group("/feed"))
	BookmarkRouter(routeGroup.Group("/bookmark"))
}
