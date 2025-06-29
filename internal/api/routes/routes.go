package routes

import (
	"github.com/gofiber/fiber/v3"
	"github.com/mcctrix/ctrix-social-go-backend/internal/domain/services"
)

func SetupRoutes(endpoint string, app *fiber.App, services *services.Services) {

	routeGroup := app.Group(endpoint)

	AuthRouter(routeGroup.Group("/auth"), services)
	UserRouter(routeGroup.Group("/profile"))
	PostRouter(routeGroup.Group("/post"))
	CommentRouter(routeGroup.Group("/comments"))
	FeedRouter(routeGroup.Group("/feed"))
	BookmarkRouter(routeGroup.Group("/bookmark"))
}
