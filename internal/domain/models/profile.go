package models

import "time"

type User_Profile struct {
	Id              string    `json:"id,omitempty"`
	First_name      string    `json:"first_name,omitempty"`
	Last_name       string    `json:"last_name,omitempty"`
	Avatar          string    `json:"avatar,omitempty"`
	Profile_picture string    `json:"profile_profile,omitempty"`
	Last_seen       time.Time `json:"last_seen,omitempty"`
	Verified_user   bool      `json:"verified_user"`
}
