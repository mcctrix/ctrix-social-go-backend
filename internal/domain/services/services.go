package services

type Services struct {
	UserService               *UserService
	ProfileService            *ProfileService
	FollowService             *FollowService
	UserSettingsService       *UserSettingService
	AdditionalUserInfoService *AdditionalUserInfoService
}

func NewServiceContainer(
	userService *UserService,
	profileService *ProfileService,
	followService *FollowService,
	userSettingsService *UserSettingService,
	additionalUserInfoService *AdditionalUserInfoService,
) *Services {
	return &Services{
		UserService:               userService,
		ProfileService:            profileService,
		FollowService:             followService,
		UserSettingsService:       userSettingsService,
		AdditionalUserInfoService: additionalUserInfoService,
	}
}
