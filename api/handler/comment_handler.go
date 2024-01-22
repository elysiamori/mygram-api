package handler

import (
	"strconv"

	"github.com/elysiamori/mygram-api/api/models"
	"github.com/elysiamori/mygram-api/api/services"
	"github.com/elysiamori/mygram-api/dto/request"
	"github.com/elysiamori/mygram-api/dto/response"
	"github.com/elysiamori/mygram-api/helpers"
	"github.com/gofiber/fiber/v2"
)

type CommentHandler struct {
	CommentService services.CommentServiceImpl
}

func NewCommentHandler(commentService services.CommentServiceImpl) *CommentHandler {
	return &CommentHandler{CommentService: commentService}
}

func (h *CommentHandler) PostComment(c *fiber.Ctx) error {
	request := new(request.CommentRequest)

	if err := c.BodyParser(request); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	userID, err := helpers.ExtractTokenID(c)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	comment := models.Comment{
		Message: request.Message,
		PhotoID: request.PhotoID,
		UserID:  userID,
	}

	_, err = h.CommentService.PostComment(&comment)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	response := response.CommentResponse{
		ID:        comment.ID,
		Message:   comment.Message,
		PhotoID:   comment.PhotoID,
		UserID:    comment.UserID,
		CreatedAt: comment.CreatedAt.Format("2006-01-02 15:04:05"),
	}
	return c.Status(fiber.StatusOK).JSON(response)
}

func (h *CommentHandler) GetAllComment(c *fiber.Ctx) error {
	comments, err := h.CommentService.GetAllComment()
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	responses := []response.GetCommentsResponse{}
	for _, comment := range comments {
		responses = append(responses, response.GetCommentsResponse{
			ID:        comment.ID,
			Message:   comment.Message,
			PhotoID:   comment.PhotoID,
			UserID:    comment.UserID,
			CreatedAt: comment.CreatedAt.Format("2006-01-02 15:04:05"),
			User: response.UserComment{
				ID:       comment.User.ID,
				Email:    comment.User.Email,
				Username: comment.User.Username,
			},
			Photo: response.PhotoComment{
				ID:       comment.Photo.ID,
				Title:    comment.Photo.Title,
				Caption:  comment.Photo.Caption,
				PhotoURL: comment.Photo.PhotoURL,
				UserID:   comment.Photo.UserID,
			},
		})
	}
	return c.Status(fiber.StatusOK).JSON(responses)
}

func (h *CommentHandler) GetCommentByID(c *fiber.Ctx) error {
	id := c.Params("id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	comment, err := h.CommentService.GetCommentByID(idInt)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	response := response.CommentResponse{
		ID:        comment.ID,
		Message:   comment.Message,
		PhotoID:   comment.PhotoID,
		UserID:    comment.UserID,
		CreatedAt: comment.CreatedAt.Format("2006-01-02 15:04:05"),
	}
	return c.Status(fiber.StatusOK).JSON(response)
}

func (h *CommentHandler) UpdateComment(c *fiber.Ctx) error {
	userID, err := helpers.ExtractTokenID(c)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	id := c.Params("id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	request := new(request.CommentUpdateRequest)

	if err := c.BodyParser(request); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	comment, err := h.CommentService.GetCommentByID(idInt)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	if userID != comment.UserID {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "You are not authorized to update this comment",
		})
	}

	commentUpdated, err := h.CommentService.UpdateComment(uint(idInt), request)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(commentUpdated)
}

func (h *CommentHandler) DeleteComment(c *fiber.Ctx) error {
	userID, err := helpers.ExtractTokenID(c)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	id := c.Params("id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	comment, err := h.CommentService.GetCommentByID(idInt)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	if userID != comment.UserID {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "You are not authorized to delete this comment",
		})
	}

	err = h.CommentService.DeleteComment(uint(idInt), comment)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Your comment has been successfully deleted",
	})
}
