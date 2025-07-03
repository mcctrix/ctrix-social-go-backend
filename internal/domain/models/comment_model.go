package models

import "time"

type User_post_Comments struct {
	Id         string    `json:"id" gorm:"column:id;type:varchar(50);primary_key;default:uuid_generate_v4()"`
	Post_id    string    `json:"post_id" gorm:"column:post_id;type:varchar(50);not null"`
	Creator_id string    `json:"creator_id" gorm:"column:creator_id;type:varchar(50);not null"`
	Created_at time.Time `json:"created_at" gorm:"column:created_at;default:now()"`
	Updated_at time.Time `json:"updated_at" gorm:"column:updated_at;default:now()"`
	Content    string    `json:"content" gorm:"column:content;type:text"`
	Giff       string    `json:"giff" gorm:"column:giff;type:text"`
}

func (User_post_Comments) TableName() string {
	return "user_post_comments"
}

type User_post_comment_like struct {
	User_id    string `json:"user_id" gorm:"column:user_id;type:varchar(50)"`
	Comment_id string `json:"comment_id" gorm:"column:comment_id;type:varchar(50);not null"`
}

func (User_post_comment_like) TableName() string {
	return "user_post_comment_like"
}
