package controllers

import (
	"fmt"

	"github.com/gofiber/fiber/v3"
	"github.com/mcctrix/ctrix-social-go-backend/db"
	"github.com/mcctrix/ctrix-social-go-backend/utils"
)

func GetUserPosts() fiber.Handler {
	return func(c fiber.Ctx) error {
		userID, err := utils.GetUserIDWithToken(c.Cookies("auth_token"))
		if err != nil {
			fmt.Println("unable to fetch user with this Token: ", err)
			return c.Status(401).SendString("unable to fetch user with this Token!")
		}

		posts, err := db.GetUserPostsByID(userID)
		if err != nil {
			fmt.Println("error while fetching user posts: ", err)
			return c.Status(500).SendString("unable to fetch user posts!")
		}

		return c.JSON(posts)
	}
}

func CreateUserPost() fiber.Handler {
	return func(c fiber.Ctx) error {
		userID, err := utils.GetUserIDWithToken(c.Cookies("auth_token"))
		if err != nil {
			fmt.Println("unable to fetch user with this Token: ", err)
			return c.Status(401).SendString("unable to fetch user with this Token!")
		}

		err = db.CreateUserPostWithByteData(c.BodyRaw(), userID)
		if err != nil {
			fmt.Println("Error creating post: ", err)
			return fiber.ErrInternalServerError
		}

		return c.SendString("Post created successfully!")
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
		userID, err := utils.GetUserIDWithToken(c.Cookies("auth_token"))
		if err != nil {
			fmt.Println("unable to fetch user with this Token: ", err)
			return c.Status(401).SendString("unable to fetch user with this Token!")
		}

		postID := c.Params("postid")
		err = db.UpdateUserPostWithByteData(postID, c.BodyRaw(), userID)
		if err != nil {
			fmt.Println("Error updating post: ", err)
			return fiber.ErrInternalServerError
		}

		return c.SendString("Post updated successfully!")
	}
}

func DeleteUserPost() fiber.Handler {
	return func(c fiber.Ctx) error {
		userID, err := utils.GetUserIDWithToken(c.Cookies("auth_token"))
		if err != nil {
			fmt.Println("unable to fetch user with this Token: ", err)
			return c.Status(401).SendString("unable to fetch user with this Token!")
		}

		postID := c.Params("postid")
		err = db.DeleteUserPost(postID, userID)
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
		userID, err := utils.GetUserIDWithToken(c.Cookies("auth_token"))
		if err != nil {
			fmt.Println("unable to fetch user with this Token: ", err)
			return c.Status(401).SendString("unable to fetch user with this Token!")
		}

		postID := c.Params("postid")
		// Modify the byte data to include the post ID
		commentData := c.BodyRaw()
		var commentMap map[string]interface{}
		if err := json.Unmarshal(commentData, &commentMap); err != nil {
			return fiber.ErrBadRequest
		}
		commentMap["post_id"] = postID
		modifiedCommentData, err := json.Marshal(commentMap)
		if err != nil {
			return fiber.ErrInternalServerError
		}

		err = db.CreatePostCommentWithByteData(modifiedCommentData, userID)
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
		userID, err := utils.GetUserIDWithToken(c.Cookies("auth_token"))
		if err != nil {
			fmt.Println("unable to fetch user with this Token: ", err)
			return c.Status(401).SendString("unable to fetch user with this Token!")
		}

		commentID := c.Params("commentid")
		err = db.UpdatePostCommentWithByteData(commentID, c.BodyRaw(), userID)
		if err != nil {
			fmt.Println("Error updating comment: ", err)
			return fiber.ErrInternalServerError
		}

		return c.SendString("Comment updated successfully!")
	}
}

func DeletePostComment() fiber.Handler {
	return func(c fiber.Ctx) error {
		userID, err := utils.GetUserIDWithToken(c.Cookies("auth_token"))
		if err != nil {
			fmt.Println("unable to fetch user with this Token: ", err)
			return c.Status(401).SendString("unable to fetch user with this Token!")
		}

		commentID := c.Params("commentid")
		err = db.DeletePostComment(commentID, userID)
		if err != nil {
			fmt.Println("Error deleting comment: ", err)
			return fiber.ErrInternalServerError
		}

		return c.SendString("Comment deleted successfully!")
	}
}
