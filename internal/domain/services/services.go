package services

type Services struct {
	UserService         *UserService
	ProfileService      *ProfileService
	FollowService       *FollowService
	UserSettingsService *UserSettingService
}

func NewServiceContainer(
	userService *UserService,
	profileService *ProfileService,
	followService *FollowService,
	userSettingsService *UserSettingService,
) *Services {
	return &Services{
		UserService:         userService,
		ProfileService:      profileService,
		FollowService:       followService,
		UserSettingsService: userSettingsService,
	}
}
