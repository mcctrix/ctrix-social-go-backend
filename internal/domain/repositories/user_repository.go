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

	//  creates a new user
	Create(user *models.User) error

	// Delete removes a user from the storage by their ID.
	Delete(id string) error
}
