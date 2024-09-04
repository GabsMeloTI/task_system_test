package routes

import (
	"awesomeProject/controllers"
	"awesomeProject/service"
	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

func SectionRoutes(r *mux.Router, db *gorm.DB) {
	sectionService := service.NewSectionService(db)
	sectionController := controllers.SectionController{
		Service: sectionService,
	}

	r.HandleFunc("/section", sectionController.GetSections).Methods("GET")
	r.HandleFunc("/section/{id:[0-9]+}", sectionController.GetSectionByID).Methods("GET")
	r.HandleFunc("/section", sectionController.CreateSection).Methods("POST")
	r.HandleFunc("/section/{id:[0-9]+}", sectionController.UpdateSection).Methods("PUT")
	r.HandleFunc("/section/{id:[0-9]+}", sectionController.DeleteSection).Methods("DELETE")
}
