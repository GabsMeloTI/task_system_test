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

// GetSections retrieves all sections
// @Summary Get all sections
// @Description Fetches all sections available in the system
// @Tags sections
// @Produce json
// @Success 200 {array} section_dto.SectionListingDTO
// @Failure 500 {string} string "Internal Server Error"
// @Router /section [get]
func (c *SectionController) GetSections(w http.ResponseWriter, r *http.Request) {
	sections, err := c.Service.GetSections()
	if err != nil {
		http.Error(w, "Error fetching sections", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(sections)
}

// GetSectionByID retrieves a section by its ID
// @Summary Get a section by ID
// @Description Fetches the details of a specific section by its ID
// @Tags sections
// @Produce json
// @Param id path string true "Section ID"
// @Success 200 {array} section_dto.SectionListingDTO
// @Success 200 {array} project_dto.ProjectBasicDTO
// @Success 200 {array} user_dto.UserBasicDTO
// @Failure 404 {string} string "Section not found"
// @Router /section/{id} [get]
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

// CreateSection creates a new section
// @Summary Create a new section
// @Description Creates a new section with the provided details
// @Tags sections
// @Accept json
// @Produce json
// @Param section body models.Section true "Section data"
// @Success 201 {string} string "Section created successfully"
// @Failure 400 {string} string "Invalid data"
// @Failure 500 {string} string "Internal Server Error"
// @Router /section [post]
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

// UpdateSection updates an existing section
// @Summary Update a section
// @Description Updates the details of an existing section by its ID
// @Tags sections
// @Accept json
// @Produce json
// @Param id path string true "Section ID"
// @Param section body models.Section true "Updated section data"
// @Success 204 "Section updated successfully"
// @Failure 400 {string} string "Invalid data"
// @Failure 500 {string} string "Internal Server Error"
// @Router /section/{id} [put]
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

// DeleteSection deletes a section by its ID
// @Summary Delete a section by ID
// @Description Deletes an existing section from the system by its ID
// @Tags sections
// @Param id path string true "Section ID"
// @Success 204 "Section deleted successfully"
// @Failure 500 {string} string "Internal Server Error"
// @Router /section/{id} [delete]
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
