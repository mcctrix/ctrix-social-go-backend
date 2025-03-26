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
	Avatar string    `json:"avatar"`
	Last_seen       time.Time `json:"last_seen"`
	Post_count      int       `json:"post_count"`
	Followers       []string  `json:"followers"`
	Followings      []string  `json:"followings"`
}

type User_Additional_Info struct {
	Id              string    `json:"id"`
	Hobbies         []string  `json:"hobbies"`
	Family_members  []string  `json:"family_members"`
	Relation_status string    `json:"relation_status"`
	Dob             time.Time `json:"dob"`
	Bio             string    `json:"bio"`
	Gender          string    `json:"gender"`
}

type User_Settings struct {
	Id          string   `json:"id"`
	Hide_post   []string `json:"hide_post"`
	Hide_story  []string `json:"hide_story"`
	Block_user  []string `json:"block_user"`
	Show_online bool     `json:"show_online"`
}

type User_Data struct {
	Id      string   `json:"id"`
	Posts   []string `json:"posts"`
	Stories []string `json:"stories"`
	Notes   []string `json:"notes"`
}

type User_Posts struct {
	Id                string    `json:"id"`
	Creator_id        string    `json:"creator_id"`
	Created_at        time.Time `json:"creator_at"`
	Group_id          string    `json:"group_id"`
	Text_content      string    `json:"text_content"`
	Pictures_attached []string  `json:"pictures_attached"`
	Liked_by          []string  `json:"liked_by"`
	Comments          []string  `json:"comments"`
}

type User_post_Comments struct {
	Id                string    `json:"id"`
	Post_id           string    `json:"post_id"`
	Creator_id        string    `json:"creator_id"`
	Created_at        time.Time `json:"created_at"`
	Content           string    `json:"content"`
	Pictures_attached []string  `json:"pictures_attached"`
	Nested_comments   []string  `json:"nested_comments"`
	Liked_by          []string  `json:"liked_by"`
}





