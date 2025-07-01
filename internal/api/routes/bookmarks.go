package routes

import (
	"github.com/gofiber/fiber/v3"
	"github.com/mcctrix/ctrix-social-go-backend/internal/api/handlers/bookmarks"
	"github.com/mcctrix/ctrix-social-go-backend/internal/api/middleware"
	"github.com/mcctrix/ctrix-social-go-backend/internal/domain/services"
)

func BookmarkRouter(router fiber.Router, services *services.Services) {
	router.Use(middleware.AuthMiddleware())

	bookmarkHandler := bookmarks.NewBookmarkHandler(services.BookmarkService)

	router.Get("/", bookmarkHandler.GetBookmark)
	router.Post("/:postID", bookmarkHandler.CreateBookmark)
	router.Delete("/:postID", bookmarkHandler.DeleteBookmark)

}
