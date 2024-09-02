package service

import (
	"awesomeProject/db"
	"awesomeProject/models"
	"awesomeProject/utils"
	"errors"
	"github.com/dgrijalva/jwt-go"
	"time"
)

var jwtKey = []byte("your_secret_key")

type UserService interface {
	GetUsers() ([]models.User, error)
	GetUserByID(id string) (models.User, error)
	CreateUser(user *models.User) error
	UpdateUser(id string, user *models.User) error
	DeleteUser(id string) error
	Login(email, password string) (string, error)
	RegisterUser(user *models.User) error
}

type userService struct{}

func NewUserService() UserService {
	return &userService{}
}

func (s *userService) GetUsers() ([]models.User, error) {
	var users []models.User
	err := db.DB.Preload("Projects").Find(&users).Error
	return users, err
}

func (s *userService) GetUserByID(id string) (models.User, error) {
	var user models.User
	err := db.DB.Preload("Projects").First(&user, id).Error
	return user, err
}

func (s *userService) CreateUser(user *models.User) error {
	var existingUser models.User
	err := db.DB.Where("email = ?", user.Email).First(&existingUser).Error
	if err == nil {
		return errors.New("email already registered")
	}

	hashedPassword, err := utils.HashPassword(user.Password)
	if err != nil {
		return err
	}
	user.Password = hashedPassword

	return db.DB.Create(user).Error
}

func (s *userService) UpdateUser(id string, user *models.User) error {
	var existingUser models.User
	err := db.DB.First(&existingUser, id).Error
	if err != nil {
		return err
	}

	return db.DB.Model(&existingUser).Updates(user).Error
}

func (s *userService) DeleteUser(id string) error {
	return db.DB.Delete(&models.User{}, id).Error
}

func (s *userService) RegisterUser(user *models.User) error {
	return s.CreateUser(user)
}

func (s *userService) Login(email, password string) (string, error) {
	var user models.User
	err := db.DB.Where("email = ?", email).First(&user).Error
	if err != nil || !utils.ComparePasswords(user.Password, password) {
		return "", errors.New("invalid credentials")
	}

	expirationTime := time.Now().Add(24 * time.Hour)
	claims := &utils.JWTClaims{
		Email: email,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
