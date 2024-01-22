package mygram

import (
	"github.com/elysiamori/mygram-api/api/models"
	"gorm.io/gorm"
)

type SocialMediaRepositoryImpl struct {
	DB *gorm.DB
}

func NewSMRepository(db *gorm.DB) *SocialMediaRepositoryImpl {
	return &SocialMediaRepositoryImpl{DB: db}
}

func (r *SocialMediaRepositoryImpl) PostSocialMedia(sm *models.SocialMedia) (*models.SocialMedia, error) {
	err := r.DB.Create(&sm).Error
	if err != nil {
		return sm, err
	}
	return sm, nil
}

func (r *SocialMediaRepositoryImpl) GetAllSocialMedia() ([]models.SocialMedia, error) {
	var userSocialMedia []models.SocialMedia
	err := r.DB.Find(&userSocialMedia).Error
	if err != nil {
		return []models.SocialMedia{}, err
	}
	return userSocialMedia, nil
}

func (r *SocialMediaRepositoryImpl) GetSocialMediaByID(id int) (*models.SocialMedia, error) {
	var sm models.SocialMedia
	err := r.DB.Where("id = ?", id).Take(&sm).Error
	if err != nil {
		return nil, err
	}
	return &sm, nil
}

func (r *SocialMediaRepositoryImpl) UpdateSocialMedia(sm *models.SocialMedia, newSM *models.SocialMedia) (*models.SocialMedia, error) {
	err := r.DB.Model(&sm).Updates(newSM).Error
	if err != nil {
		return nil, err
	}
	return sm, nil
}

func (r *SocialMediaRepositoryImpl) DeleteSocialMedia(sm *models.SocialMedia) error {
	err := r.DB.Delete(&sm).Error
	if err != nil {
		return err
	}
	return nil
}
