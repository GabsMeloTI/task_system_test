package routes

import (
	"awesomeProject/controllers"
	"awesomeProject/service"
	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

func LabelRoutes(r *mux.Router, db *gorm.DB) {
	labelService := service.NewLabelService(db)
	labelController := &controllers.LabelController{
		Service: labelService,
	}

	r.HandleFunc("/label", labelController.GetLabels).Methods("GET")
	r.HandleFunc("/label/{id:[0-9]+}", labelController.GetLabelByID).Methods("GET")
	r.HandleFunc("/label", labelController.CreateLabel).Methods("POST")
	r.HandleFunc("/label/{id:[0-9]+}", labelController.UpdateLabel).Methods("PUT")
	r.HandleFunc("/label/{id:[0-9]+}", labelController.DeleteLabel).Methods("DELETE")
}
