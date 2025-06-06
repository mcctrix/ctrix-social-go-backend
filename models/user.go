package models

import (
	"database/sql/driver"
	"fmt"
	"strings"
	"time"
)

// StringArray is a custom type to properly handle string arrays with GORM and PostgreSQL
type StringArray []string

// Value implements the driver.Valuer interface
func (a StringArray) Value() (driver.Value, error) {
	if len(a) == 0 {
		return "{}", nil
	}

	// Escape single quotes and backslashes
	escaped := make([]string, len(a))
	for i, s := range a {
		escaped[i] = strings.Replace(strings.Replace(s, "\\", "\\\\", -1), "'", "\\'", -1)
	}

	// Format as PostgreSQL array literal
	return fmt.Sprintf("{%s}", strings.Join(escaped, ",")), nil
}

// Scan implements the sql.Scanner interface
func (a *StringArray) Scan(value interface{}) error {
	var str string

	switch v := value.(type) {
	case []byte:
		str = string(v)
	case string:
		str = v
	case nil:
		*a = StringArray{}
		return nil
	default:
		return fmt.Errorf("unsupported type for StringArray: %T", value)
	}

	// Handle empty array
	if str == "{}" {
		*a = StringArray{}
		return nil
	}

	// Remove curly braces and split by comma
	trimmed := str[1 : len(str)-1]
	elements := strings.Split(trimmed, ",")

	// Unescape each element
	result := make([]string, len(elements))
	for i, e := range elements {
		result[i] = strings.Replace(strings.Replace(e, "\\'", "'", -1), "\\\\", "\\", -1)
	}

	*a = result
	return nil
}

type User_Auth struct {
	Id         string
	Email      string
	Username   string
	Password   string
	Created_at time.Time
}

type User_Profile struct {
	Id              string      `json:"id,omitempty"`
	First_name      string      `json:"first_name,omitempty"`
	Last_name       string      `json:"last_name,omitempty"`
	Profile_picture string      `json:"profile_profile,omitempty"`
	Avatar          string      `json:"avatar,omitempty"`
	Last_seen       time.Time   `json:"last_seen,omitempty"`
	Post_count      int         `json:"post_count,omitempty"`
	Followers       StringArray `json:"followers,omitempty" gorm:"type:text[]"`
	Followings      StringArray `json:"followings,omitempty" gorm:"type:text[]"`
}

type User_Additional_Info struct {
	Id              string      `json:"id,omitempty"`
	Hobbies         StringArray `json:"hobbies,omitempty" gorm:"type:text[]"`
	Family_members  StringArray `json:"family_members,omitempty" gorm:"type:text[]"`
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

type User_Data struct {
	Id      string      `json:"id,omitempty"`
	Posts   StringArray `json:"posts,omitempty" gorm:"type:text[]"`
	Stories StringArray `json:"stories,omitempty" gorm:"type:text[]"`
	Notes   StringArray `json:"notes,omitempty" gorm:"type:text[]"`
}

type User_Posts struct {
	Id                string      `json:"id,omitempty" gorm:"primaryKey;type:uuid;default:uuid_generate_v4()"`
	Creator_id        string      `json:"creator_id"`
	Created_at        time.Time   `json:"created_at"`
	Group_id          string      `json:"group_id"`
	Text_content      string      `json:"text_content"`
	Pictures_attached StringArray `json:"pictures_attached" gorm:"type:text[]"`
	Comments          StringArray `json:"comments" gorm:"type:text[]"`
}
type User_Post_Like_Table struct {
	User_id   string `json:"user_id"`
	Post_id   string `json:"post_id,omitempty"`
	Like_type string `json:"like_type"`
}

type User_post_Comments struct {
	Id                string      `json:"id" gorm:"primaryKey;type:uuid;default:uuid_generate_v4()"`
	Post_id           string      `json:"post_id"`
	Creator_id        string      `json:"creator_id"`
	Created_at        time.Time   `json:"created_at"`
	Content           string      `json:"content"`
	Pictures_attached StringArray `json:"pictures_attached,omitempty"`
	Nested_comments   StringArray `json:"nested_comments,omitempty"`
}

type User_comment_like struct {
	User_id    string `json:"user_id"`
	Comment_id string `json:"comment_id"`
	Like_type  string `json:"like_type"`
}
