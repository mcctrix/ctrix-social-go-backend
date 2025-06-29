package services

import (
	"errors"
	"fmt"

	"github.com/mcctrix/ctrix-social-go-backend/internal/domain/models"
	"github.com/mcctrix/ctrix-social-go-backend/internal/domain/repositories"
)

// UserService defines the business logic for user-related operations.
type UserService struct {
	userRepo repositories.UserRepository // UserService depends on the UserRepository interface
}

// NewUserService creates a new instance of UserService.
// It takes a UserRepository interface as a dependency.
func NewUserService(userRepo repositories.UserRepository) *UserService {
	return &UserService{
		userRepo: userRepo,
	}
}

// RegisterUser handles the business logic for user registration.
func (s *UserService) RegisterUser(email, username, password string) (*models.User, error) {
	// 1. Business Rule: Check if user with this email already exists
	existingUser, err := s.userRepo.FindByEmail(email)
	if err != nil && err.Error() != "user not found" { // Assuming repository returns specific error for not found
		return nil, fmt.Errorf("failed to check existing user: %w", err)
	}
	if existingUser != nil {
		return nil, errors.New("user with this email already exists")
	}

	// 2. Create a new User domain model (which handles password hashing)
	newUser, err := models.NewUser(email, username, password)
	if err != nil {
		return nil, fmt.Errorf("failed to create new user model: %w", err)
	}

	// 3. Use the repository to persist the new user
	if err := s.userRepo.Save(newUser); err != nil {
		return nil, fmt.Errorf("failed to save new user: %w", err)
	}

	return newUser, nil
}

// GetUserByID retrieves a user by their ID.
func (s *UserService) GetUserByID(userID string) (*models.User, error) {
	user, err := s.userRepo.FindByID(userID)
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve user: %w", err)
	}
	return user, nil
}

// AuthenticateUser handles user login and authentication.
func (s *UserService) AuthenticateUser(email, password string) (*models.User, error) {
	user, err := s.userRepo.FindByEmail(email)
	if err != nil {
		return nil, fmt.Errorf("user not found: %w", err)
	}

	if err := user.ComparePassword(password); err != nil {
		return nil, errors.New("invalid credentials")
	}

	return user, nil
}
