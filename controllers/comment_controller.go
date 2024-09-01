package controllers

import (
	"awesomeProject/db"
	"awesomeProject/dto/comment_dto"
	"awesomeProject/dto/task_dto"
	"awesomeProject/dto/user_dto"
	"awesomeProject/models"
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
)

func GetComments(w http.ResponseWriter, r *http.Request) {
	var comments []models.Comment
	var commentsDTO []comment_dto.CommentListingDTO

	if err := db.DB.Preload("User").Preload("Task").Find(&comments).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	for _, comment := range comments {
		commentDTO := comment_dto.CommentListingDTO{
			ID:          comment.ID,
			Content:     comment.Content,
			PublishedAt: comment.PublishedAt,
			ImageURL:    comment.Image,
			User: user_dto.UserBasicDTO{
				ID:   comment.User.ID,
				Name: comment.User.Name,
			},
			Task: task_dto.TaskBasicDTO{
				ID:    comment.Task.ID,
				Title: comment.Task.Title,
			},
		}
		commentsDTO = append(commentsDTO, commentDTO)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(commentsDTO)
}

func GetCommentByID(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var comment models.Comment

	if err := db.DB.Preload("User").Preload("Task").First(&comment, params["id"]).Error; err != nil {
		http.Error(w, "Comment not found", http.StatusNotFound)
		return
	}

	commentDTO := comment_dto.CommentListingDTO{
		ID:          comment.ID,
		Content:     comment.Content,
		PublishedAt: comment.PublishedAt,
		ImageURL:    comment.Image,
		User: user_dto.UserBasicDTO{
			ID:   comment.User.ID,
			Name: comment.User.Name,
		},
		Task: task_dto.TaskBasicDTO{
			ID:    comment.Task.ID,
			Title: comment.Task.Title,
		},
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(commentDTO)
}

func CreateComment(w http.ResponseWriter, r *http.Request) {
	var comment models.Comment

	if err := json.NewDecoder(r.Body).Decode(&comment); err != nil {
		http.Error(w, "Invalid data", http.StatusBadRequest)
		return
	}

	if err := db.DB.Create(&comment).Error; err != nil {
		http.Error(w, "Failed to create comment", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func UpdateComment(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var comment models.Comment

	if err := db.DB.First(&comment, params["id"]).Error; err != nil {
		http.Error(w, "Comment not found", http.StatusNotFound)
		return
	}

	if err := json.NewDecoder(r.Body).Decode(&comment); err != nil {
		http.Error(w, "Invalid data", http.StatusBadRequest)
		return
	}

	if err := db.DB.Save(&comment).Error; err != nil {
		http.Error(w, "Failed to update comment", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func DeleteComment(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	if err := db.DB.Delete(&models.Comment{}, params["id"]).Error; err != nil {
		http.Error(w, "Failed to delete comment", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
