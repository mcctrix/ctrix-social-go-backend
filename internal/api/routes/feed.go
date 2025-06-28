package routes

import (
	"github.com/gofiber/fiber/v3"
	"github.com/mcctrix/ctrix-social-go-backend/internal/api/handlers/feed"
	"github.com/mcctrix/ctrix-social-go-backend/internal/api/middleware"
)

func FeedManagementRouter(router fiber.Router) {
	router.Use(middleware.AuthMiddleware())

	router.Get("/", feed.GetFeed())
	router.Get("/follow_recommendation", feed.GetFollowRecommendation())
}