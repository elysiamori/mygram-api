package handler

import (
	"github.com/elysiamori/mygram-api/api/models"
	"github.com/elysiamori/mygram-api/api/services"
	"github.com/elysiamori/mygram-api/dto/request"
	"github.com/elysiamori/mygram-api/dto/response"
	"github.com/elysiamori/mygram-api/helpers"
	"github.com/gofiber/fiber/v2"
)

type SocialMediaHandler struct {
	SocialMediaService services.SocialMediaServiceImpl
	UserService        services.UserServiceImpl
}

func NewSocialMediaHandler(socialMediaService services.SocialMediaServiceImpl) *SocialMediaHandler {
	return &SocialMediaHandler{SocialMediaService: socialMediaService}
}

func (h *SocialMediaHandler) PostSocialMedia(c *fiber.Ctx) error {
	request := new(request.PostSocialMediaRequest)

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

	socialMedia := models.SocialMedia{
		Name:           request.Name,
		SocialMediaURL: request.SocialMediaURL,
		UserID:         userID,
	}

	_, err = h.SocialMediaService.PostSocialMedia(&socialMedia)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	response := response.SocialMediaResponse{
		ID:             socialMedia.ID,
		Name:           socialMedia.Name,
		SocialMediaURL: socialMedia.SocialMediaURL,
		UserID:         socialMedia.UserID,
		CreatedAt:      socialMedia.CreatedAt.Format("2006-01-02 15:04:05"),
	}

	return c.Status(fiber.StatusOK).JSON(response)
}

func (h *SocialMediaHandler) GetAllSocialMedia(c *fiber.Ctx) error {
	socialMedia, err := h.SocialMediaService.GetAllSocialMedia()
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	responses := []response.SocialMediaUserResponse{}

	for _, socialMedia := range socialMedia {
		socialmedia, err := h.UserService.GetUserByID(socialMedia.UserID)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"message": err.Error(),
			})
		}

		responses = append(responses, response.SocialMediaUserResponse{
			ID:             socialMedia.ID,
			Name:           socialMedia.Name,
			SocialMediaURL: socialMedia.SocialMediaURL,
			UserID:         socialMedia.UserID,
			CreatedAt:      socialMedia.CreatedAt.Format("2006-01-02 15:04:05"),
			UpdatedAt:      socialMedia.UpdatedAt.Format("2006-01-02 15:04:05"),
			User: response.UserSocialMediaResponse{
				ID:       socialmedia.ID,
				Username: socialmedia.Username,
			},
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"social_medias": responses,
	})
}
