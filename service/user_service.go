package service

import (
	"awesomeProject/configs"
	"awesomeProject/dto/user_dto"
	"awesomeProject/models"
	"awesomeProject/utils"
	"errors"
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/dgrijalva/jwt-go"
	"gorm.io/gorm"
	"mime/multipart"
	"path/filepath"
	"time"
)

var jwtKey = []byte("your_secret_key")

type UserService interface {
	GetUsers() ([]user_dto.UserListingDTO, error)
	GetUserByID(id string) (user_dto.UserListingDTO, error)
	CreateUser(user *models.User, imageFile multipart.File, imageFileName string) error
	UpdateUser(id string, user *models.User, imageFile multipart.File, imageFileName string) error
	DeleteUser(id string) error
	Login(email, password string) (string, error)
	RegisterUser(user *models.User, imageFile multipart.File, imageFileName string) error
	UpdateUserImage(id string, file multipart.File, imageFileName string) error
}

type userService struct {
	db *gorm.DB
}

func NewUserService(db *gorm.DB) UserService {
	return &userService{db: db}
}

func (s *userService) GetUsers() ([]user_dto.UserListingDTO, error) {
	var users []models.User
	var usersDto []user_dto.UserListingDTO

	err := s.db.Find(&users).Error
	if err != nil {
		return nil, err
	}

	for _, user := range users {
		userDto := user_dto.UserListingDTO{
			ID:        user.ID,
			CreatedAt: user.CreatedAt,
			Name:      user.Name,
			Email:     user.Email,
			Photo:     user.Avatar,
		}
		usersDto = append(usersDto, userDto)
	}

	return usersDto, nil
}

func (s *userService) GetUserByID(id string) (user_dto.UserListingDTO, error) {
	var user models.User
	err := s.db.First(&user, id).Error
	if err != nil {
		return user_dto.UserListingDTO{}, err
	}

	userDto := user_dto.UserListingDTO{
		ID:        user.ID,
		CreatedAt: user.CreatedAt,
		Name:      user.Name,
		Email:     user.Email,
		Photo:     user.Avatar,
	}

	return userDto, nil
}

func (s *userService) CreateUser(user *models.User, imageFile multipart.File, imageFileName string) error {
	var existingUser models.User
	err := s.db.Where("email = ?", user.Email).First(&existingUser).Error
	if err == nil {
		return errors.New("email already registered")
	}

	if imageFile != nil {
		imageURL, err := uploadFileToS3(imageFile, imageFileName)
		if err != nil {
			return err
		}
		user.Avatar = imageURL
	}

	hashedPassword, err := utils.HashPassword(user.Password)
	if err != nil {
		return err
	}
	user.Password = hashedPassword

	return s.db.Create(user).Error
}

func (s *userService) UpdateUser(id string, user *models.User, imageFile multipart.File, imageFileName string) error {
	var existingUser models.User
	err := s.db.First(&existingUser, id).Error
	if err != nil {
		return err
	}

	if imageFile != nil {
		imageURL, err := uploadFileToS3(imageFile, imageFileName)
		if err != nil {
			return err
		}
		user.Avatar = imageURL
	}

	return s.db.Model(&existingUser).Updates(user).Error
}

func (s *userService) DeleteUser(id string) error {
	return s.db.Delete(&models.User{}, id).Error
}

func (s *userService) RegisterUser(user *models.User, imageFile multipart.File, imageFileName string) error {
	return s.CreateUser(user, imageFile, imageFileName)
}

func (s *userService) Login(email, password string) (string, error) {
	var user models.User
	err := s.db.Where("email = ?", email).First(&user).Error
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

func uploadFileToS3(file multipart.File, filename string) (string, error) {
	if configs.S3Client == nil {
		return "", fmt.Errorf("S3 client is not initialized")
	}

	bucket := "task-sytem-upload-image" // Substitua pelo nome do seu bucket

	_, err := configs.S3Client.PutObject(&s3.PutObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(filepath.Base(filename)),
		Body:   file,
	})
	if err != nil {
		return "", fmt.Errorf("failed to upload file to S3: %v", err)
	}

	imageURL := fmt.Sprintf("https://%s.s3.amazonaws.com/%s", bucket, filepath.Base(filename))
	return imageURL, nil
}

func (s *userService) UpdateUserImage(id string, imageFile multipart.File, imageFileName string) error {
	var user models.User
	err := s.db.First(&user, id).Error
	if err != nil {
		return err
	}

	if imageFile != nil {
		imageURL, err := uploadFileToS3(imageFile, imageFileName)
		if err != nil {
			return err
		}
		user.Avatar = imageURL
	}

	return s.db.Model(&user).Update("avatar", user.Avatar).Error
}
