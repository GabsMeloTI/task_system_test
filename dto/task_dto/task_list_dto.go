package task_dto

import (
	"awesomeProject/dto/label_dto"
	"awesomeProject/dto/section_dto"
	"awesomeProject/dto/user_dto"
	"awesomeProject/models"
	"time"
)

type TaskListingDTO struct {
	ID                  uint                        `json:"id"`
	Name                string                      `json:"name"`
	Description         string                      `json:"description"`
	EstimatedCompletion time.Time                   `json:"estimated_completion_date"`
	Priority            models.Priority             `json:"priority"`
	CreatedAt           time.Time                   `json:"created_at"`
	Status              string                      `json:"status"`
	Labels              []label_dto.LabelListingDTO `json:"labels"`
	User                user_dto.UserBasicDTO       `json:"user"`
	Section             section_dto.SectionBasicDTO `json:"section"`
}
