package models

import "time"

type User_Post struct {
	Id             string      `json:"id,omitempty" gorm:"primaryKey;type:uuid;default:uuid_generate_v4()"`
	Creator_id     string      `json:"creator_id"`
	Created_at     time.Time   `json:"created_at"`
	Group_id       string      `json:"group_id"`
	Text_content   string      `json:"text_content"`
	Media_attached StringArray `json:"pictures_attached" gorm:"type:text[]"`
	Comments       StringArray `json:"comments" gorm:"type:text[]"`
}

type User_Post_Like_Table struct {
	User_id   string `json:"user_id"`
	Post_id   string `json:"post_id,omitempty"`
	Like_type string `json:"like_type"`
}

type User_post_Comments struct {
	Id                string      `json:"id" gorm:"primaryKey;type:uuid;default:uuid_generate_v4()"`
	Post_id           string      `json:"post_id"`
	Creator_id        string      `json:"creator_id"`
	Created_at        time.Time   `json:"created_at"`
	Content           string      `json:"content"`
	Pictures_attached StringArray `json:"pictures_attached,omitempty"`
	Nested_comments   StringArray `json:"nested_comments,omitempty"`
}

type User_comment_like struct {
	User_id    string `json:"user_id"`
	Comment_id string `json:"comment_id"`
	Like_type  string `json:"like_type"`
}
