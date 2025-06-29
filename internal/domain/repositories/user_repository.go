package repositories

import (
	"github.com/mcctrix/ctrix-social-go-backend/internal/domain/models"
	"github.com/mcctrix/ctrix-social-go-backend/internal/pkg/auth"
)

type UserRepository interface {
	// FindByID retrieves a user by their unique identifier.
	FindByID(id string) (*models.User, error)

	// FindByEmail retrieves a user by their email address.
	FindByEmail(email string) (*models.User, error)

	// FindByUsername retrieves a user by their username.
	FindByUsername(username string) (*models.User, error)

	GenerateJwtToken(user *models.User) (*auth.TokenData, error)

	// Save creates a new user or updates an existing one.
	// If the user's ID is empty, it's considered a new user to be created.
	Save(user *models.User) error

	// Delete removes a user from the storage by their ID.
	Delete(id string) error
}
