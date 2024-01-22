package mygram

import (
	"errors"
	"html"
	"strings"

	"github.com/elysiamori/mygram-api/api/models"
	"github.com/elysiamori/mygram-api/helpers"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type UserRepositoryImpl struct {
	DB *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepositoryImpl {
	return &UserRepositoryImpl{DB: db}
}

// register user
func (r *UserRepositoryImpl) RegisterUser(user *models.User) (*models.User, error) {
	err := r.DB.Create(&user).Error

	if err != nil {
		return nil, err
	}

	return user, nil
}

// save user
func (r *UserRepositoryImpl) BeforeSave(user *models.User) error {
	// turn password into hash
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user.Password = string(hashedPassword)

	// remove spaces in username
	user.Username = html.EscapeString(strings.TrimSpace(user.Username))
	return nil
}

// login user
func (r *UserRepositoryImpl) LoginUser(email, password string) (string, error) {
	var err error

	u := models.User{}

	err = r.DB.Model(models.User{}).Where("email = ?", email).Take(&u).Error

	if err != nil {
		return "", err
	}

	err = models.VerifyPassword(password, u.Password)

	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword {
		return "", err
	}

	// generate token json web token (jwt)
	token, err := helpers.GenerateToken(u.ID)

	if err != nil {
		return "", err
	}

	return token, nil
}

// login with username
// login user
func (r *UserRepositoryImpl) LoginUserWithUsername(username, password string) (string, error) {
	var err error

	u := models.User{}

	err = r.DB.Model(models.User{}).Where("username = ?", username).Take(&u).Error

	if err != nil {
		return "", err
	}

	err = models.VerifyPassword(password, u.Password)

	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword {
		return "", err
	}

	token, err := helpers.GenerateToken(u.ID)

	if err != nil {
		return "", err
	}

	return token, nil
}

func (r *UserRepositoryImpl) GetUserById(uid uint) (*models.User, error) {
	var u models.User

	err := r.DB.First(&u, uid).Error

	if err != nil {
		return nil, errors.New("User not found")
	}

	return &u, nil
}

func (r *UserRepositoryImpl) PrepareGive() {
	var user models.User
	user.Password = ""
}

func (r *UserRepositoryImpl) GetUserByEmail(email string) (models.User, error) {
	var u models.User

	err := r.DB.First(models.User{}).Where("email = ?", email).Error

	if err != nil {
		return u, errors.New("User not found")
	}

	return u, nil
}

func (r *UserRepositoryImpl) GetUserByUsername(username string) (models.User, error) {
	var u models.User

	err := r.DB.First(models.User{}).Where("username = ?", username).Error

	if err != nil {
		return u, errors.New("User not found")
	}

	return u, nil
}

func (r *UserRepositoryImpl) UpdateUser(user *models.User, newUser *models.User) (*models.User, error) {

	err := r.DB.Model(user).Updates(newUser).Error

	if err != nil {
		return nil, err
	}

	return user, nil
}

func (r *UserRepositoryImpl) DeleteUser(id uint) error {
	delUser := models.User{}
	err := r.DB.Delete(&delUser, id).Error

	if err != nil {
		return err
	}

	return nil
}
