package project_dto

import "time"

type ProjectListingDTO struct {
	ID          uint                  `json:"id"`
	Name        string                `json:"name"`
	Description string                `json:"description"`
	Status      string                `json:"status"`
	CreatedAt   time.Time             `json:"created_at"`
	User        user_dto.UserBasicDTO `json:"user"`
}
