package repositories

type FollowRepository interface {
	CreateFollow(follower_id, following_id string) error
	IsFollowing(follower_id, following_id string) bool
	UnFollow(follower_id, following_id string) error
	CountFollowAndFollowing(userID string) (struct {
		followCount    int
		followingCount int
	}, error)
}
