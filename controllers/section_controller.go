package controllers

import (
	"awesomeProject/db"
	"awesomeProject/dto/project_dto"
	"awesomeProject/dto/section_dto"
	"awesomeProject/dto/user_dto"
	"awesomeProject/models"
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
)

func GetSections(w http.ResponseWriter, r *http.Request) {
	var sections []models.Section
	var sectionsDTO []section_dto.SectionListingDTO

	db.DB.Preload("Project").Preload("User").Find(&sections)

	for _, section := range sections {
		sectionDTO := section_dto.SectionListingDTO{
			ID:          section.ID,
			Title:       section.Title,
			Description: section.Description,
			CreatedAt:   section.CreatedAt,
			User: user_dto.UserBasicDTO{
				ID:   section.User.ID,
				Name: section.User.Name,
			},
			Project: project_dto.ProjectBasicDTO{
				ID:     section.Project.ID,
				Title:  section.Project.Title,
				Status: section.Project.Status,
			},
		}
		sectionsDTO = append(sectionsDTO, sectionDTO)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(sectionsDTO)
}

func GetSectionByID(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var section models.Section

	err := db.DB.Preload("Project").Preload("User").First(&section, params["id"]).Error
	if err != nil {
		http.Error(w, "Section not found", http.StatusNotFound)
		return
	}

	sectionDTO := section_dto.SectionListingDTO{
		ID:          section.ID,
		Title:       section.Title,
		Description: section.Description,
		CreatedAt:   section.CreatedAt,
		User: user_dto.UserBasicDTO{
			ID:   section.User.ID,
			Name: section.User.Name,
		},
		Project: project_dto.ProjectBasicDTO{
			ID:     section.Project.ID,
			Title:  section.Project.Title,
			Status: section.Project.Status,
		},
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(sectionDTO)
}

func CreateSection(w http.ResponseWriter, r *http.Request) {
	var section models.Section

	err := json.NewDecoder(r.Body).Decode(&section)
	if err != nil {
		http.Error(w, "Invalid data", http.StatusBadRequest)
		return
	}

	err = db.DB.Create(&section).Error
	if err != nil {
		http.Error(w, "Error creating section", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func UpdateSection(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var section models.Section

	err := db.DB.First(&section, params["id"]).Error
	if err != nil {
		http.Error(w, "Section not found", http.StatusNotFound)
		return
	}

	err = json.NewDecoder(r.Body).Decode(&section)
	if err != nil {
		http.Error(w, "Invalid data", http.StatusBadRequest)
		return
	}

	db.DB.Save(&section)
	w.WriteHeader(http.StatusNoContent)
}

func DeleteSection(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	err := db.DB.Delete(&models.Section{}, params["id"]).Error
	if err != nil {
		http.Error(w, "Error deleting section", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
