package app

import (
	"github.com/mcctrix/ctrix-social-go-backend/internal/domain/services"
	"github.com/mcctrix/ctrix-social-go-backend/internal/infrastructure/database/repositories"
	"gorm.io/gorm"
)

func BuildApplicationDependencies(db *gorm.DB) *services.Services {
	userRepo := repositories.NewPostgreSQLUserRepository(db)
	userService := services.NewUserService(userRepo)

	return services.NewServiceContainer(userService)
}
