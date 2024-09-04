package controllers

import (
	"awesomeProject/models"
	"awesomeProject/service"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

type TaskController struct {
	Service *service.TaskService
}

// GetTasks retrieves all tasks
// @Summary Get all tasks
// @Description Fetches all tasks available in the system
// @Tags tasks
// @Produce json
// @Success 200 {array} task_dto.TaskListingDTO
// @Success 200 {array} section_dto.SectionBasicDTO
// @Success 200 {array} user_dto.UserBasicDTO
// @Success 200 {array} label_dto.LabelListingDTO
// @Failure 500 {string} string "Internal Server Error"
// @Router /task [get]
func (c *TaskController) GetTasks(ctx echo.Context) error {
	tasksDTO, err := c.Service.GetAllTasks()
	if err != nil {
		return ctx.String(http.StatusInternalServerError, err.Error())
	}

	return ctx.JSON(http.StatusOK, tasksDTO)
}

// GetTaskByID retrieves a task by its ID
// @Summary Get a task by ID
// @Description Fetches the details of a specific task by its ID
// @Tags tasks
// @Produce json
// @Param id path string true "Task ID"
// @Success 200 {array} task_dto.TaskListingDTO
// @Success 200 {array} section_dto.SectionBasicDTO
// @Success 200 {array} user_dto.UserBasicDTO
// @Success 200 {array} label_dto.LabelListingDTO
// @Failure 404 {string} string "Task not found"
// @Router /task/{id} [get]
func (c *TaskController) GetTaskByID(ctx echo.Context) error {
	id, _ := strconv.ParseUint(ctx.Param("id"), 10, 32)

	taskDTO, err := c.Service.GetTaskByID(uint(id))
	if err != nil {
		return ctx.String(http.StatusNotFound, err.Error())
	}

	return ctx.JSON(http.StatusOK, taskDTO)
}

// CreateTask creates a new task
// @Summary Create a new task
// @Description Creates a new task with the provided details
// @Tags tasks
// @Accept json
// @Produce json
// @Param task body models.Task true "Task data"
// @Success 201 {string} string "Task created successfully"
// @Failure 400 {string} string "Invalid data"
// @Failure 500 {string} string "Internal Server Error"
// @Router /task [post]
func (c *TaskController) CreateTask(ctx echo.Context) error {
	var task models.Task

	if err := ctx.Bind(&task); err != nil {
		return ctx.String(http.StatusBadRequest, "Invalid data")
	}

	if err := c.Service.CreateTask(task); err != nil {
		return ctx.String(http.StatusInternalServerError, err.Error())
	}

	return ctx.String(http.StatusCreated, "Task created successfully")
}

// UpdateTask updates an existing task
// @Summary Update a task
// @Description Updates the details of an existing task by its ID
// @Tags tasks
// @Accept json
// @Produce json
// @Param id path string true "Task ID"
// @Param task body models.Task true "Updated task data"
// @Success 204 "Task updated successfully"
// @Failure 400 {string} string "Invalid data"
// @Failure 500 {string} string "Internal Server Error"
// @Router /task/{id} [put]
func (c *TaskController) UpdateTask(ctx echo.Context) error {
	id, _ := strconv.ParseUint(ctx.Param("id"), 10, 32)

	var task models.Task
	if err := ctx.Bind(&task); err != nil {
		return ctx.String(http.StatusBadRequest, "Invalid data")
	}

	if err := c.Service.UpdateTask(uint(id), task); err != nil {
		return ctx.String(http.StatusInternalServerError, err.Error())
	}

	return ctx.NoContent(http.StatusNoContent)
}

// DeleteTask deletes a task by its ID
// @Summary Delete a task by ID
// @Description Deletes an existing task from the system by its ID
// @Tags tasks
// @Param id path string true "Task ID"
// @Success 204 "Task deleted successfully"
// @Failure 500 {string} string "Internal Server Error"
// @Router /task/{id} [delete]
func (c *TaskController) DeleteTask(ctx echo.Context) error {
	id, _ := strconv.ParseUint(ctx.Param("id"), 10, 32)

	if err := c.Service.DeleteTask(uint(id)); err != nil {
		return ctx.String(http.StatusInternalServerError, err.Error())
	}

	return ctx.NoContent(http.StatusNoContent)
}

// AssignLabelsToTask assigns labels to a task
// @Summary Assign labels to a task
// @Description Associates a set of labels with a specific task by its ID
// @Tags tasks
// @Accept json
// @Produce json
// @Param task_id path string true "Task ID"
// @Param labels body []models.Label true "List of labels"
// @Success 204 "Labels assigned successfully"
// @Failure 400 {string} string "Invalid data"
// @Failure 500 {string} string "Internal Server Error"
// @Router /task/{task_id}/labels [post]
func (c *TaskController) AssignLabelsToTask(ctx echo.Context) error {
	taskID, _ := strconv.ParseUint(ctx.Param("task_id"), 10, 32)

	var labels []models.Label
	if err := ctx.Bind(&labels); err != nil {
		return ctx.String(http.StatusBadRequest, "Invalid data")
	}

	if err := c.Service.AssignLabelsToTask(uint(taskID), labels); err != nil {
		return ctx.String(http.StatusInternalServerError, err.Error())
	}

	return ctx.NoContent(http.StatusNoContent)
}
