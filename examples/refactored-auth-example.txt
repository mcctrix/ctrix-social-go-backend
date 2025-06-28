// Example of how the auth controller would look in the improved structure
// This demonstrates the separation of concerns and better organization

// =============================================================================
// 1. DOMAIN MODELS (internal/domain/models/user.go)
// =============================================================================

package models

import (
    "time"
    "github.com/google/uuid"
)

type User struct {
    ID        string    `json:"id" db:"id"`
    Email     string    `json:"email" db:"email"`
    Username  string    `json:"username" db:"username"`
    Password  string    `json:"-" db:"password"` // "-" means don't include in JSON
    CreatedAt time.Time `json:"created_at" db:"created_at"`
    UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
}

func NewUser(email, username, password string) *User {
    return &User{
        ID:        uuid.New().String(),
        Email:     email,
        Username:  username,
        Password:  password,
        CreatedAt: time.Now(),
        UpdatedAt: time.Now(),
    }
}

// =============================================================================
// 2. REPOSITORY INTERFACE (internal/domain/repositories/user_repository.go)
// =============================================================================

package repositories

import (
    "context"
    "github.com/mcctrix/ctrix-social-go-backend/internal/domain/models"
)

type UserRepository interface {
    Create(ctx context.Context, user *models.User) error
    GetByEmail(ctx context.Context, email string) (*models.User, error)
    GetByUsername(ctx context.Context, username string) (*models.User, error)
    GetByID(ctx context.Context, id string) (*models.User, error)
    Update(ctx context.Context, user *models.User) error
    Delete(ctx context.Context, id string) error
}

// =============================================================================
// 3. BUSINESS LOGIC SERVICE (internal/domain/services/auth_service.go)
// =============================================================================

package services

import (
    "context"
    "errors"
    "time"
    
    "github.com/mcctrix/ctrix-social-go-backend/internal/domain/models"
    "github.com/mcctrix/ctrix-social-go-backend/internal/domain/repositories"
    "github.com/mcctrix/ctrix-social-go-backend/internal/pkg/auth"
)

type AuthService struct {
    userRepo repositories.UserRepository
    jwtUtil  auth.JWTUtil
}

func NewAuthService(userRepo repositories.UserRepository, jwtUtil auth.JWTUtil) *AuthService {
    return &AuthService{
        userRepo: userRepo,
        jwtUtil:  jwtUtil,
    }
}

type SignUpRequest struct {
    Email    string `json:"email" validate:"required,email"`
    Username string `json:"username" validate:"required,min=3,max=30"`
    Password string `json:"password" validate:"required,min=8"`
}

type SignUpResponse struct {
    User  *models.User `json:"user"`
    Token string       `json:"token"`
}

func (s *AuthService) SignUp(ctx context.Context, req SignUpRequest) (*SignUpResponse, error) {
    // Check if user already exists
    existingUser, _ := s.userRepo.GetByEmail(ctx, req.Email)
    if existingUser != nil {
        return nil, errors.New("user with this email already exists")
    }

    existingUser, _ = s.userRepo.GetByUsername(ctx, req.Username)
    if existingUser != nil {
        return nil, errors.New("username already taken")
    }

    // Hash password
    hashedPassword, err := auth.HashPassword(req.Password)
    if err != nil {
        return nil, err
    }

    // Create new user
    user := models.NewUser(req.Email, req.Username, hashedPassword)
    
    if err := s.userRepo.Create(ctx, user); err != nil {
        return nil, err
    }

    // Generate JWT token
    token, err := s.jwtUtil.GenerateToken(user.ID)
    if err != nil {
        return nil, err
    }

    return &SignUpResponse{
        User:  user,
        Token: token,
    }, nil
}

type LoginRequest struct {
    Username string `json:"username" validate:"required"`
    Password string `json:"password" validate:"required"`
}

type LoginResponse struct {
    User  *models.User `json:"user"`
    Token string       `json:"token"`
}

func (s *AuthService) Login(ctx context.Context, req LoginRequest) (*LoginResponse, error) {
    // Find user by username
    user, err := s.userRepo.GetByUsername(ctx, req.Username)
    if err != nil {
        return nil, errors.New("invalid credentials")
    }

    // Verify password
    if !auth.VerifyPassword(req.Password, user.Password) {
        return nil, errors.New("invalid credentials")
    }

    // Generate JWT token
    token, err := s.jwtUtil.GenerateToken(user.ID)
    if err != nil {
        return nil, err
    }

    return &LoginResponse{
        User:  user,
        Token: token,
    }, nil
}

