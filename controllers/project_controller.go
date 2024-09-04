package controllers

import (
	"awesomeProject/models"
	"awesomeProject/service"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

type ProjectController struct {
	Service service.ProjectService
}

func NewProjectController(svc service.ProjectService) *ProjectController {
	return &ProjectController{
		Service: svc,
	}
}

// GetProjects retrieves all projects
// @Summary Get all projects
// @Description Fetches all projects available in the system
// @Tags projects
// @Produce json
// @Success 200 {object} []models.Project
// @Failure 500 {string} string "Internal Server Error"
// @Router /project [get]
func (c *ProjectController) GetProjects(ctx echo.Context) error {
	projects, err := c.Service.GetProjects()
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, "Error fetching projects")
	}

	return ctx.JSON(http.StatusOK, projects)
}

// GetProjectByID retrieves a project by its ID
// @Summary Get a project by ID
// @Description Fetches the details of a specific project by its ID
// @Tags projects
// @Produce json
// @Param id path string true "Project ID"
// @Success 200 {object} models.Project
// @Failure 404 {string} string "Project not found"
// @Router /project/{id} [get]
func (c *ProjectController) GetProjectByID(ctx echo.Context) error {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, "Invalid ID format")
	}

	project, err := c.Service.GetProjectByID(uint(id))
	if err != nil {
		return ctx.JSON(http.StatusNotFound, "Project not found")
	}

	return ctx.JSON(http.StatusOK, project)
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
func (c *ProjectController) CreateProject(ctx echo.Context) error {
	var project models.Project

	if err := ctx.Bind(&project); err != nil {
		return ctx.JSON(http.StatusBadRequest, "Invalid data")
	}

	if err := c.Service.CreateProject(project); err != nil {
		return ctx.JSON(http.StatusInternalServerError, "Error creating project")
	}

	return ctx.NoContent(http.StatusCreated)
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
func (c *ProjectController) UpdateProject(ctx echo.Context) error {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, "Invalid ID format")
	}

	var project models.Project
	if err := ctx.Bind(&project); err != nil {
		return ctx.JSON(http.StatusBadRequest, "Invalid data")
	}

	if err := c.Service.UpdateProject(uint(id), project); err != nil {
		return ctx.JSON(http.StatusInternalServerError, "Error updating project")
	}

	return ctx.NoContent(http.StatusNoContent)
}

// DeleteProject deletes a project by its ID
// @Summary Delete a project by ID
// @Description Deletes an existing project from the system by its ID
// @Tags projects
// @Param id path string true "Project ID"
// @Success 204 "Project deleted successfully"
// @Failure 500 {string} string "Internal Server Error"
// @Router /project/{id} [delete]
func (c *ProjectController) DeleteProject(ctx echo.Context) error {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, "Invalid ID format")
	}

	if err := c.Service.DeleteProject(uint(id)); err != nil {
		return ctx.JSON(http.StatusInternalServerError, "Error deleting project")
	}

	return ctx.NoContent(http.StatusNoContent)
}
