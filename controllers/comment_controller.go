package controllers

import (
	"awesomeProject/service"
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

type CommentController struct {
	Service *service.CommentService
}

func (c *CommentController) GetComments(w http.ResponseWriter, r *http.Request) {
	commentsDTO, err := c.Service.GetAllComments()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(commentsDTO)
}

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

func (c *CommentController) CreateComment(w http.ResponseWriter, r *http.Request) {
	err := r.ParseMultipartForm(10 << 20)
	if err != nil {
		http.Error(w, "Unable to parse form", http.StatusBadRequest)
		return
	}

	content := r.FormValue("content")
	file, fileHeader, err := r.FormFile("image")
	if err != nil {
		if err == http.ErrMissingFile {
			file = nil
			fileHeader = nil
		} else {
			http.Error(w, "Error retrieving the file", http.StatusBadRequest)
			return
		}
	}

	var imageURL string
	if file != nil && fileHeader != nil {
		imageURL, err = c.Service.UploadToS3(file, fileHeader.Filename)
		if err != nil {
			http.Error(w, "Error uploading image", http.StatusInternalServerError)
			return
		}
	}

	err = c.Service.CreateComment(content, imageURL)
	if err != nil {
		http.Error(w, "Failed to create comment", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func (c *CommentController) UpdateComment(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, _ := strconv.ParseUint(params["id"], 10, 32)

	var requestBody struct {
		Content  string `json:"content"`
		ImageURL string `json:"image_url"`
	}

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
