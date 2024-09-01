package subtask_dto

import (
	"awesomeProject/dto/task_dto"
	"time"
)

type SubtaskListingDTO struct {
	ID          uint                  `json:"id"`
	Title       string                `json:"title"`
	Description string                `json:"description"`
	CreatedAt   time.Time             `json:"created_at"`
	Status      string                `json:"status"`
	Task        task_dto.TaskBasicDTO `json:"task"`
}
