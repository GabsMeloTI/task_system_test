package routes

import (
	"awesomeProject/controllers"
	"awesomeProject/service"
	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

func ProjectRoutes(r *mux.Router, db *gorm.DB) {
	projectService := service.NewProjectService(db)
	projectController := controllers.ProjectController{
		Service: projectService,
	}

	r.HandleFunc("/project", projectController.GetProjects).Methods("GET")
	r.HandleFunc("/project/{id:[0-9]+}", projectController.GetProjectByID).Methods("GET")
	r.HandleFunc("/project", projectController.CreateProject).Methods("POST")
	r.HandleFunc("/project/{id:[0-9]+}", projectController.UpdateProject).Methods("PUT")
	r.HandleFunc("/project/{id:[0-9]+}", projectController.DeleteProject).Methods("DELETE")
}
