package commentReaction

import (
	"github.com/gofiber/fiber/v3"
	"github.com/mcctrix/ctrix-social-go-backend/internal/domain/services"
)

type commentReactionHandler struct {
	commentReactionService *services.CommentReactionService
}

func NewCommentReactionHandler(commentReactionService *services.CommentReactionService) *commentReactionHandler {
	return &commentReactionHandler{
		commentReactionService: commentReactionService,
	}
}

func (h *commentReactionHandler) CreateCommentReaction(c fiber.Ctx) error {
	commentID := c.Params("commentid")
	userID := c.Locals("userID").(string)
	err := h.commentReactionService.CreateCommentReaction(commentID, userID)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "comment reaction created successfully"})
}

func (h *commentReactionHandler) DeleteCommentReaction(c fiber.Ctx) error {
	commentID := c.Params("commentid")
	userID := c.Locals("userID").(string)
	err := h.commentReactionService.DeleteCommentReaction(commentID, userID)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "comment reaction deleted successfully"})
}

func (h *commentReactionHandler) GetCommentReactionCount(c fiber.Ctx) error {
	commentID := c.Params("commentid")
	count, err := h.commentReactionService.GetCommentReactioCount(commentID)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(fiber.Map{"message": "comment reaction count fetched successfully", "data": count})
}
