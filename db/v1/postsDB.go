package v1

import (
	"encoding/json"
	"errors"
	"fmt"
	"time"

	"github.com/mcctrix/ctrix-social-go-backend/models"
)

// Posts Database Functions
func GetUserPostsByID(id string) ([]models.User_Post, error) {
	db, err := DBConnection()
	if err != nil {
		return nil, err
	}
	var userPosts []models.User_Post
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
	newPost := &models.User_Post{}
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

func GetPostByID(postID string) (*models.User_Post, error) {
	db, err := DBConnection()
	if err != nil {
		return nil, err
	}
	var post *models.User_Post = &models.User_Post{}
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
	if err = db.Table("user_post_comments").Order("updated_at desc").Where("post_id = ?", postID).Find(&postComments).Error; err != nil {
		return nil, err
	}
	return postComments, nil
}

func CreatePostCommentWithByteData(newCommentByte []byte, userID string, postID string) error {
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
	newComment.Created_at = time.Now()
	newComment.Updated_at = time.Now()
	newComment.Creator_id = userID
	newComment.Post_id = postID

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
	var existingPost *models.User_Post = &models.User_Post{}
	if err = db.Table("user_posts").Where("id = ? AND creator_id = ?", postID, userID).First(existingPost).Error; err != nil {
		return err
	}

	// Unmarshal the updated post data
	if err = json.Unmarshal(updatedPostByte, existingPost); err != nil {
		return err
	}

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
	result := db.Table("user_posts").Where("id = ? AND creator_id = ?", postID, userID).Delete(&models.User_Post{})
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

func GetAllPostReaction(postID string) ([]models.User_Post_Like_Table, error) {
	db, err := DBConnection()
	if err != nil {
		return nil, err
	}

	var allReacts []models.User_Post_Like_Table

	if err = db.Table("user_post_like").Where("post_id = ? ", postID).Select("user_id", "like_type").Find(&allReacts).Error; err != nil {
		fmt.Println(err)
		return nil, err
	}

	return allReacts, nil
}

func PostLikeToggler(postID string, userLikedID string, liked bool, likeType string) error {
	db, err := DBConnection()
	if err != nil {
		return err
	}

	if liked {
		post_like_data := models.User_Post_Like_Table{User_id: userLikedID, Post_id: postID, Like_type: "like"}
		if db.Table("user_post_like").Where("user_id = ?", userLikedID).Where("post_id = ?", postID).Updates(map[string]interface{}{
			"like_type": likeType,
		}).RowsAffected == 0 {
			if err = db.Table("user_post_like").Create(post_like_data).Error; err != nil {
				fmt.Println(err)
				return errors.New("unable to create post reaction")
			}
		}

	} else {

		if err = db.Table("user_post_like").Where("post_id = ?", postID).Where("user_id = ?", userLikedID).Delete(&models.User_Post_Like_Table{}).Error; err != nil {
			fmt.Println(err)
			return errors.New("unable to remove post reaction")
		}
	}

	return nil
}

func CommentLikeToggler(commentID string, userLikedID string, liked bool, likeType string) error {
	db, err := DBConnection()
	if err != nil {
		return err
	}

	if liked {
		comment_like_data := models.User_comment_like{User_id: userLikedID, Comment_id: commentID, Like_type: "like"}
		if db.Table("user_comment_like").Where("user_id = ?", userLikedID).Where("comment_id = ?", commentID).Updates(map[string]interface{}{
			"like_type": likeType,
		}).RowsAffected == 0 {
			if err = db.Table("user_comment_like").Create(comment_like_data).Error; err != nil {
				fmt.Println(err)
				return errors.New("unable to create comment reaction")
			}
		}

	} else {

		if err = db.Table("user_comment_like").Where("comment_id = ?", commentID).Where("user_id = ?", userLikedID).Delete(&models.User_comment_like{}).Error; err != nil {
			fmt.Println(err)
			return errors.New("unable to remove comment reaction")
		}
	}

	return nil
}
