package repositories

import (
	"errors"
	"fmt"

	"github.com/mcctrix/ctrix-social-go-backend/internal/domain/models"
	"github.com/mcctrix/ctrix-social-go-backend/internal/domain/repositories"
	"github.com/mcctrix/ctrix-social-go-backend/internal/pkg/auth"
	"gorm.io/gorm"
)

// PostgreSQLUserRepository is a concrete implementation of repositories.UserRepository
// for PostgreSQL database.
type PostgreSQLUserRepository struct {
	db *gorm.DB
}

// NewPostgreSQLUserRepository creates a new instance of PostgreSQLUserRepository.
func NewPostgreSQLUserRepository(db *gorm.DB) repositories.UserRepository {
	return &PostgreSQLUserRepository{db: db}
}

// FindByID retrieves a user from the database by their ID.
func (r *PostgreSQLUserRepository) FindByID(id string) (*models.User, error) {
	user := &models.User{}
	query := r.db.Table("user_auth").Where("id = ?", id)
	err := query.First(user).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("user not found")
		}
		return nil, fmt.Errorf("failed to find user by ID: %w", err)
	}
	return user, nil
}

// FindByEmail retrieves a user from the database by their email address.
func (r *PostgreSQLUserRepository) FindByEmail(email string) (*models.User, error) {
	user := &models.User{}
	query := r.db.Table("user_auth").Where("email = ?", email).Find(user)
	err := query.Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("user not found")
		}
		return nil, fmt.Errorf("failed to find user by email: %w", err)
	}
	if query.RowsAffected == 0 {
		return nil, errors.New("user not found")
	}
	return user, nil
}
func (r *PostgreSQLUserRepository) FindByUsername(username string) (*models.User, error) {
	user := &models.User{}
	query := r.db.Table("user_auth").Where("username = ?", username).Find(user)
	err := query.Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("user not found")
		}
		return nil, fmt.Errorf("failed to find user by username: %w", err)
	}
	return user, nil
}

func (r *PostgreSQLUserRepository) GenerateJwtToken(user *models.User) (*auth.TokenData, error) {
	return auth.GenerateJwtToken(user)
}

// Save creates a new user or updates an existing one in the database.
func (r *PostgreSQLUserRepository) Save(user *models.User) error {
	// Check if user exists to decide between INSERT and UPDATE
	query := r.db.Table("user_auth").Save(user)
	if err := query.Error; err != nil {
		return fmt.Errorf("failed to save user: %w", err)
	}
	return nil
}

// Delete removes a user from the database by their ID.
func (r *PostgreSQLUserRepository) Delete(id string) error {
	query := r.db.Table("user_auth").Delete(id)
	if err := query.Error; err != nil {
		return fmt.Errorf("failed to delete user: %w", err)
	}
	return nil
}
