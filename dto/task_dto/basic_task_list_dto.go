package task_dto

import (
	"awesomeProject/models"
)

type TaskBasicDTO struct {
	ID       uint            `json:"id"`
	Name     string          `json:"name"`
	Priority models.Priority `json:"priority"`
}
