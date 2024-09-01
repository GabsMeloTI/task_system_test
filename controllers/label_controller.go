package controllers

import (
	"awesomeProject/db"
	"awesomeProject/dto/label_dto"
	"awesomeProject/models"
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
)

func GetLabels(w http.ResponseWriter, r *http.Request) {
	var labels []models.Label
	var labelsDTO []label_dto.LabelListingDTO

	if err := db.DB.Preload("Task").Find(&labels).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	for _, label := range labels {
		labelDTO := label_dto.LabelListingDTO{
			ID:    label.ID,
			Name:  label.Name,
			Color: label.Color,
		}
		labelsDTO = append(labelsDTO, labelDTO)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(labelsDTO)
}

func GetLabelByID(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var label models.Label

	if err := db.DB.Preload("Task").First(&label, params["id"]).Error; err != nil {
		http.Error(w, "Label not found", http.StatusNotFound)
		return
	}

	labelDTO := label_dto.LabelListingDTO{
		ID:    label.ID,
		Name:  label.Name,
		Color: label.Color,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(labelDTO)
}

func CreateLabel(w http.ResponseWriter, r *http.Request) {
	var label models.Label

	if err := json.NewDecoder(r.Body).Decode(&label); err != nil {
		http.Error(w, "Invalid data", http.StatusBadRequest)
		return
	}

	if err := db.DB.Create(&label).Error; err != nil {
		http.Error(w, "Failed to create label", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func UpdateLabel(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var label models.Label

	if err := db.DB.First(&label, params["id"]).Error; err != nil {
		http.Error(w, "Label not found", http.StatusNotFound)
		return
	}

	if err := json.NewDecoder(r.Body).Decode(&label); err != nil {
		http.Error(w, "Invalid data", http.StatusBadRequest)
		return
	}

	if err := db.DB.Save(&label).Error; err != nil {
		http.Error(w, "Failed to update label", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func DeleteLabel(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	if err := db.DB.Delete(&models.Label{}, params["id"]).Error; err != nil {
		http.Error(w, "Failed to delete label", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
