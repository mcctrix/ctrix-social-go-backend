package models

import "time"

type User_Additional_Info struct {
	Id              string      `json:"id,omitempty"`
	Hobbies         StringArray `json:"hobbies,omitempty" gorm:"type:text[]"`
	Relation_status string      `json:"relation_status,omitempty"`
	Dob             time.Time   `json:"dob,omitempty"`
	Bio             string      `json:"bio,omitempty"`
	Gender          string      `json:"gender,omitempty"`
}
