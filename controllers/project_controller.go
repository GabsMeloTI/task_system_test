package controllers

import (
	"awesomeProject/db"
	"awesomeProject/dto/project_dto"
	"awesomeProject/dto/user_dto"
	"awesomeProject/models"
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
)

func GetProjects(w http.ResponseWriter, r *http.Request) {
	var projects []models.Project
	var projectsDTO []project_dto.ProjectListingDTO

	db.DB.Preload("User").Find(&projects)

	for _, project := range projects {
		projectDTO := project_dto.ProjectListingDTO{
			ID:          project.ID,
			Title:       project.Title,
			Description: project.Description,
			Status:      project.Status,
			CreatedAt:   project.CreatedAt,
			User: user_dto.UserBasicDTO{
				ID:   project.User.ID,
				Name: project.User.Name,
			},
		}
		projectsDTO = append(projectsDTO, projectDTO)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(projectsDTO)
}

func GetProjectByID(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var project models.Project

	err := db.DB.Preload("User").First(&project, params["id"]).Error
	if err != nil {
		http.Error(w, "Project not found", http.StatusNotFound)
		return
	}

	projectDTO := project_dto.ProjectListingDTO{
		ID:          project.ID,
		Title:       project.Title,
		Description: project.Description,
		Status:      project.Status,
		CreatedAt:   project.CreatedAt,
		User: user_dto.UserBasicDTO{
			ID:   project.User.ID,
			Name: project.User.Name,
		},
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(projectDTO)
}

func CreateProject(w http.ResponseWriter, r *http.Request) {
	var project models.Project

	err := json.NewDecoder(r.Body).Decode(&project)
	if err != nil {
		http.Error(w, "Invalid data", http.StatusBadRequest)
		return
	}

	err = db.DB.Create(&project).Error
	if err != nil {
		http.Error(w, "Error creating project", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func UpdateProject(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var project models.Project

	err := db.DB.First(&project, params["id"]).Error
	if err != nil {
		http.Error(w, "Project not found", http.StatusNotFound)
		return
	}

	err = json.NewDecoder(r.Body).Decode(&project)
	if err != nil {
		http.Error(w, "Invalid data", http.StatusBadRequest)
		return
	}

	db.DB.Save(&project)
	w.WriteHeader(http.StatusNoContent)
}

func DeleteProject(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	err := db.DB.Delete(&models.Project{}, params["id"]).Error
	if err != nil {
		http.Error(w, "Error deleting project", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
