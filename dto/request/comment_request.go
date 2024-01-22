package request

type CommentRequest struct {
	Message string `json:"message"`
	PhotoID uint   `json:"photo_id"`
}

type CommentUpdateRequest struct {
	Message string `json:"message"`
}
