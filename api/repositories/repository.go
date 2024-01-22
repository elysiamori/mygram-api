package repositories

import "github.com/elysiamori/mygram-api/api/models"

type UserRepository interface {
	RegisterUser(user *models.User) (*models.User, error)
	LoginUser(email, password string) (*models.User, error)
	LoginWithUsername(username, password string) (*models.User, error)
	BeforeSaveUser(user *models.User) error
	GetUserByID(id uint) (*models.User, error)
	GetUserByEmail(email string) (*models.User, error)
	GetUserByUsername(username string) (*models.User, error)
	UpdateUser(user *models.User, newUser *models.User) (*models.User, error)
	DeleteUser(id uint) error
}

type PhotoRepository interface {
	UploadPhoto(photo *models.Photo) (*models.Photo, error)
	GetAllPhoto() ([]*models.Photo, error)
	GetPhotoByID(id uint) (*models.Photo, error)
	UpdatePhoto(photo *models.Photo, newPhoto *models.Photo) (*models.Photo, error)
	DeletePhoto(id uint) error
}

type CommentRepository interface {
	PostComment(comment *models.Comment) (*models.Comment, error)
	GetAllComment() ([]*models.Comment, error)
	GetCommentByID(id uint) (*models.Comment, error)
	UpdateComment(comment *models.Comment, newComment *models.Comment) (*models.Comment, error)
	DeleteComment(comment *models.Comment) error
	GetCommentByPhotoID(photoID uint) ([]*models.Comment, error)
}

type SocialMediaRepository interface {
	PostSocialMedia(socialmedia *models.SocialMedia) (*models.SocialMedia, error)
	GetAllSocialMedia() ([]*models.SocialMedia, error)
	GetSocialMediaByID(id uint) (*models.SocialMedia, error)
	UpdateSocialMedia(socialmedia *models.SocialMedia, newSocialMedia *models.SocialMedia) (*models.SocialMedia, error)
	DeleteSocialMedia(socialmedia *models.SocialMedia) error
}
