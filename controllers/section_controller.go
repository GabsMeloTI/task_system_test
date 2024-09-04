package controllers

import (
	"awesomeProject/models"
	"awesomeProject/service"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

type SectionController struct {
	Service service.SectionService
}

func NewSectionController(svc service.SectionService) *SectionController {
	return &SectionController{
		Service: svc,
	}
}

// GetSections retrieves all sections
// @Summary Get all sections
// @Description Fetches all sections available in the system
// @Tags sections
// @Produce json
// @Success 200 {array} models.Section
// @Failure 500 {string} string "Internal Server Error"
// @Router /section [get]
func (c *SectionController) GetSections(ctx echo.Context) error {
	sections, err := c.Service.GetSections()
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, "Error fetching sections")
	}

	return ctx.JSON(http.StatusOK, sections)
}

// GetSectionByID retrieves a section by its ID
// @Summary Get a section by ID
// @Description Fetches the details of a specific section by its ID
// @Tags sections
// @Produce json
// @Param id path string true "Section ID"
// @Success 200 {object} models.Section
// @Failure 404 {string} string "Section not found"
// @Router /section/{id} [get]
func (c *SectionController) GetSectionByID(ctx echo.Context) error {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, "Invalid ID format")
	}

	section, err := c.Service.GetSectionByID(uint(id))
	if err != nil {
		return ctx.JSON(http.StatusNotFound, "Section not found")
	}

	return ctx.JSON(http.StatusOK, section)
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
func (c *SectionController) CreateSection(ctx echo.Context) error {
	var section models.Section

	if err := ctx.Bind(&section); err != nil {
		return ctx.JSON(http.StatusBadRequest, "Invalid data")
	}

	if err := c.Service.CreateSection(section); err != nil {
		return ctx.JSON(http.StatusInternalServerError, "Error creating section")
	}

	return ctx.NoContent(http.StatusCreated)
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
func (c *SectionController) UpdateSection(ctx echo.Context) error {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, "Invalid ID format")
	}

	var section models.Section
	if err := ctx.Bind(&section); err != nil {
		return ctx.JSON(http.StatusBadRequest, "Invalid data")
	}

	if err := c.Service.UpdateSection(uint(id), section); err != nil {
		return ctx.JSON(http.StatusInternalServerError, "Error updating section")
	}

	return ctx.NoContent(http.StatusNoContent)
}

// DeleteSection deletes a section by its ID
// @Summary Delete a section by ID
// @Description Deletes an existing section from the system by its ID
// @Tags sections
// @Param id path string true "Section ID"
// @Success 204 "Section deleted successfully"
// @Failure 500 {string} string "Internal Server Error"
// @Router /section/{id} [delete]
func (c *SectionController) DeleteSection(ctx echo.Context) error {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, "Invalid ID format")
	}

	if err := c.Service.DeleteSection(uint(id)); err != nil {
		return ctx.JSON(http.StatusInternalServerError, "Error deleting section")
	}

	return ctx.NoContent(http.StatusNoContent)
}
