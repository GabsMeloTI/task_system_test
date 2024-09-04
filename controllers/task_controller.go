package controllers

import (
	"awesomeProject/models"
	"awesomeProject/service"
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

type TaskController struct {
	Service *service.TaskService
}

// GetTasks retrieves all tasks
// @Summary Get all tasks
// @Description Fetches all tasks available in the system
// @Tags tasks
// @Produce json
// @Success 200 {array} task_dto.TaskListingDTO
// @Success 200 {array} section_dto.SectionBasicDTO
// @Success 200 {array} user_dto.UserBasicDTO
// @Success 200 {array} label_dto.LabelListingDTO
// @Failure 500 {string} string "Internal Server Error"
// @Router /task [get]
func (c *TaskController) GetTasks(w http.ResponseWriter, r *http.Request) {
	tasksDTO, err := c.Service.GetAllTasks()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(tasksDTO)
}

// GetTaskByID retrieves a task by its ID
// @Summary Get a task by ID
// @Description Fetches the details of a specific task by its ID
// @Tags tasks
// @Produce json
// @Param id path string true "Task ID"
// @Success 200 {array} task_dto.TaskListingDTO
// @Success 200 {array} section_dto.SectionBasicDTO
// @Success 200 {array} user_dto.UserBasicDTO
// @Success 200 {array} label_dto.LabelListingDTO
// @Failure 404 {string} string "Task not found"
// @Router /task/{id} [get]
func (c *TaskController) GetTaskByID(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, _ := strconv.ParseUint(params["id"], 10, 32)

	taskDTO, err := c.Service.GetTaskByID(uint(id))
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(taskDTO)
}

// CreateTask creates a new task
// @Summary Create a new task
// @Description Creates a new task with the provided details
// @Tags tasks
// @Accept json
// @Produce json
// @Param task body models.Task true "Task data"
// @Success 201 {string} string "Task created successfully"
// @Failure 400 {string} string "Invalid data"
// @Failure 500 {string} string "Internal Server Error"
// @Router /task [post]
func (c *TaskController) CreateTask(w http.ResponseWriter, r *http.Request) {
	var task models.Task

	err := json.NewDecoder(r.Body).Decode(&task)
	if err != nil {
		http.Error(w, "Invalid data", http.StatusBadRequest)
		return
	}

	err = c.Service.CreateTask(task)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

// UpdateTask updates an existing task
// @Summary Update a task
// @Description Updates the details of an existing task by its ID
// @Tags tasks
// @Accept json
// @Produce json
// @Param id path string true "Task ID"
// @Param task body models.Task true "Updated task data"
// @Success 204 "Task updated successfully"
// @Failure 400 {string} string "Invalid data"
// @Failure 500 {string} string "Internal Server Error"
// @Router /task/{id} [put]
func (c *TaskController) UpdateTask(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, _ := strconv.ParseUint(params["id"], 10, 32)

	var task models.Task
	err := json.NewDecoder(r.Body).Decode(&task)
	if err != nil {
		http.Error(w, "Invalid data", http.StatusBadRequest)
		return
	}

	err = c.Service.UpdateTask(uint(id), task)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

// DeleteTask deletes a task by its ID
// @Summary Delete a task by ID
// @Description Deletes an existing task from the system by its ID
// @Tags tasks
// @Param id path string true "Task ID"
// @Success 204 "Task deleted successfully"
// @Failure 500 {string} string "Internal Server Error"
// @Router /task/{id} [delete]
func (c *TaskController) DeleteTask(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, _ := strconv.ParseUint(params["id"], 10, 32)

	err := c.Service.DeleteTask(uint(id))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

// AssignLabelsToTask assigns labels to a task
// @Summary Assign labels to a task
// @Description Associates a set of labels with a specific task by its ID
// @Tags tasks
// @Accept json
// @Produce json
// @Param task_id path string true "Task ID"
// @Param labels body []models.Label true "List of labels"
// @Success 204 "Labels assigned successfully"
// @Failure 400 {string} string "Invalid data"
// @Failure 500 {string} string "Internal Server Error"
// @Router /task/{task_id}/labels [post]
func (c *TaskController) AssignLabelsToTask(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	taskID, _ := strconv.ParseUint(params["task_id"], 10, 32)

	var labels []models.Label
	err := json.NewDecoder(r.Body).Decode(&labels)
	if err != nil {
		http.Error(w, "Invalid data", http.StatusBadRequest)
		return
	}

	err = c.Service.AssignLabelsToTask(uint(taskID), labels)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
