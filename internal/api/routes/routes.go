package routes

import (
	"github.com/gofiber/fiber/v3"
	"github.com/mcctrix/ctrix-social-go-backend/internal/domain/services"
)

func SetupRoutes(endpoint string, app *fiber.App, services *services.Services) {

	routeGroup := app.Group(endpoint)

	AuthRouter(routeGroup.Group("/auth"), services)
	UserRouter(routeGroup.Group("/profile"), services)
	FollowRouter(routeGroup.Group("/follow"), services)
	UserSettingsRouter(routeGroup.Group("/user_settings"), services)
	AdditionalUserInfoRouter(routeGroup.Group("/additional_info_user"), services)
	BookmarkRouter(routeGroup.Group("/bookmark"), services)
	CommentRouter(routeGroup.Group("/comment"), services)
	PostRouter(routeGroup.Group("/post"))
	FeedRouter(routeGroup.Group("/feed"))
}
