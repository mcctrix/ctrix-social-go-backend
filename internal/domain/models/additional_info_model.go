package models

import "time"

type User_Additional_Info struct {
	Id              string      `json:"id,omitempty" gorm:"column:id;type:varchar(50);primary_key"`
	Hobbies         StringArray `json:"hobbies,omitempty" gorm:"column:hobbies;type:text[]"`
	Relation_status string      `json:"relation_status,omitempty" gorm:"column:relation_status;type:varchar(12)"`
	Dob             time.Time   `json:"dob,omitempty" gorm:"column:dob;type:date"`
	Bio             string      `json:"bio,omitempty" gorm:"column:bio;type:varchar(250)"`
	Gender          string      `json:"gender,omitempty" gorm:"column:gender;type:varchar(6)"`
}

func (User_Additional_Info) TableName() string {
	return "user_additional_info"
}
