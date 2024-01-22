package request

type RegisterRequest struct {
	Age      uint   `json:"age" validate:"required,numeric"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
	Username string `json:"username" validate:"required"`
}

type LoginRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=8"`
}

type UpdateRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Username string `json:"username" validate:"required"`
}

type LoginRequestUsername struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required,min=8"`
}
