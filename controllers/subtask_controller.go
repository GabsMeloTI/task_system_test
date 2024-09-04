package controllers

import (
	"awesomeProject/models"
	"awesomeProject/service"
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

type SubtaskController struct {
	Service *service.SubtaskService
}

// GetSubtasks retrieves all subtasks
// @Summary Get all subtasks
// @Description Fetches all subtasks available in the system
// @Tags subtasks
// @Produce json
// @Success 200 {array} subtask_dto.SubtaskListingDTO
// @Success 200 {array} task_dto.TaskBasicDTO
// @Failure 500 {string} string "Internal Server Error"
// @Router /subtask [get]
func (c *SubtaskController) GetSubtasks(w http.ResponseWriter, r *http.Request) {
	subtasksDTO, err := c.Service.GetAllSubtasks()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(subtasksDTO)
}

// GetSubtaskByID retrieves a subtask by its ID
// @Summary Get a subtask by ID
// @Description Fetches the details of a specific subtask by its ID
// @Tags subtasks
// @Produce json
// @Param id path string true "Subtask ID"
// @Success 200 {array} subtask_dto.SubtaskListingDTO
// @Success 200 {array} task_dto.TaskBasicDTO
// @Failure 404 {string} string "Subtask not found"
// @Router /subtask/{id} [get]
func (c *SubtaskController) GetSubtaskByID(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, _ := strconv.ParseUint(params["id"], 10, 32)

	subtaskDTO, err := c.Service.GetSubtaskByID(uint(id))
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(subtaskDTO)
}

// CreateSubtask creates a new subtask
// @Summary Create a new subtask
// @Description Creates a new subtask with the provided details
// @Tags subtasks
// @Accept json
// @Produce json
// @Param subtask body models.Subtask true "Subtask data"
// @Success 201 {string} string "Subtask created successfully"
// @Failure 400 {string} string "Invalid data"
// @Failure 500 {string} string "Internal Server Error"
// @Router /subtask [post]
func (c *SubtaskController) CreateSubtask(w http.ResponseWriter, r *http.Request) {
	var subtask models.Subtask

	err := json.NewDecoder(r.Body).Decode(&subtask)
	if err != nil {
		http.Error(w, "Invalid data", http.StatusBadRequest)
		return
	}

	err = c.Service.CreateSubtask(subtask)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

// UpdateSubtask updates an existing subtask
// @Summary Update a subtask
// @Description Updates the details of an existing subtask by its ID
// @Tags subtasks
// @Accept json
// @Produce json
// @Param id path string true "Subtask ID"
// @Param subtask body models.Subtask true "Updated subtask data"
// @Success 204 "Subtask updated successfully"
// @Failure 400 {string} string "Invalid data"
// @Failure 500 {string} string "Internal Server Error"
// @Router /subtask/{id} [put]
func (c *SubtaskController) UpdateSubtask(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, _ := strconv.ParseUint(params["id"], 10, 32)

	var subtask models.Subtask
	err := json.NewDecoder(r.Body).Decode(&subtask)
	if err != nil {
		http.Error(w, "Invalid data", http.StatusBadRequest)
		return
	}

	err = c.Service.UpdateSubtask(uint(id), subtask)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

// DeleteSubtask deletes a subtask by its ID
// @Summary Delete a subtask by ID
// @Description Deletes an existing subtask from the system by its ID
// @Tags subtasks
// @Param id path string true "Subtask ID"
// @Success 204 "Subtask deleted successfully"
// @Failure 500 {string} string "Internal Server Error"
// @Router /subtask/{id} [delete]
func (c *SubtaskController) DeleteSubtask(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, _ := strconv.ParseUint(params["id"], 10, 32)

	err := c.Service.DeleteSubtask(uint(id))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
