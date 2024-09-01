package project_dto

import (
	"awesomeProject/dto/user_dto"
	"time"
)

type ProjectListingDTO struct {
	ID          uint                  `json:"id"`
	Title       string                `json:"title"`
	Description string                `json:"description"`
	Status      string                `json:"status"`
	CreatedAt   time.Time             `json:"created_at"`
	User        user_dto.UserBasicDTO `json:"user"`
}
