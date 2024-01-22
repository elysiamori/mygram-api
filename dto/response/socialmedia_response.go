package response

type SocialMediaResponse struct {
	ID             uint   `json:"id"`
	Name           string `json:"name"`
	SocialMediaURL string `json:"social_media_url"`
	UserID         uint   `json:"user_id"`
	CreatedAt      string `json:"created_at"`
}

type SocialMediaUserResponse struct {
	ID             uint                    `json:"id"`
	Name           string                  `json:"name"`
	SocialMediaURL string                  `json:"social_media_url"`
	UserID         uint                    `json:"user_id"`
	CreatedAt      string                  `json:"created_at"`
	UpdatedAt      string                  `json:"updated_at"`
	User           UserSocialMediaResponse `json:"user"`
}

type UserSocialMediaResponse struct {
	ID       uint   `json:"id"`
	Username string `json:"username"`
}
