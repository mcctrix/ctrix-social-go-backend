package models

import (
	"time"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

// User represents the core user authentication and identity.
type User struct {
	Id         string    `json:"id,omitempty"`
	Email      string    `json:"email,omitempty"`
	Username   string    `json:"username,omitempty"`
	Password   string    `json:"-"` // Exclude from JSON output
	Created_at time.Time `json:"created_at,omitempty"`
}

// NewUser creates a new User instance with a hashed password.
func NewUser(email, username, password string) (*User, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}
	userID := uuid.New().String()
	return &User{
		Id:         userID,
		Email:      email,
		Username:   username,
		Password:   string(hashedPassword),
		Created_at: time.Now(),
	}, nil
}

// ComparePassword compares a plaintext password with the stored hash.
func (u *User) ComparePassword(password string) error {
	return bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
}

// --- Other User-related structs (kept as is for now) ---

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
