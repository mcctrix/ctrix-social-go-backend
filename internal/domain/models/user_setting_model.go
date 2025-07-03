package models

type User_Settings struct {
	Id          string      `json:"id,omitempty" gorm:"column:id;type:varchar(50);primary_key"`
	Hide_post   StringArray `json:"hide_post,omitempty" gorm:"column:hide_post;type:text[]"`
	Hide_story  StringArray `json:"hide_story,omitempty" gorm:"column:hide_story;type:text[]"`
	Block_user  StringArray `json:"block_user,omitempty" gorm:"column:block_user;type:text[]"`
	Show_online bool        `json:"show_online" gorm:"column:show_online"`
}

func (User_Settings) TableName() string {
	return "user_settings"
}
