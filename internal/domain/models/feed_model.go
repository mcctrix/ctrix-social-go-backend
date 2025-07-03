package models

type PostWithUserDetails struct {
	User_Post
	Username       string `json:"username,omitempty"`
	Avatar         string `json:"avatar,omitempty"`
	ProfilePicture string `json:"profile_picture,omitempty"`
	VerifiedUser   bool   `json:"verified_user"`
	Bio            string `json:"bio,omitempty"`
	IsLiked        bool   `json:"is_liked"`
	LikesCount     int    `json:"likes_count"`
}

type FollowRecommendation struct {
	User_Profile
	Username string `json:"username,omitempty"`
}
