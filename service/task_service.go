package service

import (
	"awesomeProject/db"
	"awesomeProject/dto/label_dto"
	"awesomeProject/dto/section_dto"
	"awesomeProject/dto/task_dto"
	"awesomeProject/dto/user_dto"
	"awesomeProject/models"
	"errors"
	_ "gorm.io/gorm"
)

type TaskService struct{}

func (s *TaskService) GetAllTasks() ([]task_dto.TaskListingDTO, error) {
	var tasks []models.Task
	var tasksDTO []task_dto.TaskListingDTO

	err := db.DB.Preload("User").Preload("Section").Preload("Labels").Find(&tasks).Error
	if err != nil {
		return nil, err
	}

	for _, task := range tasks {
		var labelsDTO []label_dto.LabelListingDTO
		for _, label := range task.Labels {
			labelsDTO = append(labelsDTO, label_dto.LabelListingDTO{
				ID:    label.ID,
				Name:  label.Name,
				Color: label.Color,
			})
		}

		tasksDTO = append(tasksDTO, task_dto.TaskListingDTO{
			ID:                 task.ID,
			Title:              task.Title,
			Description:        task.Description,
			ExpectedCompletion: task.ExpectedCompletion,
			Priority:           task.Priority,
			CreatedAt:          task.CreatedAt,
			Status:             task.Status,
			Labels:             labelsDTO,
			User: user_dto.UserBasicDTO{
				ID:   task.User.ID,
				Name: task.User.Name,
			},
			Section: section_dto.SectionBasicDTO{
				ID:    task.Section.ID,
				Title: task.Section.Title,
			},
		})
	}

	return tasksDTO, nil
}

func (s *TaskService) GetTaskByID(id uint) (task_dto.TaskListingDTO, error) {
	var task models.Task

	err := db.DB.Preload("User").Preload("Section").First(&task, id).Error
	if err != nil {
		return task_dto.TaskListingDTO{}, err
	}

	var labelsDTO []label_dto.LabelListingDTO
	for _, label := range task.Labels {
		labelsDTO = append(labelsDTO, label_dto.LabelListingDTO{
			ID:    label.ID,
			Name:  label.Name,
			Color: label.Color,
		})
	}

	taskDTO := task_dto.TaskListingDTO{
		ID:                 task.ID,
		Title:              task.Title,
		Description:        task.Description,
		ExpectedCompletion: task.ExpectedCompletion,
		Priority:           task.Priority,
		CreatedAt:          task.CreatedAt,
		Status:             task.Status,
		Labels:             labelsDTO,
		User: user_dto.UserBasicDTO{
			ID:   task.User.ID,
			Name: task.User.Name,
		},
		Section: section_dto.SectionBasicDTO{
			ID:    task.Section.ID,
			Title: task.Section.Title,
		},
	}

	return taskDTO, nil
}

func (s *TaskService) CreateTask(task models.Task) error {
	var user models.User
	var section models.Section

	err := db.DB.First(&user, task.UserID).Error
	if err != nil {
		return errors.New("user not found")
	}

	err = db.DB.First(&section, task.SectionID).Error
	if err != nil {
		return errors.New("section not found")
	}

	task.User = user
	task.Section = section

	return db.DB.Create(&task).Error
}

func (s *TaskService) UpdateTask(id uint, task models.Task) error {
	var existingTask models.Task

	err := db.DB.First(&existingTask, id).Error
	if err != nil {
		return err
	}

	task.ID = existingTask.ID
	return db.DB.Save(&task).Error
}

func (s *TaskService) DeleteTask(id uint) error {
	return db.DB.Delete(&models.Task{}, id).Error
}

func (s *TaskService) AssignLabelsToTask(taskID uint, labels []models.Label) error {
	var task models.Task
	err := db.DB.Preload("Labels").First(&task, taskID).Error
	if err != nil {
		return err
	}

	task.Labels = labels
	return db.DB.Save(&task).Error
}
