package controllers

import (
	"awesomeProject/models"
	"awesomeProject/service"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

type LabelController struct {
	Service *service.LabelService
}

// GetLabels retrieves all labels
// @Summary Get all labels
// @Description Fetches all labels available in the system
// @Tags labels
// @Produce json
// @Success 200 {array} label_dto.LabelListingDTO
// @Failure 500 {string} string "Internal Server Error"
// @Router /label [get]
func (c *LabelController) GetLabels(ctx echo.Context) error {
	labelsDTO, err := c.Service.GetAllLabels()
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return ctx.JSON(http.StatusOK, labelsDTO)
}

// GetLabelByID retrieves a label by its ID
// @Summary Get a label by ID
// @Description Fetches the details of a specific label by its ID
// @Tags labels
// @Produce json
// @Param id path string true "Label ID"
// @Success 200 {array} label_dto.LabelListingDTO
// @Failure 404 {string} string "Label not found"
// @Router /label/{id} [get]
func (c *LabelController) GetLabelByID(ctx echo.Context) error {
	id, _ := strconv.ParseUint(ctx.Param("id"), 10, 32)

	labelDTO, err := c.Service.GetLabelByID(uint(id))
	if err != nil {
		return ctx.JSON(http.StatusNotFound, map[string]string{"error": "Label not found"})
	}

	return ctx.JSON(http.StatusOK, labelDTO)
}

// CreateLabel creates a new label
// @Summary Create a new label
// @Description Creates a new label with the provided details
// @Tags labels
// @Accept json
// @Produce json
// @Param label body models.Label true "Label data"
// @Success 201 {string} string "Label created successfully"
// @Failure 400 {string} string "Invalid data"
// @Failure 500 {string} string "Internal Server Error"
// @Router /label [post]
func (c *LabelController) CreateLabel(ctx echo.Context) error {
	var label models.Label

	if err := ctx.Bind(&label); err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid data"})
	}

	if err := c.Service.CreateLabel(label); err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to create label"})
	}

	return ctx.JSON(http.StatusCreated, "Label created successfully")
}

// UpdateLabel updates an existing label
// @Summary Update a label
// @Description Updates the details of an existing label by its ID
// @Tags labels
// @Accept json
// @Produce json
// @Param id path string true "Label ID"
// @Param label body models.Label true "Updated label data"
// @Success 204 "Label updated successfully"
// @Failure 400 {string} string "Invalid data"
// @Failure 500 {string} string "Internal Server Error"
// @Router /label/{id} [put]
func (c *LabelController) UpdateLabel(ctx echo.Context) error {
	id, _ := strconv.ParseUint(ctx.Param("id"), 10, 32)

	var label models.Label
	if err := ctx.Bind(&label); err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid data"})
	}

	if err := c.Service.UpdateLabel(uint(id), label); err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to update label"})
	}

	return ctx.NoContent(http.StatusNoContent)
}

// DeleteLabel deletes a label by its ID
// @Summary Delete a label by ID
// @Description Deletes an existing label from the system by its ID
// @Tags labels
// @Param id path string true "Label ID"
// @Success 204 "Label deleted successfully"
// @Failure 500 {string} string "Internal Server Error"
// @Router /label/{id} [delete]
func (c *LabelController) DeleteLabel(ctx echo.Context) error {
	id, _ := strconv.ParseUint(ctx.Param("id"), 10, 32)

	if err := c.Service.DeleteLabel(uint(id)); err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to delete label"})
	}

	return ctx.NoContent(http.StatusNoContent)
}
