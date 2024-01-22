package response

import "time"

type PhotoResponse struct {
	ID        uint   `json:"id"`
	Title     string `json:"title"`
	Caption   string `json:"caption"`
	PhotoURL  string `json:"photo_url"`
	UserID    uint   `json:"user_id"`
	CreatedAt string `json:"created_at"`
}

type PhotoUserResponse struct {
	ID        uint              `json:"id"`
	Title     string            `json:"title"`
	Caption   string            `json:"caption"`
	PhotoURL  string            `json:"photo_url"`
	UserID    uint              `json:"user_id"`
	CreatedAt time.Time         `json:"created_at"`
	UpdatedAt time.Time         `json:"updated_at"`
	User      UserPhotoResponse `json:"User"`
}

type PhotoResponseUpdate struct {
	ID        uint   `json:"id"`
	Title     string `json:"title"`
	Caption   string `json:"caption"`
	PhotoURL  string `json:"photo_url"`
	UserID    uint   `json:"user_id"`
	UpdatedAt string `json:"updated_at"`
}

type PhotoComment struct {
	ID       uint   `json:"id"`
	Title    string `json:"title"`
	Caption  string `json:"caption"`
	PhotoURL string `json:"photo_url"`
	UserID   uint   `json:"user_id"`
}
