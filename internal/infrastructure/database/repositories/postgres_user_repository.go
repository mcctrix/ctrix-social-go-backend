package repositories

import (
	"database/sql"
	"errors"
	"fmt"

	"github.com/mcctrix/ctrix-social-go-backend/internal/domain/models"
	"github.com/mcctrix/ctrix-social-go-backend/internal/domain/repositories"
)

// PostgreSQLUserRepository is a concrete implementation of repositories.UserRepository
// for PostgreSQL database.
type PostgreSQLUserRepository struct {
	db *sql.DB
}

// NewPostgreSQLUserRepository creates a new instance of PostgreSQLUserRepository.
func NewPostgreSQLUserRepository(db *sql.DB) repositories.UserRepository {
	return &PostgreSQLUserRepository{db: db}
}

// FindByID retrieves a user from the database by their ID.
func (r *PostgreSQLUserRepository) FindByID(id string) (*models.User, error) {
	user := &models.User{}
	query := `SELECT id, email, username, password, created_at FROM users WHERE id = $1`
	err := r.db.QueryRow(query, id).Scan(&user.Id, &user.Email, &user.Username, &user.Password, &user.Created_at)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errors.New("user not found")
		}
		return nil, fmt.Errorf("failed to find user by ID: %w", err)
	}
	return user, nil
}

// FindByEmail retrieves a user from the database by their email address.
func (r *PostgreSQLUserRepository) FindByEmail(email string) (*models.User, error) {
	user := &models.User{}
	query := `SELECT id, email, username, password, created_at FROM users WHERE email = $1`
	err := r.db.QueryRow(query, email).Scan(&user.Id, &user.Email, &user.Username, &user.Password, &user.Created_at)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errors.New("user not found")
		}
		return nil, fmt.Errorf("failed to find user by email: %w", err)
	}
	return user, nil
}

// Save creates a new user or updates an existing one in the database.
func (r *PostgreSQLUserRepository) Save(user *models.User) error {
	// Check if user exists to decide between INSERT and UPDATE
	var exists bool
	checkQuery := `SELECT EXISTS(SELECT 1 FROM users WHERE id = $1)`
	err := r.db.QueryRow(checkQuery, user.Id).Scan(&exists)
	if err != nil {
		return fmt.Errorf("failed to check user existence: %w", err)
	}

	if exists {
		// Update existing user
		updateQuery := `UPDATE users SET email = $1, username = $2, password = $3 WHERE id = $4`
		_, err = r.db.Exec(updateQuery, user.Email, user.Username, user.Password, user.Id)
		if err != nil {
			return fmt.Errorf("failed to update user: %w", err)
		}
	} else {
		// Insert new user
		insertQuery := `INSERT INTO users (id, email, username, password, created_at) VALUES ($1, $2, $3, $4, $5)`
		_, err = r.db.Exec(insertQuery, user.Id, user.Email, user.Username, user.Password, user.Created_at)
		if err != nil {
			return fmt.Errorf("failed to insert user: %w", err)
		}
	}
	return nil
}

// Delete removes a user from the database by their ID.
func (r *PostgreSQLUserRepository) Delete(id string) error {
	query := `DELETE FROM users WHERE id = $1`
	result, err := r.db.Exec(query, id)
	if err != nil {
		return fmt.Errorf("failed to delete user: %w", err)
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to get rows affected: %w", err)
	}
	if rowsAffected == 0 {
		return errors.New("user not found for deletion")
	}
	return nil
}
