package service

import (
	"awesomeProject/db"
	"awesomeProject/dto/subtask_dto"
	"awesomeProject/dto/task_dto"
	"awesomeProject/models"
)

type SubtaskService struct{}

func (s *SubtaskService) GetAllSubtasks() ([]subtask_dto.SubtaskListingDTO, error) {
	var subtasks []models.Subtask
	var subtasksDTO []subtask_dto.SubtaskListingDTO

	err := db.DB.Preload("Task").Find(&subtasks).Error
	if err != nil {
		return nil, err
	}

	for _, subtask := range subtasks {
		subtasksDTO = append(subtasksDTO, subtask_dto.SubtaskListingDTO{
			ID:          subtask.ID,
			Title:       subtask.Title,
			Description: subtask.Description,
			CreatedAt:   subtask.CreatedAt,
			Status:      subtask.Status,
			Task: task_dto.TaskBasicDTO{
				ID:       subtask.Task.ID,
				Title:    subtask.Task.Title,
				Priority: subtask.Task.Priority,
			},
		})
	}

	return subtasksDTO, nil
}

func (s *SubtaskService) GetSubtaskByID(id uint) (subtask_dto.SubtaskListingDTO, error) {
	var subtask models.Subtask

	err := db.DB.Preload("Task").First(&subtask, id).Error
	if err != nil {
		return subtask_dto.SubtaskListingDTO{}, err
	}

	subtaskDTO := subtask_dto.SubtaskListingDTO{
		ID:          subtask.ID,
		Title:       subtask.Title,
		Description: subtask.Description,
		CreatedAt:   subtask.CreatedAt,
		Status:      subtask.Status,
		Task: task_dto.TaskBasicDTO{
			ID:       subtask.Task.ID,
			Title:    subtask.Task.Title,
			Priority: subtask.Task.Priority,
		},
	}

	return subtaskDTO, nil
}

func (s *SubtaskService) CreateSubtask(subtask models.Subtask) error {
	return db.DB.Create(&subtask).Error
}

func (s *SubtaskService) UpdateSubtask(id uint, subtask models.Subtask) error {
	var existingSubtask models.Subtask

	err := db.DB.First(&existingSubtask, id).Error
	if err != nil {
		return err
	}

	subtask.ID = existingSubtask.ID
	return db.DB.Save(&subtask).Error
}

func (s *SubtaskService) DeleteSubtask(id uint) error {
	return db.DB.Delete(&models.Subtask{}, id).Error
}
