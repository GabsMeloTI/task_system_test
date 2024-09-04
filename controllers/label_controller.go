package controllers

import (
	"awesomeProject/models"
	"awesomeProject/service"
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

type LabelController struct {
	Service *service.LabelService
}

// GetLabels retrieves all labels
// @Summary Get all labels
// @Description Fetches all labels available in the system
// @Tags labels
// @Produce json
// @Success 200 {array} label_dto.LabelListingDTO
// @Failure 500 {string} string "Internal Server Error"
// @Router /label [get]
func (c *LabelController) GetLabels(w http.ResponseWriter, r *http.Request) {
	labelsDTO, err := c.Service.GetAllLabels()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(labelsDTO)
}

// GetLabelByID retrieves a label by its ID
// @Summary Get a label by ID
// @Description Fetches the details of a specific label by its ID
// @Tags labels
// @Produce json
// @Param id path string true "Label ID"
// @Success 200 {array} label_dto.LabelListingDTO
// @Failure 404 {string} string "Label not found"
// @Router /label/{id} [get]
func (c *LabelController) GetLabelByID(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, _ := strconv.ParseUint(params["id"], 10, 32)

	labelDTO, err := c.Service.GetLabelByID(uint(id))
	if err != nil {
		http.Error(w, "Label not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(labelDTO)
}

// CreateLabel creates a new label
// @Summary Create a new label
// @Description Creates a new label with the provided details
// @Tags labels
// @Accept json
// @Produce json
// @Param label body models.Label true "Label data"
// @Success 201 {string} string "Label created successfully"
// @Failure 400 {string} string "Invalid data"
// @Failure 500 {string} string "Internal Server Error"
// @Router /label [post]
func (c *LabelController) CreateLabel(w http.ResponseWriter, r *http.Request) {
	var label models.Label

	if err := json.NewDecoder(r.Body).Decode(&label); err != nil {
		http.Error(w, "Invalid data", http.StatusBadRequest)
		return
	}

	err := c.Service.CreateLabel(label)
	if err != nil {
		http.Error(w, "Failed to create label", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

// UpdateLabel updates an existing label
// @Summary Update a label
// @Description Updates the details of an existing label by its ID
// @Tags labels
// @Accept json
// @Produce json
// @Param id path string true "Label ID"
// @Param label body models.Label true "Updated label data"
// @Success 204 "Label updated successfully"
// @Failure 400 {string} string "Invalid data"
// @Failure 500 {string} string "Internal Server Error"
// @Router /label/{id} [put]
func (c *LabelController) UpdateLabel(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, _ := strconv.ParseUint(params["id"], 10, 32)

	var label models.Label

	if err := json.NewDecoder(r.Body).Decode(&label); err != nil {
		http.Error(w, "Invalid data", http.StatusBadRequest)
		return
	}

	err := c.Service.UpdateLabel(uint(id), label)
	if err != nil {
		http.Error(w, "Failed to update label", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

// DeleteLabel deletes a label by its ID
// @Summary Delete a label by ID
// @Description Deletes an existing label from the system by its ID
// @Tags labels
// @Param id path string true "Label ID"
// @Success 204 "Label deleted successfully"
// @Failure 500 {string} string "Internal Server Error"
// @Router /label/{id} [delete]
func (c *LabelController) DeleteLabel(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, _ := strconv.ParseUint(params["id"], 10, 32)

	err := c.Service.DeleteLabel(uint(id))
	if err != nil {
		http.Error(w, "Failed to delete label", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
