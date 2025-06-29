package models

import "time"

type User_post_Comments struct {
	Id         string    `json:"id" gorm:"primaryKey;type:uuid;default:uuid_generate_v4()"`
	Post_id    string    `json:"post_id"`
	Creator_id string    `json:"creator_id"`
	Created_at time.Time `json:"created_at"`
	Updated_at time.Time `json:"updated_at"`
	Content    string    `json:"content"`
	Giff       string    `json:"giff"`
}

type User_post_comment_like struct {
	User_id    string `json:"user_id"`
	Comment_id string `json:"comment_id"`
	Like_type  string `json:"like_type"`
}
