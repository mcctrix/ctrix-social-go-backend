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
