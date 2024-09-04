package routes

import (
	"awesomeProject/controllers"
	"awesomeProject/service"
	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

func TaskRoutes(r *mux.Router, db *gorm.DB) {
	taskService := service.NewTaskService(db)
	taskController := &controllers.TaskController{
		Service: (*service.TaskService)(taskService),
	}

	r.HandleFunc("/task", taskController.GetTasks).Methods("GET")
	r.HandleFunc("/task/{id:[0-9]+}", taskController.GetTaskByID).Methods("GET")
	r.HandleFunc("/task", taskController.CreateTask).Methods("POST")
	r.HandleFunc("/task/{id:[0-9]+}/label", taskController.AssignLabelsToTask).Methods("POST")
	r.HandleFunc("/task/{id:[0-9]+}", taskController.UpdateTask).Methods("PUT")
	r.HandleFunc("/task/{id:[0-9]+}", taskController.DeleteTask).Methods("DELETE")
}
