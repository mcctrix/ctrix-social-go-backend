package services

type Services struct {
	UserService               *UserService
	ProfileService            *ProfileService
	FollowService             *FollowService
	UserSettingsService       *UserSettingService
	AdditionalUserInfoService *AdditionalUserInfoService
	BookmarkService           *BookmarkService
	CommentService            *CommentService
	CommentReactionService    *CommentReactionService
}

func NewServiceContainer(
	userService *UserService,
	profileService *ProfileService,
	followService *FollowService,
	userSettingsService *UserSettingService,
	additionalUserInfoService *AdditionalUserInfoService,
	bookmarkService *BookmarkService,
	commentService *CommentService,
	commentReactionService *CommentReactionService,
) *Services {
	return &Services{
		UserService:               userService,
		ProfileService:            profileService,
		FollowService:             followService,
		UserSettingsService:       userSettingsService,
		AdditionalUserInfoService: additionalUserInfoService,
		BookmarkService:           bookmarkService,
		CommentService:            commentService,
		CommentReactionService:    commentReactionService,
	}
}
