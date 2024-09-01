package subtask_dto

import (
	"time"
)

type SubtaskListingDTO struct {
	ID          uint         `json:"id"`
	Name        string       `json:"name"`
	Description string       `json:"description"`
	CreatedAt   time.Time    `json:"created_at"`
	Status      string       `json:"status"`
	Task        TaskBasicDTO `json:"task"`
}
