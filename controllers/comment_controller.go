package controllers

import (
	"awesomeProject/dto/comment_dto"
	"awesomeProject/service"
	"encoding/json"
	"github.com/gorilla/mux"
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
// @Success 200 {array} task_dto.TaskBasicDTO
// @Success 200 {array} user_dto.UserBasicDTO
// @Failure 500 {string} string "Internal Server Error"
// @Router /comment [get]
func (c *CommentController) GetComments(w http.ResponseWriter, r *http.Request) {
	commentsDTO, err := c.Service.GetAllComments()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(commentsDTO)
}

// GetCommentByID retrieves a comment by its ID
// @Summary Get a comment by ID
// @Description Fetches the details of a specific comment by its ID
// @Tags comments
// @Produce json
// @Param id path string true "Comment ID"
// @Success 200 {object} comment_dto.CommentListingDTO
// @Success 200 {array} task_dto.TaskBasicDTO
// @Success 200 {array} user_dto.UserBasicDTO
// @Failure 404 {string} string "Comment not found"
// @Router /comment/{id} [get]
func (c *CommentController) GetCommentByID(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, _ := strconv.ParseUint(params["id"], 10, 32)

	commentDTO, err := c.Service.GetCommentByID(uint(id))
	if err != nil {
		http.Error(w, "Comment not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(commentDTO)
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
func (c *CommentController) CreateComment(w http.ResponseWriter, r *http.Request) {
	err := r.ParseMultipartForm(10 << 20) // Limit to 10 MB
	if err != nil {
		http.Error(w, "Error parsing form", http.StatusBadRequest)
		return
	}

	content := r.FormValue("content")
	userID, err := strconv.ParseUint(r.FormValue("user_id"), 10, 32)
	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}
	taskID, err := strconv.ParseUint(r.FormValue("task_id"), 10, 32)
	if err != nil {
		http.Error(w, "Invalid task ID", http.StatusBadRequest)
		return
	}

	file, fileHeader, err := r.FormFile("image")
	if err != nil {
		if err.Error() == "http: no such file" {
			file = nil
			fileHeader = nil
		} else {
			log.Printf("Error getting file: %v", err)
			http.Error(w, "Error getting file", http.StatusBadRequest)
			return
		}
	}
	defer func() {
		if file != nil {
			file.Close()
		}
	}()

	var imageURL string
	if file != nil && fileHeader != nil {
		// Upload the image to S3
		imageURL, err = service.UploadFileToS3(file, fileHeader.Filename)
		if err != nil {
			log.Printf("Error uploading file to S3: %v", err)
			http.Error(w, "Error uploading image", http.StatusInternalServerError)
			return
		}
		log.Printf("Image successfully uploaded to S3. URL: %s", imageURL)
	}

	err = c.Service.CreateComment(content, imageURL, int(userID), int(taskID))
	if err != nil {
		log.Printf("Error creating comment: %v", err)
		http.Error(w, "Failed to create comment", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("Comment created successfully"))
	log.Println("Comment created successfully")
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
func (c *CommentController) UpdateComment(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, _ := strconv.ParseUint(params["id"], 10, 32)

	var requestBody comment_dto.UpdateCommentRequestDTO

	err := json.NewDecoder(r.Body).Decode(&requestBody)
	if err != nil {
		http.Error(w, "Invalid data", http.StatusBadRequest)
		return
	}

	err = c.Service.UpdateComment(uint(id), requestBody.Content, requestBody.ImageURL)
	if err != nil {
		http.Error(w, "Failed to update comment", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

// DeleteComment deletes a comment by its ID
// @Summary Delete a comment by ID
// @Description Deletes an existing comment from the system by its ID
// @Tags comments
// @Param id path string true "Comment ID"
// @Success 204 "Comment deleted successfully"
// @Failure 500 {string} string "Internal Server Error"
// @Router /comment/{id} [delete]
func (c *CommentController) DeleteComment(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, _ := strconv.ParseUint(params["id"], 10, 32)

	err := c.Service.DeleteComment(uint(id))
	if err != nil {
		http.Error(w, "Failed to delete comment", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
