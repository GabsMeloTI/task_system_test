package routes

import (
	"awesomeProject/controllers"
	"github.com/gorilla/mux"
	_ "net/http"
)

// registra as rotas relacionadas a usuários
func UserRoutes(r *mux.Router) {
	// retorna todos os usuários
	r.HandleFunc("/user", controllers.GetUsuarios).Methods("GET")

	// retorna usuário específico por id
	r.HandleFunc("/user/{id:[0-9]+}", controllers.GetUsuarioId).Methods("GET")

	// criar um novo usuário
	r.HandleFunc("/user", controllers.CreateUsuario).Methods("POST")

	// atualiza usuário existente por id
	r.HandleFunc("/user/{id:[0-9]+}", controllers.UpdateUsuario).Methods("PUT")

	// deletar um usuário existente por id
	r.HandleFunc("/user/{id:[0-9]+}", controllers.DeleteUsuario).Methods("DELETE")

	r.HandleFunc("/register", controllers.RegisterUsuario).Methods("POST")

	r.HandleFunc("/login", controllers.Login).Methods("POST")
}
