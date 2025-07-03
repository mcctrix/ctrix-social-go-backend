package models

import (
	"time"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

// User represents the core user authentication and identity.
type User struct {
	Id         string    `json:"id,omitempty" gorm:"column:id;type:varchar(50);primary_key"`
	Email      string    `json:"email,omitempty" gorm:"column:email;type:varchar(50);unique;not null"`
	Username   string    `json:"username,omitempty" gorm:"column:username;type:varchar(30);unique;not null"`
	Password   string    `json:"-" gorm:"column:password;type:varchar(60);not null"` // Exclude from JSON output
	Created_at time.Time `json:"created_at,omitempty" gorm:"column:created_at;default:now()"`
	Updated_at time.Time `json:"updated_at,omitempty" gorm:"column:updated_at;default:now()"`
}

func (User) TableName() string {
	return "user_auth"
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
