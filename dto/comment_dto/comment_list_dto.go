package comment_dto

import (
	"awesomeProject/dto/task_dto"
	"awesomeProject/dto/user_dto"
	"time"
)

type CommentListingDTO struct {
	ID          uint                  `json:"id"`
	Content     string                `json:"content"`
	PublishedAt time.Time             `json:"published_at"`
	ImageURL    string                `json:"image_url"`
	User        user_dto.UserBasicDTO `json:"user"`
	Task        task_dto.TaskBasicDTO `json:"task"`
}
