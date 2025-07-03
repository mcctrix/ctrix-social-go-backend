package models

import "time"

type Follow struct {
	Follower_id  string    `json:"follower_id,omitempty" gorm:"column:follower_id;type:text;not null"`
	Following_id string    `json:"following_id,omitempty" gorm:"column:following_id;type:text;not null"`
	Created_at   time.Time `json:"created_at,omitempty" gorm:"column:created_at;default:CURRENT_TIMESTAMP"`
}

func (Follow) TableName() string {
	return "follow"
}

func NewFollow(follower_id, following_id string) *Follow {
	return &Follow{
		Follower_id:  follower_id,
		Following_id: following_id,
		Created_at:   time.Now(),
	}
}
