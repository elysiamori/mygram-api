package mygram

import (
	"github.com/elysiamori/mygram-api/api/models"
	"gorm.io/gorm"
)

type CommentRepositoryImpl struct {
	DB *gorm.DB
}

func NewCommentRepository(db *gorm.DB) *CommentRepositoryImpl {
	return &CommentRepositoryImpl{DB: db}
}

func (r *CommentRepositoryImpl) PostComment(comment *models.Comment) (*models.Comment, error) {
	err := r.DB.Create(&comment).Error
	if err != nil {
		return comment, err
	}
	return comment, nil
}

func (r *CommentRepositoryImpl) GetAllComment() ([]models.Comment, error) {
	var usercomments []models.Comment
	err := r.DB.Preload("User").Preload("Photo").Find(&usercomments).Error
	if err != nil {
		return []models.Comment{}, err
	}
	return usercomments, nil
}

func (r *CommentRepositoryImpl) GetCommentByID(id int) (*models.Comment, error) {
	var comment models.Comment
	err := r.DB.Where("id = ?", id).Take(&comment).Error
	if err != nil {
		return nil, err
	}
	return &comment, nil
}

func (r *CommentRepositoryImpl) UpdateComment(comment *models.Comment, newComment *models.Comment) (*models.Comment, error) {
	err := r.DB.Model(&comment).Updates(newComment).Error
	if err != nil {
		return nil, err
	}
	return comment, nil
}

func (r *CommentRepositoryImpl) DeleteComment(comment *models.Comment) error {
	err := r.DB.Delete(&comment).Error
	if err != nil {
		return err
	}
	return nil
}
