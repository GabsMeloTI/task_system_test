package task_dto

import (
	"awesomeProject/dto/label_dto"
	"awesomeProject/dto/section_dto"
	"awesomeProject/dto/user_dto"
	"awesomeProject/models"
	"time"
)

type TaskListingDTO struct {
	ID                 uint                        `json:"id"`
	Title              string                      `json:"title"`
	Description        string                      `json:"description"`
	ExpectedCompletion time.Time                   `json:"expected_completion"`
	Priority           models.Priority             `json:"priority"`
	CreatedAt          time.Time                   `json:"created_at"`
	Status             string                      `json:"status"`
	Labels             []label_dto.LabelListingDTO `json:"labels"`
	User               user_dto.UserBasicDTO       `json:"user"`
	Section            section_dto.SectionBasicDTO `json:"section"`
}
