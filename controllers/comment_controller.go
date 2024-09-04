package controllers

import (
	"awesomeProject/dto/comment_dto"
	"awesomeProject/service"
	"github.com/labstack/echo/v4"
	"log"
	"net/http"
	"strconv"
)

type CommentController struct {
	Service *service.CommentService
}

// GetComments retrieves all comments
// @Summary Get all comments
// @Description Fetches all comments available in the system
// @Tags comments
// @Produce json
// @Success 200 {object} comment_dto.CommentListingDTO
// @Failure 500 {string} string "Internal Server Error"
// @Router /comment [get]
func (c *CommentController) GetComments(ctx echo.Context) error {
	commentsDTO, err := c.Service.GetAllComments()
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return ctx.JSON(http.StatusOK, commentsDTO)
}

// GetCommentByID retrieves a comment by its ID
// @Summary Get a comment by ID
// @Description Fetches the details of a specific comment by its ID
// @Tags comments
// @Produce json
// @Param id path string true "Comment ID"
// @Success 200 {object} comment_dto.CommentListingDTO
// @Failure 404 {string} string "Comment not found"
// @Router /comment/{id} [get]
func (c *CommentController) GetCommentByID(ctx echo.Context) error {
	id, _ := strconv.ParseUint(ctx.Param("id"), 10, 32)

	commentDTO, err := c.Service.GetCommentByID(uint(id))
	if err != nil {
		return ctx.JSON(http.StatusNotFound, map[string]string{"error": "Comment not found"})
	}
	return ctx.JSON(http.StatusOK, commentDTO)
}

// CreateComment creates a new comment
// @Summary Create a new comment
// @Description Creates a new comment with the provided details, including optional image upload
// @Tags comments
// @Accept multipart/form-data
// @Produce json
// @Param content formData string true "Content of the comment"
// @Param user_id formData uint true "ID of the user creating the comment"
// @Param task_id formData uint true "ID of the task associated with the comment"
// @Param image formData file false "Optional image file"
// @Success 201 {string} string "Comment created successfully"
// @Failure 400 {string} string "Invalid data"
// @Failure 500 {string} string "Internal Server Error"
// @Router /comment [post]
func (c *CommentController) CreateComment(ctx echo.Context) error {
	err := ctx.Request().ParseMultipartForm(10 << 20) // Limit to 10 MB
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{"error": "Error parsing form"})
	}

	content := ctx.FormValue("content")
	userID, err := strconv.ParseUint(ctx.FormValue("user_id"), 10, 32)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid user ID"})
	}
	taskID, err := strconv.ParseUint(ctx.FormValue("task_id"), 10, 32)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid task ID"})
	}

	file, fileHeader, err := ctx.FormFile("image")
	var imageURL string
	if err != nil && err.Error() != "http: no such file" {
		log.Printf("Error getting file: %v", err)
		return ctx.JSON(http.StatusBadRequest, map[string]string{"error": "Error getting file"})
	}
	defer func() {
		if file != nil {
			file.Close()
		}
	}()

	if file != nil && fileHeader != nil {
		// Upload the image to S3
		imageURL, err = service.UploadFileToS3(file, fileHeader.Filename)
		if err != nil {
			log.Printf("Error uploading file to S3: %v", err)
			return ctx.JSON(http.StatusInternalServerError, map[string]string{"error": "Error uploading image"})
		}
		log.Printf("Image successfully uploaded to S3. URL: %s", imageURL)
	}

	err = c.Service.CreateComment(content, imageURL, int(userID), int(taskID))
	if err != nil {
		log.Printf("Error creating comment: %v", err)
		return ctx.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to create comment"})
	}

	return ctx.JSON(http.StatusCreated, "Comment created successfully")
}

// UpdateComment updates an existing comment
// @Summary Update a comment
// @Description Updates the content and optional image URL of an existing comment by its ID
// @Tags comments
// @Accept json
// @Produce json
// @Param id path string true "Comment ID"
// @Param comment body models.Comment true "Updated comment data"
// @Success 204 "Comment updated successfully"
// @Failure 400 {string} string "Invalid data"
// @Failure 500 {string} string "Internal Server Error"
// @Router /comment/{id} [put]
func (c *CommentController) UpdateComment(ctx echo.Context) error {
	id, _ := strconv.ParseUint(ctx.Param("id"), 10, 32)

	var requestBody comment_dto.UpdateCommentRequestDTO
	if err := ctx.Bind(&requestBody); err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid data"})
	}

	err := c.Service.UpdateComment(uint(id), requestBody.Content, requestBody.ImageURL)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to update comment"})
	}

	return ctx.NoContent(http.StatusNoContent)
}

// DeleteComment deletes a comment by its ID
// @Summary Delete a comment by ID
// @Description Deletes an existing comment from the system by its ID
// @Tags comments
// @Param id path string true "Comment ID"
// @Success 204 "Comment deleted successfully"
// @Failure 500 {string} string "Internal Server Error"
// @Router /comment/{id} [delete]
func (c *CommentController) DeleteComment(ctx echo.Context) error {
	id, _ := strconv.ParseUint(ctx.Param("id"), 10, 32)

	err := c.Service.DeleteComment(uint(id))
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to delete comment"})
	}

	return ctx.NoContent(http.StatusNoContent)
}
