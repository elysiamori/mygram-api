package request

import "github.com/elysiamori/mygram-api/api/models"

type PhotoRequest struct {
	Title    string `json:"title" validate:"required"`
	Caption  string `json:"caption" validate:"required"`
	PhotoURL string `json:"photo_url" validate:"required"`
}

type UpdatePhotoRequest struct {
	Title    string `json:"title" validate:"required"`
	Caption  string `json:"caption" validate:"required"`
	PhotoURL string `json:"photo_url" validate:"required"`
}

func (p *PhotoRequest) ToDomain() *models.Photo {
	return &models.Photo{
		Title:    p.Title,
		Caption:  p.Caption,
		PhotoURL: p.PhotoURL,
	}
}
