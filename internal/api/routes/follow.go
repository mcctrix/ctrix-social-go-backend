package routes

import (
	"github.com/gofiber/fiber/v3"
	"github.com/mcctrix/ctrix-social-go-backend/internal/api/handlers/follow"
	"github.com/mcctrix/ctrix-social-go-backend/internal/api/middleware"
	"github.com/mcctrix/ctrix-social-go-backend/internal/domain/services"
)

func FollowRouter(router fiber.Router, services *services.Services) {
	router.Use(middleware.AuthMiddleware())

	followHandler := follow.NewFollowHandler(services.FollowService)

	router.Post("/:followingid", followHandler.FollowUser)
	router.Delete("/:followingid", followHandler.UnFollowUser)
	router.Get("/check/:followingid", followHandler.CheckFollowing)
	router.Get("/count/:userid", followHandler.CountFollowAndFollowing)
}
