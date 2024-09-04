package routes

import (
	"awesomeProject/controllers"
	"awesomeProject/service"
	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

func CommentRoutes(r *mux.Router, db *gorm.DB) {
	commentService := service.NewCommentService(db)
	commentController := &controllers.CommentController{
		Service: commentService,
	}

	r.HandleFunc("/comment", commentController.GetComments).Methods("GET")
	r.HandleFunc("/comment/{id:[0-9]+}", commentController.GetCommentByID).Methods("GET")
	r.HandleFunc("/comment", commentController.CreateComment).Methods("POST")
	r.HandleFunc("/comment/{id:[0-9]+}", commentController.UpdateComment).Methods("PUT")
	r.HandleFunc("/comment/{id:[0-9]+}", commentController.DeleteComment).Methods("DELETE")
}
