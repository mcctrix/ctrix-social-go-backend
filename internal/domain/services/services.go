package services

type Services struct {
	UserService *UserService
}

func NewServiceContainer(
	userService *UserService,
) *Services {
	return &Services{
		UserService: userService,
	}
}
