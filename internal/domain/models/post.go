package models

import "time"

type User_Post struct {
	Id             string      `json:"id,omitempty" gorm:"column:id;type:varchar(50);primary_key;default:uuid_generate_v4()"`
	Creator_id     string      `json:"creator_id" gorm:"column:creator_id;type:varchar(50);not null"`
	Created_at     time.Time   `json:"created_at" gorm:"column:created_at;default:now()"`
	Group_id       string      `json:"group_id" gorm:"column:group_id;type:varchar(50)"`
	Updated_at     time.Time   `json:"updated_at" gorm:"column:updated_at;default:now()"`
	Text_content   string      `json:"text_content" gorm:"column:text_content;type:text"`
	Media_attached StringArray `json:"media_attached" gorm:"column:media_attached;type:text[]"`
}

func (User_Post) TableName() string {
	return "user_posts"
}

type User_Post_Like_Table struct {
	User_id   string `json:"user_id" gorm:"column:user_id;type:varchar(50)"`
	Post_id   string `json:"post_id,omitempty" gorm:"column:post_id;type:varchar(50);not null"`
	Like_type string `json:"like_type" gorm:"column:like_type;type:varchar(20)"`
}

func (User_Post_Like_Table) TableName() string {
	return "user_post_like"
}
