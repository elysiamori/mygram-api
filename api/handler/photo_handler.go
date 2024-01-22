package handler

import (
	"encoding/json"
	"io"
	"net/http"
	"strconv"

	"github.com/elysiamori/mygram-api/api/models"
	"github.com/elysiamori/mygram-api/api/services"
	"github.com/elysiamori/mygram-api/dto/request"
	"github.com/elysiamori/mygram-api/dto/response"
	"github.com/elysiamori/mygram-api/helpers"
	"github.com/gofiber/fiber/v2"
)

type PhotoHandler struct {
	PhotoService services.PhotoServiceImpl
	UserService  services.UserServiceImpl
}

func NewPhotoHandler(photoService services.PhotoServiceImpl, userService services.UserServiceImpl) *PhotoHandler {
	return &PhotoHandler{PhotoService: photoService, UserService: userService}
}

func (h *PhotoHandler) UploadPhoto(c *fiber.Ctx) error {
	request := new(request.PhotoRequest)

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

	photo := models.Photo{
		Title:    request.Title,
		Caption:  request.Caption,
		PhotoURL: request.PhotoURL,
		UserID:   userID,
	}

	_, err = h.PhotoService.UploadPhoto(&photo)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	response := response.PhotoResponse{
		ID:        photo.ID,
		Title:     photo.Title,
		Caption:   photo.Caption,
		PhotoURL:  photo.PhotoURL,
		UserID:    photo.UserID,
		CreatedAt: photo.CreatedAt.Format("2006-01-02 15:04:05"),
	}

	return c.Status(fiber.StatusOK).JSON(response)

}

func (h *PhotoHandler) GetAllPhoto(c *fiber.Ctx) error {
	photos, err := h.PhotoService.GetAllPhoto()
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	responses := []response.PhotoUserResponse{}

	for _, photo := range photos {
		user, err := h.UserService.GetUserByID(photo.UserID)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"message": err.Error(),
			})
		}
		responses = append(responses, response.PhotoUserResponse{
			ID:        photo.ID,
			Title:     photo.Title,
			Caption:   photo.Caption,
			PhotoURL:  photo.PhotoURL,
			UserID:    photo.UserID,
			CreatedAt: photo.CreatedAt,
			UpdatedAt: photo.UpdatedAt,
			User: response.UserPhotoResponse{
				Email:    user.Email,
				Username: user.Username,
			},
		})
	}

	return c.Status(fiber.StatusOK).JSON(responses)
}
func (h *PhotoHandler) GetPhotoByID(c *fiber.Ctx) error {
	id := c.Params("id")

	idInt, err := strconv.Atoi(id)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	photo, err := h.PhotoService.GetPhotoByID(idInt)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	response := response.PhotoResponse{
		ID:        photo.ID,
		Title:     photo.Title,
		Caption:   photo.Caption,
		PhotoURL:  photo.PhotoURL,
		UserID:    photo.UserID,
		CreatedAt: photo.CreatedAt.Format("2006-01-02 15:04:05"),
	}

	return c.Status(fiber.StatusOK).JSON(response)
}

func (h *PhotoHandler) UpdatePhoto(c *fiber.Ctx) error {

	// get user id from token
	userID, err := helpers.ExtractTokenID(c)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	// get photo id from params
	id := c.Params("id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	request := new(request.UpdatePhotoRequest)

	if err := c.BodyParser(request); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	// get photo by id
	photo, err := h.PhotoService.GetPhotoByID(idInt)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	// check if user id from token is the same with user id from photo
	if userID != photo.UserID {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "You are not authorized to update this photo",
		})
	}

	photoUpdated, err := h.PhotoService.UpdatePhoto(uint(idInt), request)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(photoUpdated)
}

func (h *PhotoHandler) DeletePhoto(c *fiber.Ctx) error {

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

	photo, err := h.PhotoService.GetPhotoByID(idInt)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	if userID != photo.UserID {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "You are not authorized to delete this photo",
		})
	}

	err = h.PhotoService.DeletePhoto(uint(idInt))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Your photo has been successfully deleted",
	})
}

/*----------------------------Fetch API----------------------------------*/

type Content struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
	Desc  string `json:"desc"`
	Date  string `json:"date"`
}

// get all data
type Contents struct {
	Contents []Content `json:"contents"`
}

func getApi(apiURL string) (Contents, error) {
	resp, err := http.Get(apiURL)
	if err != nil {
		return Contents{}, err
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return Contents{}, err
	}

	var c Contents
	err = json.Unmarshal(body, &c)
	if err != nil {
		return Contents{}, err
	}

	return c, nil
}

func (h *PhotoHandler) GetAPI(c *fiber.Ctx) error {
	apiURL := "http://localhost:3000/api/content"

	response, err := getApi(apiURL)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(response)
}

/*------------------------------------------------------------------------*/
