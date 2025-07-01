package app

import (
	"github.com/mcctrix/ctrix-social-go-backend/internal/domain/services"
	"github.com/mcctrix/ctrix-social-go-backend/internal/infrastructure/database/repositories"
	"gorm.io/gorm"
)

func BuildApplicationDependencies(db *gorm.DB) *services.Services {
	userRepo := repositories.NewPostgreSQLUserRepository(db)
	userService := services.NewUserService(userRepo)

	profileRepo := repositories.NewPostgresProfileRepository(db)
	profileService := services.NewProfileService(profileRepo)

	followRepo := repositories.NewPostgresFollowRepository(db)
	followService := services.NewFollowService(followRepo)

	userSettingsRepo := repositories.NewPostgresUserSettingRepository(db)
	userSettingsService := services.NewUserSettingService(userSettingsRepo)

	additionalUserInfoRepo := repositories.NewPostgresAdditionalInfoRepository(db)
	additionalUserInfoService := services.NewAdditionalService(additionalUserInfoRepo)

	bookmarkRepo := repositories.NewPostgresBookmarkRepository(db)
	bookmarkService := services.NewBookmarkService(bookmarkRepo)

	return services.NewServiceContainer(userService, profileService, followService, userSettingsService, additionalUserInfoService, bookmarkService)
}
