package response

type CommentResponse struct {
	ID        uint   `json:"id"`
	Message   string `json:"message"`
	PhotoID   uint   `json:"photo_id"`
	UserID    uint   `json:"user_id"`
	CreatedAt string `json:"created_at"`
}

type GetCommentsResponse struct {
	ID        uint         `json:"id"`
	Message   string       `json:"message"`
	PhotoID   uint         `json:"photo_id"`
	UserID    uint         `json:"user_id"`
	CreatedAt string       `json:"created_at"`
	User      UserComment  `json:"User"`
	Photo     PhotoComment `json:"Photo"`
}

type CommentUpdate struct {
	ID        uint   `json:"id"`
	Message   string `json:"message"`
	PhotoID   uint   `json:"photo_id"`
	UserID    uint   `json:"user_id"`
	UpdatedAt string `json:"updated_at"`
}
