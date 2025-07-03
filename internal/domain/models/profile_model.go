package models

import "time"

type User_Profile struct {
	Id              string    `json:"id,omitempty" gorm:"column:id;type:varchar(50);primary_key"`
	First_name      string    `json:"first_name,omitempty" gorm:"column:first_name;type:varchar(30)"`
	Last_name       string    `json:"last_name,omitempty" gorm:"column:last_name;type:varchar(30)"`
	Avatar          string    `json:"avatar,omitempty" gorm:"column:avatar;type:varchar(25)"`
	Profile_picture string    `json:"profile_profile,omitempty" gorm:"column:profile_picture;type:varchar(200)"`
	Last_seen       time.Time `json:"last_seen,omitempty" gorm:"column:last_seen"`
	Verified_user   bool      `json:"verified_user" gorm:"column:verified_user;default:false"`
}

func (User_Profile) TableName() string {
	return "user_profile"
}
