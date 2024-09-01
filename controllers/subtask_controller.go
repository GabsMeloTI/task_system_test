package controllers

import (
	"awesomeProject/db"
	"awesomeProject/dto/subtask_dto"
	"awesomeProject/dto/task_dto"
	"awesomeProject/models"
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
)

func GetSubtasks(w http.ResponseWriter, r *http.Request) {
	var subtasks []models.Subtask

	db.DB.Preload("Task").Find(&subtasks)

	var subtasksDTO []subtask_dto.SubtaskListingDTO
	for _, subtask := range subtasks {
		subtasksDTO = append(subtasksDTO, subtask_dto.SubtaskListingDTO{
			ID:          subtask.ID,
			Title:       subtask.Title,
			Description: subtask.Description,
			CreatedAt:   subtask.CreatedAt,
			Status:      subtask.Status,
			Task: task_dto.TaskBasicDTO{
				ID:    subtask.Task.ID,
				Title: subtask.Task.Title,
			},
		})
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(subtasksDTO)
}

func GetSubtaskByID(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var subtask models.Subtask

	err := db.DB.Preload("Task").First(&subtask, params["id"]).Error
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	subtaskDTO := subtask_dto.SubtaskListingDTO{
		ID:          subtask.ID,
		Title:       subtask.Title,
		Description: subtask.Description,
		CreatedAt:   subtask.CreatedAt,
		Status:      subtask.Status,
		Task: task_dto.TaskBasicDTO{
			ID:    subtask.Task.ID,
			Title: subtask.Task.Title,
		},
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(subtaskDTO)
}

func CreateSubtask(w http.ResponseWriter, r *http.Request) {
	var subtask models.Subtask

	err := json.NewDecoder(r.Body).Decode(&subtask)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = db.DB.Create(&subtask).Error
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func UpdateSubtask(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var subtask models.Subtask

	err := db.DB.First(&subtask, params["id"]).Error
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	err = json.NewDecoder(r.Body).Decode(&subtask)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	db.DB.Save(&subtask)
	w.WriteHeader(http.StatusNoContent)
}

func DeleteSubtask(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	err := db.DB.Delete(&models.Subtask{}, params["id"]).Error
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
