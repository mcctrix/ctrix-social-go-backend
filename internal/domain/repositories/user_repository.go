package repositories

import "github.com/mcctrix/ctrix-social-go-backend/internal/domain/models"

// UserRepository defines the interface for user data operations.
type UserRepository interface {
	// FindByID retrieves a user by their unique identifier.
	FindByID(id string) (*models.User, error)

	// FindByEmail retrieves a user by their email address.
	FindByEmail(email string) (*models.User, error)

	// Save creates a new user or updates an existing one.
	// If the user's ID is empty, it's considered a new user to be created.
	Save(user *models.User) error

	// Delete removes a user from the storage by their ID.
	Delete(id string) error
}
