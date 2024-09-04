package service

import (
	"awesomeProject/configs"
	"awesomeProject/db"
	"awesomeProject/dto/comment_dto"
	"awesomeProject/dto/task_dto"
	"awesomeProject/dto/user_dto"
	"awesomeProject/models"
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3"
	"gorm.io/gorm"
	"mime/multipart"
	"path/filepath"
	"time"
)

type CommentService struct {
	DB *gorm.DB
}

func NewCommentService(db *gorm.DB) *CommentService {
	return &CommentService{DB: db}
}

func (s *CommentService) GetAllComments() ([]comment_dto.CommentListingDTO, error) {
	var comments []models.Comment
	var commentsDTO []comment_dto.CommentListingDTO

	if err := db.DB.Preload("User").Preload("Task").Find(&comments).Error; err != nil {
		return nil, err
	}

	for _, comment := range comments {
		commentDTO := comment_dto.CommentListingDTO{
			ID:          comment.ID,
			Content:     comment.Content,
			PublishedAt: comment.PublishedAt,
			ImageURL:    comment.Image,
			User: user_dto.UserBasicDTO{
				ID:    comment.User.ID,
				Name:  comment.User.Name,
				Email: comment.User.Email,
			},
			Task: task_dto.TaskBasicDTO{
				ID:       comment.Task.ID,
				Title:    comment.Task.Title,
				Priority: comment.Task.Priority,
			},
		}
		commentsDTO = append(commentsDTO, commentDTO)
	}

	return commentsDTO, nil
}

func (s *CommentService) GetCommentByID(id uint) (comment_dto.CommentListingDTO, error) {
	var comment models.Comment

	if err := db.DB.Preload("User").Preload("Task").First(&comment, id).Error; err != nil {
		return comment_dto.CommentListingDTO{}, err
	}

	commentDTO := comment_dto.CommentListingDTO{
		ID:          comment.ID,
		Content:     comment.Content,
		PublishedAt: comment.PublishedAt,
		ImageURL:    comment.Image,
		User: user_dto.UserBasicDTO{
			ID:    comment.User.ID,
			Name:  comment.User.Name,
			Email: comment.User.Email,
		},
		Task: task_dto.TaskBasicDTO{
			ID:    comment.Task.ID,
			Title: comment.Task.Title,
		},
	}

	return commentDTO, nil
}

func (s *CommentService) CreateComment(content string, imageURL string, userID int, taskID int) error {
	comment := models.Comment{
		Content:     content,
		Image:       imageURL,
		UserID:      userID,
		TaskID:      taskID,
		PublishedAt: time.Now(),
	}

	return db.DB.Create(&comment).Error
}

func (s *CommentService) UpdateComment(id uint, content string, imageURL string) error {
	var comment models.Comment

	if err := db.DB.First(&comment, id).Error; err != nil {
		return err
	}

	comment.Content = content
	comment.Image = imageURL

	return db.DB.Save(&comment).Error
}

func (s *CommentService) DeleteComment(id uint) error {
	return db.DB.Delete(&models.Comment{}, id).Error
}

func UploadFileToS3(file multipart.File, filename string) (string, error) {
	if configs.S3Client == nil {
		return "", fmt.Errorf("S3 client is not initialized")
	}

	bucket := "task-sytem-upload-image"

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
