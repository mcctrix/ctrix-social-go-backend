package routes

import (
	"github.com/gofiber/fiber/v3"
	"github.com/mcctrix/ctrix-social-go-backend/controllers"
	"github.com/mcctrix/ctrix-social-go-backend/middleware"
)

func FeedManagementRouter(router fiber.Router) {
	router.Use(middleware.AuthMiddleware())

	router.Get("/", controllers.GetFeed())
	router.Get("/follow_recommendation", controllers.GetFollowRecommendation())
}
