package routes

import (
	"awesomeProject/controllers"
	"github.com/gorilla/mux"
	_ "net/http"
)

// registra as rotas relacionadas a usuários
func UserRoutes(r *mux.Router) {
	// retorna todos os usuários
	r.HandleFunc("/user", controllers.GetUsers).Methods("GET")

	// retorna usuário específico por id
	r.HandleFunc("/user/{id:[0-9]+}", controllers.GetUserByID).Methods("GET")

	// criar um novo usuário
	r.HandleFunc("/user", controllers.CreateUser).Methods("POST")

	// atualiza usuário existente por id
	r.HandleFunc("/user/{id:[0-9]+}", controllers.UpdateUser).Methods("PUT")

	// deletar um usuário existente por id
	r.HandleFunc("/user/{id:[0-9]+}", controllers.DeleteUser).Methods("DELETE")

	r.HandleFunc("/register", controllers.RegisterUser).Methods("POST")

	r.HandleFunc("/login", controllers.Login).Methods("POST")
}