// =============================================================================
// 4. HTTP HANDLER (internal/api/handlers/auth/handler.go)
// =============================================================================

package auth

import (
    "github.com/gofiber/fiber/v3"
    "github.com/mcctrix/ctrix-social-go-backend/internal/domain/services"
    "github.com/mcctrix/ctrix-social-go-backend/internal/pkg/errors"
    "github.com/mcctrix/ctrix-social-go-backend/internal/pkg/response"
)

type AuthHandler struct {
    authService services.AuthService
}

func NewAuthHandler(authService services.AuthService) *AuthHandler {
    return &AuthHandler{
        authService: authService,
    }
}

func (h *AuthHandler) SignUp(c fiber.Ctx) error {
    var req services.SignUpRequest
    if err := c.BodyParser(&req); err != nil {
        return response.BadRequest(c, "Invalid request body")
    }

    // Validate request
    if err := req.Validate(); err != nil {
        return response.BadRequest(c, err.Error())
    }

    // Check if user is already logged in
    if token := c.Cookies("auth_token"); token != "" {
        return response.BadRequest(c, "User already logged in")
    }

    // Call business logic
    result, err := h.authService.SignUp(c.Context(), req)
    if err != nil {
        return response.BadRequest(c, err.Error())
    }

    // Set cookie
    h.setAuthCookie(c, result.Token)

    return response.Success(c, result)
}

func (h *AuthHandler) Login(c fiber.Ctx) error {
    var req services.LoginRequest
    if err := c.BodyParser(&req); err != nil {
        return response.BadRequest(c, "Invalid request body")
    }

    // Validate request
    if err := req.Validate(); err != nil {
        return response.BadRequest(c, err.Error())
    }

    // Check if user is already logged in
    if token := c.Cookies("auth_token"); token != "" {
        return response.BadRequest(c, "User already logged in")
    }

    // Call business logic
    result, err := h.authService.Login(c.Context(), req)
    if err != nil {
        return response.Unauthorized(c, err.Error())
    }

    // Set cookie
    h.setAuthCookie(c, result.Token)

    return response.Success(c, result)
}

func (h *AuthHandler) Logout(c fiber.Ctx) error {
    c.ClearCookie("auth_token")
    return response.Success(c, fiber.Map{"message": "User logged out successfully"})
}

func (h *AuthHandler) setAuthCookie(c fiber.Ctx, token string) {
    // Cookie configuration based on environment
    secure := os.Getenv("APP_ENV") != "dev"
    domain := ""
    if os.Getenv("APP_ENV") != "dev" {
        domain = "ctrix-social.vercel.app"
    }

    c.Cookie(&fiber.Cookie{
        Name:     "auth_token",
        Value:    token,
        Path:     "/",
        HTTPOnly: true,
        Secure:   secure,
        Domain:   domain,
        SameSite: "Lax",
        Expires:  time.Now().Add(24 * time.Hour),
    })
}

// =============================================================================
// 5. REQUEST VALIDATION (internal/api/handlers/auth/validator.go)
// =============================================================================

package auth

import (
    "github.com/go-playground/validator/v10"
)

var validate = validator.New()

func (r *services.SignUpRequest) Validate() error {
    return validate.Struct(r)
}

func (r *services.LoginRequest) Validate() error {
    return validate.Struct(r)
}

// =============================================================================
// 6. REPOSITORY IMPLEMENTATION (internal/infrastructure/database/repositories/user_repository.go)
// =============================================================================

package repositories

import (
    "context"
    "database/sql"
    "time"
    
    "github.com/mcctrix/ctrix-social-go-backend/internal/domain/models"
    "github.com/mcctrix/ctrix-social-go-backend/internal/domain/repositories"
)

type userRepository struct {
    db *sql.DB
}

func NewUserRepository(db *sql.DB) repositories.UserRepository {
    return &userRepository{db: db}
}

func (r *userRepository) Create(ctx context.Context, user *models.User) error {
    query := `
        INSERT INTO user_auth (id, email, username, password, created_at, updated_at)
        VALUES ($1, $2, $3, $4, $5, $6)
    `
    
    _, err := r.db.ExecContext(ctx, query,
        user.ID, user.Email, user.Username, user.Password,
        user.CreatedAt, user.UpdatedAt,
    )
    
    return err
}

