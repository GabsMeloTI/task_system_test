package controllers

import (
	"awesomeProject/models"
	"awesomeProject/service"
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

type SectionController struct {
	Service service.SectionService
}

func (c *SectionController) GetSections(w http.ResponseWriter, r *http.Request) {
	sections, err := c.Service.GetSections()
	if err != nil {
		http.Error(w, "Error fetching sections", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(sections)
}

func (c *SectionController) GetSectionByID(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, _ := strconv.ParseUint(params["id"], 10, 32)

	section, err := c.Service.GetSectionByID(uint(id))
	if err != nil {
		http.Error(w, "Section not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(section)
}

func (c *SectionController) CreateSection(w http.ResponseWriter, r *http.Request) {
	var section models.Section

	err := json.NewDecoder(r.Body).Decode(&section)
	if err != nil {
		http.Error(w, "Invalid data", http.StatusBadRequest)
		return
	}

	err = c.Service.CreateSection(section)
	if err != nil {
		http.Error(w, "Error creating section", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func (c *SectionController) UpdateSection(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, _ := strconv.ParseUint(params["id"], 10, 32)

	var section models.Section
	err := json.NewDecoder(r.Body).Decode(&section)
	if err != nil {
		http.Error(w, "Invalid data", http.StatusBadRequest)
		return
	}

	err = c.Service.UpdateSection(uint(id), section)
	if err != nil {
		http.Error(w, "Error updating section", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func (c *SectionController) DeleteSection(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, _ := strconv.ParseUint(params["id"], 10, 32)

	err := c.Service.DeleteSection(uint(id))
	if err != nil {
		http.Error(w, "Error deleting section", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
