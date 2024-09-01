package controllers

import (
	"awesomeProject/db"
	"awesomeProject/dto/label_dto"
	"awesomeProject/dto/section_dto"
	"awesomeProject/dto/task_dto"
	"awesomeProject/dto/user_dto"
	"awesomeProject/models"
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"strconv"
)

func GetTasks(w http.ResponseWriter, r *http.Request) {
	var tasks []models.Task
	var tasksDTO []task_dto.TaskListingDTO

	err := db.DB.Preload("User").Preload("Section").Preload("Labels").Find(&tasks).Error
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	for _, task := range tasks {
		log.Printf("Task: %+v\n", task) // Log for debugging

		var labelsDTO []label_dto.LabelListingDTO
		for _, label := range task.Labels {
			log.Printf("Label: %+v\n", label) // Log for debugging
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

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(tasksDTO)
}

func GetTaskByID(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var task models.Task

	err := db.DB.Preload("User").Preload("Section").First(&task, params["id"]).Error
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
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

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(taskDTO)
}

func CreateTask(w http.ResponseWriter, r *http.Request) {
	var task models.Task

	err := json.NewDecoder(r.Body).Decode(&task)
	if err != nil {
		http.Error(w, "Invalid data", http.StatusBadRequest)
		return
	}

	userID := task.UserID
	sectionID := task.SectionID

	var user models.User
	var section models.Section

	err = db.DB.First(&user, userID).Error
	if err != nil {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	err = db.DB.First(&section, sectionID).Error
	if err != nil {
		http.Error(w, "Section not found", http.StatusNotFound)
		return
	}

	task.User = user
	task.Section = section

	err = db.DB.Create(&task).Error
	if err != nil {
		http.Error(w, "Error creating task", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func UpdateTask(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var task models.Task

	err := db.DB.First(&task, params["id"]).Error
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	err = json.NewDecoder(r.Body).Decode(&task)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	db.DB.Save(&task)

	w.WriteHeader(http.StatusNoContent)
}

func DeleteTask(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	err := db.DB.Delete(&models.Task{}, params["id"]).Error
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func AssignLabelsToTask(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	taskID, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, "Invalid task ID", http.StatusBadRequest)
		return
	}

	var task models.Task
	err = db.DB.Preload("Labels").First(&task, taskID).Error
	if err != nil {
		http.Error(w, "Task not found", http.StatusNotFound)
		return
	}

	var labels []models.Label
	err = json.NewDecoder(r.Body).Decode(&labels)
	if err != nil {
		http.Error(w, "Invalid data", http.StatusBadRequest)
		return
	}

	task.Labels = labels

	err = db.DB.Save(&task).Error
	if err != nil {
		http.Error(w, "Error assigning labels", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