func (r *userRepository) GetByEmail(ctx context.Context, email string) (*models.User, error) {
    query := `
        SELECT id, email, username, password, created_at, updated_at
        FROM user_auth WHERE email = $1
    `
    
    user := &models.User{}
    err := r.db.QueryRowContext(ctx, query, email).Scan(
        &user.ID, &user.Email, &user.Username, &user.Password,
        &user.CreatedAt, &user.UpdatedAt,
    )
    
    if err == sql.ErrNoRows {
        return nil, nil
    }
    
    return user, err
}

func (r *userRepository) GetByUsername(ctx context.Context, username string) (*models.User, error) {
    query := `
        SELECT id, email, username, password, created_at, updated_at
        FROM user_auth WHERE username = $1
    `
    
    user := &models.User{}
    err := r.db.QueryRowContext(ctx, query, username).Scan(
        &user.ID, &user.Email, &user.Username, &user.Password,
        &user.CreatedAt, &user.UpdatedAt,
    )
    
    if err == sql.ErrNoRows {
        return nil, nil
    }
    
    return user, err
}

func (r *userRepository) GetByID(ctx context.Context, id string) (*models.User, error) {
    query := `
        SELECT id, email, username, password, created_at, updated_at
        FROM user_auth WHERE id = $1
    `
    
    user := &models.User{}
    err := r.db.QueryRowContext(ctx, query, id).Scan(
        &user.ID, &user.Email, &user.Username, &user.Password,
        &user.CreatedAt, &user.UpdatedAt,
    )
    
    if err == sql.ErrNoRows {
        return nil, nil
    }
    
    return user, err
}

func (r *userRepository) Update(ctx context.Context, user *models.User) error {
    user.UpdatedAt = time.Now()
    
    query := `
        UPDATE user_auth 
        SET email = $2, username = $3, password = $4, updated_at = $5
        WHERE id = $1
    `
    
    _, err := r.db.ExecContext(ctx, query,
        user.ID, user.Email, user.Username, user.Password, user.UpdatedAt,
    )
    
    return err
}

func (r *userRepository) Delete(ctx context.Context, id string) error {
    query := `DELETE FROM user_auth WHERE id = $1`
    _, err := r.db.ExecContext(ctx, query, id)
    return err
}

// =============================================================================
// 7. ROUTES (internal/api/routes/auth.go)
// =============================================================================

package routes

import (
    "github.com/gofiber/fiber/v3"
    "github.com/mcctrix/ctrix-social-go-backend/internal/api/handlers/auth"
)

func SetupAuthRoutes(app fiber.Router, authHandler *auth.AuthHandler) {
    app.Post("/signup", authHandler.SignUp)
    app.Post("/login", authHandler.Login)
    app.Post("/logout", authHandler.Logout)
    app.Post("/refresh", authHandler.RefreshToken)
}

// =============================================================================
// 8. MAIN APPLICATION SETUP (cmd/server/main.go)
// =============================================================================

package main

import (
    "database/sql"
    "log"
    
    "github.com/gofiber/fiber/v3"
    "github.com/mcctrix/ctrix-social-go-backend/internal/api/handlers/auth"
    "github.com/mcctrix/ctrix-social-go-backend/internal/api/routes"
    "github.com/mcctrix/ctrix-social-go-backend/internal/config"
    "github.com/mcctrix/ctrix-social-go-backend/internal/domain/services"
    "github.com/mcctrix/ctrix-social-go-backend/internal/infrastructure/database/repositories"
    "github.com/mcctrix/ctrix-social-go-backend/internal/pkg/auth"
)

func main() {
    // Load configuration
    cfg, err := config.Load()
    if err != nil {
        log.Fatal("Failed to load config:", err)
    }

    // Initialize database
    db, err := initDatabase(cfg.Database)
    if err != nil {
        log.Fatal("Failed to connect to database:", err)
    }
    defer db.Close()

    // Initialize dependencies
    userRepo := repositories.NewUserRepository(db)
    jwtUtil := auth.NewJWTUtil(cfg.JWT.Secret, cfg.JWT.Expiration)
    authService := services.NewAuthService(userRepo, jwtUtil)
    authHandler := auth.NewAuthHandler(authService)

    // Setup Fiber app
    app := fiber.New()

    // Setup routes
    authGroup := app.Group("/api/auth")
    routes.SetupAuthRoutes(authGroup, authHandler)

    // Start server
    log.Fatal(app.Listen(":" + cfg.Server.Port))
}

func initDatabase(cfg config.DatabaseConfig) (*sql.DB, error) {
    // Database connection logic here
    return nil, nil
} 