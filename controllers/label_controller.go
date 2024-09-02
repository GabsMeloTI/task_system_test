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

func (c *LabelController) GetLabels(w http.ResponseWriter, r *http.Request) {
	labelsDTO, err := c.Service.GetAllLabels()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(labelsDTO)
}

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
