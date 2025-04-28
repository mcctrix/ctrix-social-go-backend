package db

import (
	"encoding/json"
	"errors"
	"fmt"
	"time"

	"github.com/gofiber/fiber/v3"
	"github.com/mcctrix/ctrix-social-go-backend/models"
	"gorm.io/gorm"
)

// Posts Database Functions
func GetUserPostsByID(id string) ([]models.User_Posts, error) {
	db, err := DBConnection()
	if err != nil {
		return nil, err
	}
	var userPosts []models.User_Posts
	if err = db.Table("user_posts").Order("created_at desc").Where("creator_id = ?", id).Find(&userPosts).Error; err != nil {
		return nil, err
	}
	return userPosts, nil
}

func CreateUserPostWithByteData(newPostByte []byte, userID string) error {
	db, err := DBConnection()
	if err != nil {
		return err
	}

	// Create a new post
	newPost := &models.User_Posts{}
	if err = json.Unmarshal(newPostByte, newPost); err != nil {
		return err
	}

	// Set the creator ID to the authenticated user
	newPost.Creator_id = userID
	newPost.Created_at = time.Now()

	// Save the post1
	if err = db.Table("user_posts").Create(newPost).Error; err != nil {
		return err
	}

	return nil
}

func GetPostByID(postID string) (*models.User_Posts, error) {
	db, err := DBConnection()
	if err != nil {
		return nil, err
	}
	var post *models.User_Posts = &models.User_Posts{}
	if err = db.Table("user_posts").Where("id = ?", postID).First(post).Error; err != nil {
		return nil, err
	}
	return post, nil
}

// Post Comments Database Functions
func GetPostCommentsByPostID(postID string) ([]*models.User_post_Comments, error) {
	db, err := DBConnection()
	if err != nil {
		return nil, err
	}
	var postComments []*models.User_post_Comments
	if err = db.Table("user_post_comments").Order("created_at desc").Where("post_id = ?", postID).Find(&postComments).Error; err != nil {
		return nil, err
	}
	return postComments, nil
}

func CreatePostCommentWithByteData(newCommentByte []byte, userID string) error {
	db, err := DBConnection()
	if err != nil {
		return err
	}

	// Create a new comment
	newComment := &models.User_post_Comments{}
	if err = json.Unmarshal(newCommentByte, newComment); err != nil {
		return err
	}

	// Set the creator ID to the authenticated user
	newComment.Creator_id = userID

	// Save the comment
	if err = db.Table("user_post_comments").Create(newComment).Error; err != nil {
		return err
	}

	return nil
}

func GetCommentByID(commentID string) (*models.User_post_Comments, error) {
	db, err := DBConnection()
	if err != nil {
		return nil, err
	}
	var comment *models.User_post_Comments = &models.User_post_Comments{}
	if err = db.Table("user_post_comments").Where("id = ?", commentID).First(comment).Error; err != nil {
		return nil, err
	}
	return comment, nil
}

// Update functions for Posts and Comments
func UpdateUserPostWithByteData(postID string, updatedPostByte []byte, userID string) error {
	db, err := DBConnection()
	if err != nil {
		return err
	}

	// Find the existing post
	var existingPost *models.User_Posts = &models.User_Posts{}
	if err = db.Table("user_posts").Where("id = ? AND creator_id = ?", postID, userID).First(existingPost).Error; err != nil {
		return err
	}

	fmt.Println("existing data:", existingPost)
	fmt.Println("New Data", string(updatedPostByte))
	// Unmarshal the updated post data
	if err = json.Unmarshal(updatedPostByte, existingPost); err != nil {
		return err
	}
	fmt.Println("UPDATED existing data:", existingPost)

	// Save the updated post
	if err = db.Table("user_posts").Save(existingPost).Error; err != nil {
		return err
	}

	return nil
}

func UpdatePostCommentWithByteData(commentID string, updatedCommentByte []byte, userID string) error {
	db, err := DBConnection()
	if err != nil {
		return err
	}

	// Find the existing comment
	var existingComment *models.User_post_Comments = &models.User_post_Comments{}
	if err = db.Table("user_post_comments").Where("id = ? AND creator_id = ?", commentID, userID).First(existingComment).Error; err != nil {
		return err
	}

	// Unmarshal the updated comment data
	if err = json.Unmarshal(updatedCommentByte, existingComment); err != nil {
		return err
	}

	// Save the updated comment
	if err = db.Table("user_post_comments").Save(existingComment).Error; err != nil {
		return err
	}

	return nil
}

// Delete functions for Posts and Comments
func DeleteUserPost(postID string, userID string) error {
	db, err := DBConnection()
	if err != nil {
		return err
	}

	// Delete the post, ensuring it belongs to the user
	result := db.Table("user_posts").Where("id = ? AND creator_id = ?", postID, userID).Delete(&models.User_Posts{})
	if result.Error != nil {
		return result.Error
	}

	// Check if a post was actually deleted
	if result.RowsAffected == 0 {
		return errors.New("no post found or unauthorized to delete")
	}

	return nil
}

func DeletePostComment(commentID string, userID string) error {
	db, err := DBConnection()
	if err != nil {
		return err
	}

	// Delete the comment, ensuring it belongs to the user
	result := db.Table("user_post_comments").Where("id = ? AND creator_id = ?", commentID, userID).Delete(&models.User_post_Comments{})
	if result.Error != nil {
		return result.Error
	}

	// Check if a comment was actually deleted
	if result.RowsAffected == 0 {
		return errors.New("no comment found or unauthorized to delete")
	}

	return nil
}

func PostLikeToggler(postID string, userToAddInLikedList string, like bool) error {
	db, err := DBConnection()
	if err != nil {
		return err
	}

	if like {
		// arrayLiteral := fmt.Sprintf("'{%s}'", userToAddInLikedList)
		if err = db.Table("user_posts").Where("id = ?", postID).Update("liked_by", gorm.Expr("array_cat(liked_by, ARRAY[?])", userToAddInLikedList)).Error; err != nil {
			fmt.Println("Error", err)
			return fiber.ErrInternalServerError
		}
	} else {
		if err = db.Table("user_posts").Where("id = ?", postID).Update("liked_by", gorm.Expr("array_remove(liked_by, ?::text)", []string{userToAddInLikedList})).Error; err != nil {
			fmt.Println("Error", err)
			return fiber.ErrInternalServerError
		}
	}

	return nil
}
