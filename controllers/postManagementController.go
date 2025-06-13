package controllers

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/gofiber/fiber/v3"
	"github.com/google/uuid"
	db "github.com/mcctrix/ctrix-social-go-backend/db/v1"
	"github.com/mcctrix/ctrix-social-go-backend/models"
	"github.com/mcctrix/ctrix-social-go-backend/utils"
	"github.com/mcctrix/ctrix-social-go-backend/utils/cloudinary"
)

func GetUserPosts() fiber.Handler {
	return func(c fiber.Ctx) error {

		posts, err := db.GetUserPostsByID(c.Locals("userID").(string))
		if err != nil {
			fmt.Println("error while fetching user posts: ", err)
			return c.Status(fiber.StatusNotFound).SendString("unable to fetch user posts!")
		}

		return c.Status(fiber.StatusOK).JSON(posts)
	}
}

func GetPostReactions() fiber.Handler {
	return func(c fiber.Ctx) error {
		postID := c.Params("postid")
		reactions, err := db.GetAllPostReaction(postID)
		if err != nil {
			fmt.Println(err)
			return c.Status(fiber.StatusInternalServerError).SendString("Unable to fetch Reactions for the post!")
		}

		return c.Status(fiber.StatusOK).JSON(reactions)
	}
}

func CreateUserPost() fiber.Handler {
	return func(c fiber.Ctx) error {

		postData := &models.User_Post{}

		cloudinaryURLs, err := cloudinary.UploadMediaHandler(c)
		if err != nil {
			fmt.Println("Error uploading media: ", err)
			return fiber.ErrInternalServerError
		}
		if len(cloudinaryURLs) > 0 {
			postData.Media_attached = cloudinaryURLs
		}

		formData := c.FormValue("post_data")

		postData.Id = uuid.NewString()
		postData.Created_at = time.Now()
		postData.Updated_at = time.Now()
		postData.Creator_id = c.Locals("userID").(string)
		postData.Comments = nil

		bodyData, err := utils.ClearStruct(postData, []byte(formData))
		if err != nil {
			fmt.Println("Error clearing struct: ", err)
			return fiber.ErrInternalServerError
		}

		err = db.CreateUserPostWithByteData(bodyData, c.Locals("userID").(string))
		if err != nil {
			fmt.Println("Error creating post: ", err)
			return fiber.ErrInternalServerError
		}

		return c.Status(fiber.StatusCreated).SendString("Post created successfully!")
	}
}

func GetPostByID() fiber.Handler {
	return func(c fiber.Ctx) error {
		postID := c.Params("postid")
		post, err := db.GetPostByID(postID)
		if err != nil {
			fmt.Println("unable to fetch post: ", err)
			return c.Status(500).SendString("unable to fetch post!")
		}
		return c.JSON(post)
	}
}

func UpdateUserPost() fiber.Handler {
	return func(c fiber.Ctx) error {
		bodyData := &struct {
			Updated_at     time.Time          `json:"updated_at"`
			Text_content   string             `json:"text_content"`
			Media_attached models.StringArray `json:"pictures_attached" gorm:"type:text[]"`
		}{}

		cloudinaryURLs, err := cloudinary.UploadMediaHandler(c)
		if err != nil {
			fmt.Println("Error uploading media: ", err)
			return fiber.ErrInternalServerError
		}
		if len(cloudinaryURLs) > 0 {
			bodyData.Media_attached = cloudinaryURLs
		}

		formData := c.FormValue("post_data")

		bodyData.Updated_at = time.Now()

		rawData, err := utils.ClearStruct(bodyData, []byte(formData))
		if err != nil {
			fmt.Println("Error clearing struct: ", err)
			return fiber.ErrInternalServerError
		}

		postID := c.Params("postid")
		err = db.UpdateUserPostWithByteData(postID, rawData, c.Locals("userID").(string))
		if err != nil {
			fmt.Println("Error updating post: ", err)
			return fiber.ErrInternalServerError
		}

		return c.SendString("Post updated successfully!")
	}
}

