package services

import (
	"github.com/elysiamori/mygram-api/api/models"
	"github.com/elysiamori/mygram-api/api/repositories/mygram"
	"github.com/elysiamori/mygram-api/dto/request"
	"github.com/elysiamori/mygram-api/dto/response"
)

type UserServiceImpl struct {
	UserRepo mygram.UserRepositoryImpl
}

func NewUserService(userRepo mygram.UserRepositoryImpl) *UserServiceImpl {
	return &UserServiceImpl{UserRepo: userRepo}
}

func (u *UserServiceImpl) RegisterUser(user *models.User) (*models.User, error) {
	_, err := u.UserRepo.RegisterUser(user)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (u *UserServiceImpl) BeforeSaveUser(user *models.User) error {
	err := u.UserRepo.BeforeSave(user)
	if err != nil {
		return err
	}

	return nil
}

func (u *UserServiceImpl) LoginUser(email, password string) (string, error) {
	user, err := u.UserRepo.LoginUser(email, password)
	if err != nil {
		return user, err
	}

	return user, err
}

func (u *UserServiceImpl) LoginUserWithUsername(username, password string) (string, error) {
	user, err := u.UserRepo.LoginUser(username, password)
	if err != nil {
		return user, err
	}

	return user, err
}

func (u *UserServiceImpl) GetUserByID(id uint) (*models.User, error) {
	user, err := u.UserRepo.GetUserById(id)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (u *UserServiceImpl) GetUserByUsername(username string) (*models.User, error) {
	user, err := u.UserRepo.GetUserByUsername(username)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (u *UserServiceImpl) UpdateUser(userID uint, user *request.UpdateRequest) (*response.UpdateResponse, error) {
	oldUser, err := u.UserRepo.GetUserById(userID)

	if err != nil {
		return nil, err
	}
	newUser := &models.User{
		Email:    user.Email,
		Username: user.Username,
	}

	userUp, err := u.UserRepo.UpdateUser(oldUser, newUser)
	if err != nil {
		return nil, err
	}

	response := &response.UpdateResponse{
		ID:       userUp.ID,
		Email:    userUp.Email,
		Username: userUp.Username,
		Age:      userUp.Age,
		UpdateAt: userUp.UpdateAt.Format("2006-01-02 15:04:05"),
	}

	return response, nil
}

func (u *UserServiceImpl) DeletedUser(userID uint) (*response.DeleteResponse, error) {
	err := u.UserRepo.DeleteUser(userID)

	if err != nil {
		return nil, err
	}

	delResponse := &response.DeleteResponse{
		Message: "Your account has been successfully deleted",
	}

	return delResponse, nil
}
