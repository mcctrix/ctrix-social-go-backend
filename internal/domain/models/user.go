package models

import (
	"time"
)

type User_Auth struct {
	Id         string    `json:"id,omitempty"`
	Email      string    `json:"email,omitempty"`
	Username   string    `json:"username,omitempty"`
	Password   string    `json:"password,omitempty"`
	Created_at time.Time `json:"created_at,omitempty"`
}

type User_Profile struct {
	Id              string    `json:"id,omitempty"`
	First_name      string    `json:"first_name,omitempty"`
	Last_name       string    `json:"last_name,omitempty"`
	Avatar          string    `json:"avatar,omitempty"`
	Profile_picture string    `json:"profile_profile,omitempty"`
	Last_seen       time.Time `json:"last_seen,omitempty"`
	Verified_user   bool      `json:"verified_user"`
}

type User_Additional_Info struct {
	Id              string      `json:"id,omitempty"`
	Hobbies         StringArray `json:"hobbies,omitempty" gorm:"type:text[]"`
	Relation_status string      `json:"relation_status,omitempty"`
	Dob             time.Time   `json:"dob,omitempty"`
	Bio             string      `json:"bio,omitempty"`
	Gender          string      `json:"gender,omitempty"`
}

type User_Settings struct {
	Id          string      `json:"id,omitempty"`
	Hide_post   StringArray `json:"hide_post,omitempty" gorm:"type:text[]"`
	Hide_story  StringArray `json:"hide_story,omitempty" gorm:"type:text[]"`
	Block_user  StringArray `json:"block_user,omitempty" gorm:"type:text[]"`
	Show_online bool        `json:"show_online" gorm:"type:text[]"`
}

type Follows struct {
	Follower_id  string    `json:"follower_id,omitempty"`
	Following_id string    `json:"following_id,omitempty"`
	Created_at   time.Time `json:"created_at,omitempty"`
}
