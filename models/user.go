package models

import "time"

type User_Auth struct {
	Id         string
	Email      string
	Username   string
	Password   string
	Created_at time.Time
}

type User_Profile struct {
	Id              string    `json:"id,omitempty"`
	First_name      string    `json:"first_name,omitempty"`
	Last_name       string    `json:"last_name,omitempty"`
	Profile_picture string    `json:"profile_profile,omitempty"`
	Avatar          string    `json:"avatar,omitempty"`
	Last_seen       time.Time `json:"last_seen,omitempty"`
	Post_count      int       `json:"post_count,omitempty"`
	Followers       []string  `json:"followers,omitempty" gorm:"type:text[]"`
	Followings      []string  `json:"followings,omitempty" gorm:"type:text[]"`
}

type User_Additional_Info struct {
	Id              string    `json:"id,omitempty"`
	Hobbies         []string  `json:"hobbies,omitempty" gorm:"type:text[]"`
	Family_members  []string  `json:"family_members,omitempty" gorm:"type:text[]"`
	Relation_status string    `json:"relation_status,omitempty"`
	Dob             time.Time `json:"dob,omitempty"`
	Bio             string    `json:"bio,omitempty"`
	Gender          string    `json:"gender,omitempty"`
}

type User_Settings struct {
	Id          string   `json:"id,omitempty"`
	Hide_post   []string `json:"hide_post,omitempty" gorm:"type:text[]"`
	Hide_story  []string `json:"hide_story,omitempty" gorm:"type:text[]"`
	Block_user  []string `json:"block_user,omitempty" gorm:"type:text[]"`
	Show_online bool     `json:"show_online" gorm:"type:text[]"`
}

type User_Data struct {
	Id      string   `json:"id,omitempty"`
	Posts   []string `json:"posts,omitempty" gorm:"type:text[]"`
	Stories []string `json:"stories,omitempty" gorm:"type:text[]"`
	Notes   []string `json:"notes,omitempty" gorm:"type:text[]"`
}

type User_Posts struct {
	Id                string    `json:"id,omitempty" gorm:"primaryKey;type:uuid;default:uuid_generate_v4()"`
	Creator_id        string    `json:"creator_id"`
	Created_at        time.Time `json:"created_at"`
	Group_id          string    `json:"group_id"`
	Text_content      string    `json:"text_content"`
	Pictures_attached []string  `json:"pictures_attached" gorm:"type:text[]"`
	Comments          []string  `json:"comments" gorm:"type:text[]"`
}
type User_Post_Like_Table struct {
	User_id   string `json:"user_id"`
	Post_id   string `json:"post_id,omitempty"`
	Like_type string `json:"like_type"`
}

type User_post_Comments struct {
	Id                string    `json:"id" gorm:"primaryKey;type:uuid;default:uuid_generate_v4()"`
	Post_id           string    `json:"post_id"`
	Creator_id        string    `json:"creator_id"`
	Created_at        time.Time `json:"created_at"`
	Content           string    `json:"content"`
	Pictures_attached []string  `json:"pictures_attached,omitempty"`
	Nested_comments   []string  `json:"nested_comments,omitempty"`
	Liked_by          []string  `json:"liked_by,omitempty"`
}
