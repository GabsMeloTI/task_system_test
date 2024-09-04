package controllers

import (
	"awesomeProject/models"
	"awesomeProject/service"
	"github.com/labstack/echo/v4"
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
func (c *SubtaskController) GetSubtasks(ctx echo.Context) error {
	subtasksDTO, err := c.Service.GetAllSubtasks()
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return ctx.JSON(http.StatusOK, subtasksDTO)
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
func (c *SubtaskController) GetSubtaskByID(ctx echo.Context) error {
	id, _ := strconv.ParseUint(ctx.Param("id"), 10, 32)

	subtaskDTO, err := c.Service.GetSubtaskByID(uint(id))
	if err != nil {
		return ctx.JSON(http.StatusNotFound, map[string]string{"error": err.Error()})
	}

	return ctx.JSON(http.StatusOK, subtaskDTO)
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
func (c *SubtaskController) CreateSubtask(ctx echo.Context) error {
	var subtask models.Subtask

	if err := ctx.Bind(&subtask); err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid data"})
	}

	if err := c.Service.CreateSubtask(subtask); err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return ctx.JSON(http.StatusCreated, "Subtask created successfully")
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
func (c *SubtaskController) UpdateSubtask(ctx echo.Context) error {
	id, _ := strconv.ParseUint(ctx.Param("id"), 10, 32)

	var subtask models.Subtask
	if err := ctx.Bind(&subtask); err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid data"})
	}

	if err := c.Service.UpdateSubtask(uint(id), subtask); err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return ctx.NoContent(http.StatusNoContent)
}

// DeleteSubtask deletes a subtask by its ID
// @Summary Delete a subtask by ID
// @Description Deletes an existing subtask from the system by its ID
// @Tags subtasks
// @Param id path string true "Subtask ID"
// @Success 204 "Subtask deleted successfully"
// @Failure 500 {string} string "Internal Server Error"
// @Router /subtask/{id} [delete]
func (c *SubtaskController) DeleteSubtask(ctx echo.Context) error {
	id, _ := strconv.ParseUint(ctx.Param("id"), 10, 32)

	if err := c.Service.DeleteSubtask(uint(id)); err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return ctx.NoContent(http.StatusNoContent)
}
