package mygram

import (
	"github.com/elysiamori/mygram-api/api/models"
	"gorm.io/gorm"
)

type PhotoRepositoryImpl struct {
	DB *gorm.DB
}

func NewPhotoRepository(db *gorm.DB) *PhotoRepositoryImpl {
	return &PhotoRepositoryImpl{DB: db}
}

func (r *PhotoRepositoryImpl) UploadPhoto(photo *models.Photo) (*models.Photo, error) {
	err := r.DB.Create(&photo).Error
	if err != nil {
		return photo, err
	}
	return photo, nil
}

func (r *PhotoRepositoryImpl) GetAllPhoto() ([]models.Photo, error) {
	var userphotos []models.Photo
	err := r.DB.Find(&userphotos).Error
	if err != nil {
		return []models.Photo{}, err
	}
	return userphotos, nil
}

func (r *PhotoRepositoryImpl) GetPhotoByID(id int) (*models.Photo, error) {
	var photo models.Photo
	err := r.DB.Where("id = ?", id).Take(&photo).Error
	if err != nil {
		return nil, err
	}
	return &photo, nil
}

func (r *PhotoRepositoryImpl) UpdatePhoto(photo *models.Photo, newPhoto *models.Photo) (*models.Photo, error) {
	err := r.DB.Model(&photo).Updates(newPhoto).Error
	if err != nil {
		return nil, err
	}
	return photo, nil
}

func (r *PhotoRepositoryImpl) DeletePhoto(id uint) error {
	var photo models.Photo
	if err := r.DB.Preload("Comment").First(&photo, id).Error; err != nil {
		return err
	}

	for _, comment := range photo.Comment {
		if err := r.DB.Delete(&comment).Error; err != nil {
			return err
		}
	}

	if err := r.DB.Delete(&photo).Error; err != nil {
		return err
	}

	return nil
}
