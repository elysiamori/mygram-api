package services

import (
	"github.com/elysiamori/mygram-api/api/models"
	"github.com/elysiamori/mygram-api/api/repositories/mygram"
)

type SocialMediaServiceImpl struct {
	SocialMediaRepo mygram.SocialMediaRepositoryImpl
}

func NewSocialMediaService(socialMediaRepo mygram.SocialMediaRepositoryImpl) *SocialMediaServiceImpl {
	return &SocialMediaServiceImpl{SocialMediaRepo: socialMediaRepo}
}

func (s *SocialMediaServiceImpl) PostSocialMedia(sm *models.SocialMedia) (*models.SocialMedia, error) {
	sm, err := s.SocialMediaRepo.PostSocialMedia(sm)
	if err != nil {
		return sm, err
	}
	return sm, nil
}

func (s *SocialMediaServiceImpl) GetAllSocialMedia() ([]models.SocialMedia, error) {
	userSocialMedia, err := s.SocialMediaRepo.GetAllSocialMedia()
	if err != nil {
		return nil, err
	}
	return userSocialMedia, nil
}

func (s *SocialMediaServiceImpl) GetSocialMediaByID(id int) (*models.SocialMedia, error) {
	userSocialMedia, err := s.SocialMediaRepo.GetSocialMediaByID(id)
	if err != nil {
		return nil, err
	}

	return userSocialMedia, nil
}

func (s *SocialMediaServiceImpl) UpdateSocialMedia(sm *models.SocialMedia, newSM *models.SocialMedia) (*models.SocialMedia, error) {
	sm, err := s.SocialMediaRepo.UpdateSocialMedia(sm, newSM)
	if err != nil {
		return nil, err
	}
	return sm, nil
}

func (s *SocialMediaServiceImpl) DeleteSocialMedia(sm *models.SocialMedia) error {
	err := s.SocialMediaRepo.DeleteSocialMedia(sm)
	if err != nil {
		return err
	}
	return nil
}
