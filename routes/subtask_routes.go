package routes

import (
	"awesomeProject/controllers"
	"awesomeProject/service"
	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

func SubtaskRoutes(r *mux.Router, db *gorm.DB) {
	subtaskService := service.NewSubtaskService(db)
	subtaskController := &controllers.SubtaskController{
		Service: (*service.SubtaskService)(subtaskService),
	}

	r.HandleFunc("/subtask", subtaskController.GetSubtasks).Methods("GET")
	r.HandleFunc("/subtask/{id:[0-9]+}", subtaskController.GetSubtaskByID).Methods("GET")
	r.HandleFunc("/subtask", subtaskController.CreateSubtask).Methods("POST")
	r.HandleFunc("/subtask/{id:[0-9]+}", subtaskController.UpdateSubtask).Methods("PUT")
	r.HandleFunc("/subtask/{id:[0-9]+}", subtaskController.DeleteSubtask).Methods("DELETE")
}
