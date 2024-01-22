package handler

import (
	"github.com/elysiamori/mygram-api/api/models"
	"github.com/elysiamori/mygram-api/api/repositories/mygram"
	"github.com/elysiamori/mygram-api/api/services"
	"github.com/elysiamori/mygram-api/dto/request"
	"github.com/elysiamori/mygram-api/dto/response"
	"github.com/elysiamori/mygram-api/helpers"
	"github.com/gofiber/fiber/v2"
)

type UserHandler struct {
	userService services.UserServiceImpl
}

func NewUserHandler(userService services.UserServiceImpl, userRepo mygram.UserRepositoryImpl) *UserHandler {
	return &UserHandler{userService: userService}
}

func (u *UserHandler) ProfileUser(c *fiber.Ctx) error {
	user_id, err := helpers.ExtractTokenID(c)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	user, err := u.userService.GetUserByID(user_id)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(user)
}

func (u *UserHandler) RegisterUser(c *fiber.Ctx) error {
	// get request body
	request := new(request.RegisterRequest)

	if err := c.BodyParser(request); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	ru := models.User{
		Age:      request.Age,
		Email:    request.Email,
		Password: request.Password,
		Username: request.Username,
	}

	if err := u.userService.BeforeSaveUser(&ru); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	_, err := u.userService.RegisterUser(&ru)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	response := response.RegisterResponse{
		Age:      ru.Age,
		Email:    ru.Email,
		ID:       ru.ID,
		Username: ru.Username,
	}

	return c.Status(fiber.StatusCreated).JSON(response)
}

func (u *UserHandler) LoginUser(c *fiber.Ctx) error {
	request := new(request.LoginRequest)

	if err := c.BodyParser(request); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	token, err := u.userService.LoginUser(request.Email, request.Password)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "account not found",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"token": token,
	})
}

func (u *UserHandler) LoginUserWithUsername(c *fiber.Ctx) error {
	request := new(request.LoginRequestUsername)

	if err := c.BodyParser(request); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	token, err := u.userService.LoginUser(request.Username, request.Password)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "account not found",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"token": token,
	})
}

func (u *UserHandler) UpdateUser(c *fiber.Ctx) error {
	user_id, err := helpers.ExtractTokenID(c)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	user := new(request.UpdateRequest)

	if err := c.BodyParser(user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	userUp, err := u.userService.UpdateUser(user_id, user)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(userUp)
}

func (u *UserHandler) DeleteUser(c *fiber.Ctx) error {
	user_id, err := helpers.ExtractTokenID(c)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	delResponse, err := u.userService.DeletedUser(user_id)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(delResponse)
}
