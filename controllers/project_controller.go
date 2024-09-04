package controllers

import (
	"awesomeProject/models"
	"awesomeProject/service"
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

type ProjectController struct {
	Service service.ProjectService
}

// GetProjects retrieves all projects
// @Summary Get all projects
// @Description Fetches all projects available in the system
// @Tags projects
// @Produce json
// @Success 200 {object} project_dto.ProjectListingDTO
// @Success 200 {object} user_dto.UserBasicDTO
// @Failure 500 {string} string "Internal Server Error"
// @Router /project [get]
func (c *ProjectController) GetProjects(w http.ResponseWriter, r *http.Request) {
	projects, err := c.Service.GetProjects()
	if err != nil {
		http.Error(w, "Error fetching projects", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(projects)
}

// GetProjectByID retrieves a project by its ID
// @Summary Get a project by ID
// @Description Fetches the details of a specific project by its ID
// @Tags projects
// @Produce json
// @Param id path string true "Project ID"
// @Success 200 {object} project_dto.ProjectListingDTO
// @Success 200 {object} user_dto.UserBasicDTO
// @Failure 404 {string} string "Project not found"
// @Router /project/{id} [get]
func (c *ProjectController) GetProjectByID(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, _ := strconv.ParseUint(params["id"], 10, 32)

	project, err := c.Service.GetProjectByID(uint(id))
	if err != nil {
		http.Error(w, "Project not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(project)
}

// CreateProject creates a new project
// @Summary Create a new project
// @Description Creates a new project with the provided details
// @Tags projects
// @Accept json
// @Produce json
// @Param project body models.Project true "Project data"
// @Success 201 {string} string "Project created successfully"
// @Failure 400 {string} string "Invalid data"
// @Failure 500 {string} string "Internal Server Error"
// @Router /project [post]
func (c *ProjectController) CreateProject(w http.ResponseWriter, r *http.Request) {
	var project models.Project

	err := json.NewDecoder(r.Body).Decode(&project)
	if err != nil {
		http.Error(w, "Invalid data", http.StatusBadRequest)
		return
	}

	err = c.Service.CreateProject(project)
	if err != nil {
		http.Error(w, "Error creating project", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

// UpdateProject updates an existing project
// @Summary Update a project
// @Description Updates the details of an existing project by its ID
// @Tags projects
// @Accept json
// @Produce json
// @Param id path string true "Project ID"
// @Param project body models.Project true "Updated project data"
// @Success 204 "Project updated successfully"
// @Failure 400 {string} string "Invalid data"
// @Failure 500 {string} string "Internal Server Error"
// @Router /project/{id} [put]
func (c *ProjectController) UpdateProject(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, _ := strconv.ParseUint(params["id"], 10, 32)

	var project models.Project
	err := json.NewDecoder(r.Body).Decode(&project)
	if err != nil {
		http.Error(w, "Invalid data", http.StatusBadRequest)
		return
	}

	err = c.Service.UpdateProject(uint(id), project)
	if err != nil {
		http.Error(w, "Error updating project", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

// DeleteProject deletes a project by its ID
// @Summary Delete a project by ID
// @Description Deletes an existing project from the system by its ID
// @Tags projects
// @Param id path string true "Project ID"
// @Success 204 "Project deleted successfully"
// @Failure 500 {string} string "Internal Server Error"
// @Router /project/{id} [delete]
func (c *ProjectController) DeleteProject(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, _ := strconv.ParseUint(params["id"], 10, 32)

	err := c.Service.DeleteProject(uint(id))
	if err != nil {
		http.Error(w, "Error deleting project", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
