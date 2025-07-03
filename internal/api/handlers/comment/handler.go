package comments

import (
	"encoding/json"

	"github.com/gofiber/fiber/v3"
	"github.com/google/uuid"
	"github.com/mcctrix/ctrix-social-go-backend/internal/domain/models"
	"github.com/mcctrix/ctrix-social-go-backend/internal/domain/services"
	"github.com/mcctrix/ctrix-social-go-backend/internal/pkg/utils"
)

type CommentHandler struct {
	commentService *services.CommentService
}

func NewCommentHandler(commentService *services.CommentService) *CommentHandler {
	return &CommentHandler{
		commentService: commentService,
	}
}

func (h *CommentHandler) CreateComment(c fiber.Ctx) error {
	postID := c.Params("postid")
	rawData := c.BodyRaw()
	comment := &models.User_post_Comments{}
	err := json.Unmarshal(rawData, comment)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	comment.Post_id = postID
	comment.Id = uuid.New().String()
	comment.Creator_id = c.Locals("userID").(string)

	err = h.commentService.CreateComment(comment)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "comment created successfully"})
}

func (h *CommentHandler) GetCommentByID(c fiber.Ctx) error {
	commentID := c.Params("commentid")

	comment, err := h.commentService.GetCommentByID(commentID)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(fiber.Map{"message": "comment fetched successfully", "data": comment})
}

func (h *CommentHandler) GetCommentsByPostID(c fiber.Ctx) error {
	postID := c.Params("postid")
	limit := utils.QueryLimit(c.Query("limit"))

	comments, err := h.commentService.GetCommentsByPostID(postID, limit)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(fiber.Map{"message": "comments fetched successfully", "data": comments})
}

func (h *CommentHandler) GetCommentByUserID(c fiber.Ctx) error {
	userID := c.Params("userid")
	limit := utils.QueryLimit(c.Query("limit"))

	comments, err := h.commentService.GetCommentByUserID(userID, limit)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(fiber.Map{"message": "comments fetched successfully", "data": comments})
}

func (h *CommentHandler) UpdateCommentByID(c fiber.Ctx) error {
	commentID := c.Params("commentid")
	rawData := c.BodyRaw()
	comment := &models.User_post_Comments{}
	err := json.Unmarshal(rawData, comment)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	err = h.commentService.UpdateCommentByID(commentID, comment, c.Locals("userID").(string))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "comment updated successfully"})
}

func (h *CommentHandler) DeleteCommentByID(c fiber.Ctx) error {
	commentID := c.Params("commentid")
	err := h.commentService.DeleteCommentByID(commentID, c.Locals("userID").(string))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "comment deleted successfully"})
}
