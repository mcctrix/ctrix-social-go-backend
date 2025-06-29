package models

import "time"

type Bookmark struct {
	User_id    string    `json:"user_id,omitempty"`
	Created_at time.Time `json:"created_at,omitempty"`
	Post_id    string    `json:"post_id,omitempty"`
}
