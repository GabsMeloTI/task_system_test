package task_dto

import (
	"awesomeProject/models"
)

type TaskBasicDTO struct {
	ID       uint            `json:"id"`
	Title    string          `json:"title"`
	Priority models.Priority `json:"priority"`
}
