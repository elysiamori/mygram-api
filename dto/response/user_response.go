package response

type RegisterResponse struct {
	Age      uint   `json:"age"`
	Email    string `json:"email"`
	ID       uint   `json:"id"`
	Username string `json:"username"`
}

type UpdateResponse struct {
	ID       uint   `json:"id"`
	Email    string `json:"email"`
	Username string `json:"username"`
	Age      uint   `json:"age"`
	UpdateAt string `json:"update_at"`
}

type DeleteResponse struct {
	Message string `json:"message"`
}

type UserPhotoResponse struct {
	Email    string `json:"email"`
	Username string `json:"username"`
}

type UserComment struct {
	ID       uint   `json:"id"`
	Email    string `json:"email"`
	Username string `json:"username"`
}
