package service

import (
	"awesomeProject/db"
	"awesomeProject/dto/comment_dto"
	"awesomeProject/dto/task_dto"
	"awesomeProject/dto/user_dto"
	"awesomeProject/models"
	"bytes"
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"io"
	"mime/multipart"
	"time"
)

type CommentService struct{}

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
				ID:   comment.User.ID,
				Name: comment.User.Name,
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
			ID:   comment.User.ID,
			Name: comment.User.Name,
		},
		Task: task_dto.TaskBasicDTO{
			ID:    comment.Task.ID,
			Title: comment.Task.Title,
		},
	}

	return commentDTO, nil
}

func (s *CommentService) CreateComment(content string, imageURL string) error {
	comment := models.Comment{
		Content:     content,
		Image:       imageURL,
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

func (s *CommentService) UploadToS3(file multipart.File, filename string) (string, error) {
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String("us-west-2"), // ajuste para sua região
	})
	if err != nil {
		return "", fmt.Errorf("failed to create AWS session: %w", err)
	}

	uploader := s3manager.NewUploader(sess)

	fileBytes, err := io.ReadAll(file)
	if err != nil {
		return "", fmt.Errorf("failed to read file: %w", err)
	}

	result, err := uploader.Upload(&s3manager.UploadInput{
		Bucket: aws.String("task-sytem-upload-image--use1-az4--x-s3"),
		Key:    aws.String(fmt.Sprintf("images/%d_%s", time.Now().UnixNano(), filename)),
		Body:   bytes.NewReader(fileBytes),
	})
	if err != nil {
		return "", fmt.Errorf("failed to upload file to S3: %w", err)
	}

	return result.Location, nil
}
