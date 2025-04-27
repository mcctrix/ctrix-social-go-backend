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
	Id              string    `json:"id"`
	First_name      string    `json:"first_name"`
	Last_name       string    `json:"last_name"`
	Profile_picture string    `json:"profile_profile"`
	Avatar          string    `json:"avatar"`
	Last_seen       time.Time `json:"last_seen"`
	Post_count      int       `json:"post_count"`
	Followers       []string  `json:"followers" gorm:"type:text[]"`
	Followings      []string  `json:"followings" gorm:"type:text[]"`
}

type User_Additional_Info struct {
	Id              string    `json:"id"`
	Hobbies         []string  `json:"hobbies" gorm:"type:text[]"`
	Family_members  []string  `json:"family_members" gorm:"type:text[]"`
	Relation_status string    `json:"relation_status"`
	Dob             time.Time `json:"dob"`
	Bio             string    `json:"bio"`
	Gender          string    `json:"gender"`
}

type User_Settings struct {
	Id          string   `json:"id"`
	Hide_post   []string `json:"hide_post" gorm:"type:text[]"`
	Hide_story  []string `json:"hide_story" gorm:"type:text[]"`
	Block_user  []string `json:"block_user" gorm:"type:text[]"`
	Show_online bool     `json:"show_online" gorm:"type:text[]"`
}

type User_Data struct {
	Id      string   `json:"id"`
	Posts   []string `json:"posts" gorm:"type:text[]"`
	Stories []string `json:"stories" gorm:"type:text[]"`
	Notes   []string `json:"notes" gorm:"type:text[]"`
}

type User_Posts struct {
	Id                string    `json:"id" gorm:"primaryKey;type:uuid;default:uuid_generate_v4()"`
	Creator_id        string    `json:"creator_id"`
	Created_at        time.Time `json:"created_at"`
	Group_id          string    `json:"group_id"`
	Text_content      string    `json:"text_content"`
	Pictures_attached []string  `json:"pictures_attached" gorm:"type:text[]"`
	Liked_by          []string  `json:"liked_by" gorm:"type:text[]"`
	Comments          []string  `json:"comments" gorm:"type:text[]"`
}

type User_post_Comments struct {
	Id                string    `json:"id" gorm:"-"`
	Post_id           string    `json:"post_id"`
	Creator_id        string    `json:"creator_id"`
	Created_at        time.Time `json:"created_at"`
	Content           string    `json:"content"`
	Pictures_attached []string  `json:"pictures_attached"`
	Nested_comments   []string  `json:"nested_comments"`
	Liked_by          []string  `json:"liked_by"`
}
