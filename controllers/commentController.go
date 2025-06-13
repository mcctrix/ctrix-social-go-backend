package controllers

import (
	"fmt"
	"time"

	"github.com/gofiber/fiber/v3"
	db "github.com/mcctrix/ctrix-social-go-backend/db/v1"
	"github.com/mcctrix/ctrix-social-go-backend/utils"
)

func GetCommentByID() fiber.Handler {
	return func(c fiber.Ctx) error {
		commentID := c.Params("commentid")
		fmt.Println(commentID)
		comment, err := db.GetCommentByID(commentID)
		if err != nil {
			fmt.Println("unable to fetch comment: ", err)
			return c.Status(fiber.StatusNotFound).SendString("unable to find comment!")
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
			return c.Status(fiber.StatusBadRequest).SendString("Unable to update comment!")
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
			return err
		}

		return c.SendString("Comment deleted successfully!")
	}
}
