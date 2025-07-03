package models

import "time"

type Bookmark struct {
	User_id    string    `json:"user_id,omitempty" gorm:"column:user_id;type:varchar(50);not null"`
	Created_at time.Time `json:"created_at,omitempty" gorm:"column:created_at;default:now()"`
	Post_id    string    `json:"post_id,omitempty" gorm:"column:post_id;type:varchar(50);not null"`
}

func (Bookmark) TableName() string {
	return "bookmark"
}

func NewBookmark(user_id, post_id string) *Bookmark {
	return &Bookmark{
		User_id:    user_id,
		Created_at: time.Now(),
		Post_id:    post_id,
	}
}
