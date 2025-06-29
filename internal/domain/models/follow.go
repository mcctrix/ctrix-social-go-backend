package models

import "time"

type Follow struct {
	Follower_id  string    `json:"follower_id,omitempty"`
	Following_id string    `json:"following_id,omitempty"`
	Created_at   time.Time `json:"created_at,omitempty"`
}

func NewFollow(follower_id, following_id string) *Follow {
	return &Follow{
		Follower_id:  follower_id,
		Following_id: following_id,
		Created_at:   time.Now(),
	}
}
