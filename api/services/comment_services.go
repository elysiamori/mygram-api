package services

import (
	"github.com/elysiamori/mygram-api/api/models"
	"github.com/elysiamori/mygram-api/api/repositories/mygram"
	"github.com/elysiamori/mygram-api/dto/request"
	"github.com/elysiamori/mygram-api/dto/response"
)

type CommentServiceImpl struct {
	CommentRepo mygram.CommentRepositoryImpl
}

func NewCommentService(commentRepo mygram.CommentRepositoryImpl) *CommentServiceImpl {
	return &CommentServiceImpl{CommentRepo: commentRepo}
}

func (s *CommentServiceImpl) PostComment(comment *models.Comment) (*models.Comment, error) {
	comment, err := s.CommentRepo.PostComment(comment)
	if err != nil {
		return comment, err
	}
	return comment, nil
}

func (s *CommentServiceImpl) GetAllComment() ([]models.Comment, error) {
	usercomments, err := s.CommentRepo.GetAllComment()
	if err != nil {
		return nil, err
	}
	return usercomments, nil
}

func (s *CommentServiceImpl) GetCommentByID(id int) (*models.Comment, error) {
	usercomments, err := s.CommentRepo.GetCommentByID(id)
	if err != nil {
		return nil, err
	}

	return usercomments, nil
}

func (s *CommentServiceImpl) UpdateComment(commentID uint, comment *request.CommentUpdateRequest) (*response.CommentUpdate, error) {
	oldComment, err := s.CommentRepo.GetCommentByID(int(commentID))
	if err != nil {
		return nil, err
	}

	newComment := models.Comment{
		Message: comment.Message,
	}

	commentUpdated, err := s.CommentRepo.UpdateComment(oldComment, &newComment)
	if err != nil {
		return nil, err
	}

	response := &response.CommentUpdate{
		ID:        commentUpdated.ID,
		Message:   commentUpdated.Message,
		PhotoID:   commentUpdated.PhotoID,
		UserID:    commentUpdated.UserID,
		UpdatedAt: commentUpdated.UpdateAt.Format("2006-01-02 15:04:05"),
	}

	return response, nil
}

func (s *CommentServiceImpl) DeleteComment(commentID uint, comment *models.Comment) error {
	comment_id, err := s.GetCommentByID(int(commentID))
	if err != nil {
		return err
	}

	err = s.CommentRepo.DeleteComment(comment_id)
	if err != nil {
		return err
	}

	return nil
}
