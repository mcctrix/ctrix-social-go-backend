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
	Last_seen       time.Time `json:"last_seen"`
	Post_count      int       `json:"post_count"`
	Followers       []string  `json:"followers"`
	Followings      []string  `json:"followings"`
}

type User_Additional_Info struct {
	Id              string   `json:"id"`
	Hobbies         []string `json:"hobbies"`
	Family_members  []string
	Relation_status string
	Dob             time.Time
	Bio             string
	Gender          string
}

type User_Settings struct {
	Id          string
	Hide_post   []string
	Hide_story  []string
	Block_user  []string
	Show_online bool
}

type User_Data struct {
	Id      string
	Posts   []string
	Stories []string
	Notes   []string
}

type User_Posts struct {
	Id                string
	Creator_id        string
	Created_at        time.Time
	Group_id          string
	Text_content      string
	Pictures_attached []string
	Liked_by          []string
	Comments          []string
}

type User_post_Comments struct {
	Id                string
	Post_id           string
	Creator_id        string
	Created_at        time.Time
	Content           string
	Pictures_attached []string
	Nested_comments   []string
	Liked_by          []string
}
