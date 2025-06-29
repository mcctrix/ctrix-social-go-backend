package models

type User_Settings struct {
	Id          string      `json:"id,omitempty"`
	Hide_post   StringArray `json:"hide_post,omitempty" gorm:"type:text[]"`
	Hide_story  StringArray `json:"hide_story,omitempty" gorm:"type:text[]"`
	Block_user  StringArray `json:"block_user,omitempty" gorm:"type:text[]"`
	Show_online bool        `json:"show_online" gorm:"type:text[]"`
}
