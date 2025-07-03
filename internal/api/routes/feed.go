package routes

import (
	"github.com/gofiber/fiber/v3"
	"github.com/mcctrix/ctrix-social-go-backend/internal/api/handlers/feed"
	"github.com/mcctrix/ctrix-social-go-backend/internal/api/middleware"
	"github.com/mcctrix/ctrix-social-go-backend/internal/domain/services"
)

func FeedRouter(router fiber.Router, services *services.Services) {
	router.Use(middleware.AuthMiddleware())

	handler := feed.NewFeedHandler(services.FeedService)

	router.Get("/", handler.GetFeed)
	router.Get("/follow_recommendation", handler.GetFollowRecommendation)
}
