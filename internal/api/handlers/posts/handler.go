package posts

import (
	"encoding/json"
	"time"

	"github.com/gofiber/fiber/v3"
	"github.com/google/uuid"
	"github.com/mcctrix/ctrix-social-go-backend/internal/domain/models"
	"github.com/mcctrix/ctrix-social-go-backend/internal/domain/services"
	"github.com/mcctrix/ctrix-social-go-backend/internal/pkg/utils"
)

type PostHandler struct {
	postService *services.PostService
}

func NewPostHandler(postService *services.PostService) *PostHandler {
	return &PostHandler{postService: postService}
}

func (h *PostHandler) CreatePost(c fiber.Ctx) error {
	post := &models.User_Post{}
	err := json.Unmarshal(c.BodyRaw(), post)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}
	post.Id = uuid.New().String()
	post.Created_at = time.Now()
	post.Updated_at = time.Now()
	post.Creator_id = c.Locals("userID").(string)

	err = h.postService.CreatePost(post)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.Status(fiber.StatusCreated).JSON(fiber.Map{"message": "Post created successfully!"})
}

func (h *PostHandler) GetUserPosts(c fiber.Ctx) error {
	userID := c.Locals("userID").(string)
	limit := utils.QueryLimit(c.Query("limit", "5"))
	posts, err := h.postService.GetUserPostsByID(userID, limit)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "unable to fetch user posts!"})
	}
	return c.Status(fiber.StatusOK).JSON(posts)
}

func (h *PostHandler) GetPostByID(c fiber.Ctx) error {
	userID := c.Locals("userID").(string)
	postID := c.Params("postid")
	post, err := h.postService.GetPostByID(postID, userID)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "unable to fetch post!"})
	}
	return c.JSON(post)
}

func (h *PostHandler) UpdatePost(c fiber.Ctx) error {
	postID := c.Params("postid")
	userID := c.Locals("userID").(string)
	updatedPost := &models.User_Post{}
	err := json.Unmarshal(c.BodyRaw(), updatedPost)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}
	updatedPost.Updated_at = time.Now()
	err = h.postService.UpdatePost(postID, updatedPost, userID)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(fiber.Map{"message": "Post updated successfully!"})
}

func (h *PostHandler) DeletePost(c fiber.Ctx) error {
	postID := c.Params("postid")
	userID := c.Locals("userID").(string)
	err := h.postService.DeletePost(postID, userID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(fiber.Map{"message": "Post deleted successfully!"})
}

func (h *PostHandler) GetPostReactions(c fiber.Ctx) error {
	postID := c.Params("postid")
	reactions, err := h.postService.GetPostReactions(postID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Unable to fetch reactions for the post!"})
	}
	return c.Status(fiber.StatusOK).JSON(reactions)
}

func (h *PostHandler) LikePost(c fiber.Ctx) error {
	postID := c.Params("postid")
	userID := c.Locals("userID").(string)
	err := h.postService.LikePost(postID, userID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "Post liked successfully!"})
}

func (h *PostHandler) DislikePost(c fiber.Ctx) error {
	postID := c.Params("postid")
	userID := c.Locals("userID").(string)
	err := h.postService.DislikePost(postID, userID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "Post disliked successfully!"})
}
