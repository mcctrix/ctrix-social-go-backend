package services

type Services struct {
	UserService    *UserService
	ProfileService *ProfileService
	FollowService  *FollowService
}

func NewServiceContainer(
	userService *UserService,
	profileService *ProfileService,
	followService *FollowService,
) *Services {
	return &Services{
		UserService:    userService,
		ProfileService: profileService,
		FollowService:  followService,
	}
}
