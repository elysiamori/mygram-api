package services

import (
	"github.com/elysiamori/mygram-api/api/models"
	"github.com/elysiamori/mygram-api/api/repositories/mygram"
	"github.com/elysiamori/mygram-api/dto/request"
	"github.com/elysiamori/mygram-api/dto/response"
)

type PhotoServiceImpl struct {
	PhotoRepo mygram.PhotoRepositoryImpl
}

func NewPhotoService(photoRepo mygram.PhotoRepositoryImpl) *PhotoServiceImpl {
	return &PhotoServiceImpl{PhotoRepo: photoRepo}
}

func (s *PhotoServiceImpl) UploadPhoto(photo *models.Photo) (*models.Photo, error) {
	photo, err := s.PhotoRepo.UploadPhoto(photo)
	if err != nil {
		return photo, err
	}
	return photo, nil
}

func (s *PhotoServiceImpl) GetAllPhoto() ([]models.Photo, error) {
	userphotos, err := s.PhotoRepo.GetAllPhoto()
	if err != nil {
		return nil, err
	}
	return userphotos, nil
}

func (s *PhotoServiceImpl) GetPhotoByID(id int) (*models.Photo, error) {
	photo, err := s.PhotoRepo.GetPhotoByID(id)
	if err != nil {
		return nil, err
	}
	return photo, nil
}

func (s *PhotoServiceImpl) UpdatePhoto(photoID uint, photo *request.UpdatePhotoRequest) (*response.PhotoResponseUpdate, error) {
	oldPhoto, err := s.PhotoRepo.GetPhotoByID(int(photoID))
	if err != nil {
		return nil, err
	}

	newPhoto := models.Photo{
		Title:    photo.Title,
		Caption:  photo.Caption,
		PhotoURL: photo.PhotoURL,
	}

	photoUpdated, err := s.PhotoRepo.UpdatePhoto(oldPhoto, &newPhoto)
	if err != nil {
		return nil, err
	}

	response := &response.PhotoResponseUpdate{
		ID:        photoUpdated.ID,
		Title:     photoUpdated.Title,
		Caption:   photoUpdated.Caption,
		PhotoURL:  photoUpdated.PhotoURL,
		UserID:    photoUpdated.UserID,
		UpdatedAt: photoUpdated.UpdatedAt.Format("2006-01-02 15:04:05"),
	}

	return response, nil
}

func (s *PhotoServiceImpl) DeletePhoto(photoID uint) error {
	photo_ID, err := s.PhotoRepo.GetPhotoByID(int(photoID))
	if err != nil {
		return err
	}

	if photo_ID.ID == 0 {
		return err
	}

	err = s.PhotoRepo.DeletePhoto(photoID)
	if err != nil {
		return err
	}

	return nil
}