func DeleteUserPost() fiber.Handler {
	return func(c fiber.Ctx) error {

		postID := c.Params("postid")
		err := db.DeleteUserPost(postID, c.Locals("userID").(string))
		if err != nil {
			fmt.Println("Error deleting post: ", err)
			return fiber.ErrInternalServerError
		}

		return c.SendString("Post deleted successfully!")
	}
}

func GetPostComments() fiber.Handler {
	return func(c fiber.Ctx) error {
		postID := c.Params("postid")
		comments, err := db.GetPostCommentsByPostID(postID)
		if err != nil {
			fmt.Println("error while fetching post comments: ", err)
			return c.Status(500).SendString("unable to fetch post comments!")
		}

		return c.JSON(comments)
	}
}

func CreatePostComment() fiber.Handler {
	return func(c fiber.Ctx) error {

		postID := c.Params("postid")
		commentData := c.BodyRaw()

		err := db.CreatePostCommentWithByteData(commentData, c.Locals("userID").(string), postID)
		if err != nil {
			fmt.Println("Error creating comment: ", err)
			return fiber.ErrInternalServerError
		}

		return c.SendString("Comment created successfully!")
	}
}

func GetCommentByID() fiber.Handler {
	return func(c fiber.Ctx) error {
		commentID := c.Params("commentid")
		comment, err := db.GetCommentByID(commentID)
		if err != nil {
			fmt.Println("unable to fetch comment: ", err)
			return c.Status(500).SendString("unable to fetch comment!")
		}
		return c.JSON(comment)
	}
}

func UpdatePostComment() fiber.Handler {
	return func(c fiber.Ctx) error {
		commentData := &struct {
			Updated_at time.Time `json:"updated_at"`
			Content    string    `json:"content"`
			Giff       string    `json:"giff"`
		}{}

		commentData.Updated_at = time.Now()

		rawData, err := utils.ClearStruct(commentData, c.BodyRaw())
		if err != nil {
			fmt.Println("Error clearing struct: ", err)
			return fiber.ErrInternalServerError
		}

		commentID := c.Params("commentid")
		err = db.UpdatePostCommentWithByteData(commentID, rawData, c.Locals("userID").(string))
		if err != nil {
			fmt.Println("Error updating comment: ", err)
			return fiber.ErrInternalServerError
		}

		return c.SendString("Comment updated successfully!")
	}
}

func DeletePostComment() fiber.Handler {
	return func(c fiber.Ctx) error {

		commentID := c.Params("commentid")
		err := db.DeletePostComment(commentID, c.Locals("userID").(string))
		if err != nil {
			fmt.Println("Error deleting comment: ", err)
			return fiber.ErrInternalServerError
		}

		return c.SendString("Comment deleted successfully!")
	}
}

func PostLikeToggler() fiber.Handler {
	return func(c fiber.Ctx) error {

		bodyData := &struct {
			Toggle    bool
			Like_type string
		}{}
		rawData := c.Body()

		err := json.Unmarshal(rawData, bodyData)
		if err != nil {
			fmt.Println(err)
			return fiber.ErrInternalServerError
		}

		if err = db.PostLikeToggler(c.Params("postid"), c.Locals("userID").(string), bodyData.Toggle, bodyData.Like_type); err != nil {
			fmt.Println(err)
			return err
		}

		return c.Status(fiber.StatusOK).SendString("Like Updated Successfully!")
	}
}

func CommentLikeToggler() fiber.Handler {
	return func(c fiber.Ctx) error {

		bodyData := &struct {
			Toggle    bool
			Like_type string
		}{}
		rawData := c.Body()

		err := json.Unmarshal(rawData, bodyData)
		if err != nil {
			fmt.Println(err)
			return fiber.ErrInternalServerError
		}

		if err = db.CommentLikeToggler(c.Params("commentid"), c.Locals("userID").(string), bodyData.Toggle, bodyData.Like_type); err != nil {
			fmt.Println(err)
			return err
		}

		return c.Status(fiber.StatusOK).SendString("Like Updated Successfully!")
	}
}
