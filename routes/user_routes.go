package routes

import (
	"awesomeProject/controllers"
	"awesomeProject/service"
	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

func UserRoutes(r *mux.Router, db *gorm.DB) {
	userService := service.NewUserService(db)
	userController := &controllers.UserController{
		Service: userService,
	}

	r.HandleFunc("/user", userController.GetUsers).Methods("GET")
	r.HandleFunc("/user/{id:[0-9]+}", userController.GetUserByID).Methods("GET")
	r.HandleFunc("/user", userController.CreateUser).Methods("POST")
	r.HandleFunc("/user/{id:[0-9]+}", userController.UpdateUser).Methods("PUT")
	r.HandleFunc("/user/{id:[0-9]+}/image", userController.UpdateUserImage).Methods("PATCH")
	r.HandleFunc("/user/{id:[0-9]+}", userController.DeleteUser).Methods("DELETE")
	r.HandleFunc("/register", userController.RegisterUser).Methods("POST")
	r.HandleFunc("/login", userController.Login).Methods("POST")
}
