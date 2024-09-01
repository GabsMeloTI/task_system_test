package routes

import (
	"awesomeProject/controllers"
	"github.com/gorilla/mux"
)

// registra as rotas relacionadas a usuários
func SectionRoutes(r *mux.Router) {
	// retorna todos os usuários
	r.HandleFunc("/section", controllers.GetSecao).Methods("GET")

	// retorna usuário específico por id
	r.HandleFunc("/section/{id:[0-9]+}", controllers.GetSecaoId).Methods("GET")

	// criar um novo usuário
	r.HandleFunc("/section", controllers.CreateSecao).Methods("POST")

	// atualiza usuário existente por id
	r.HandleFunc("/section/{id:[0-9]+}", controllers.UpdateSecao).Methods("PUT")

	// deletar um usuário existente por id
	r.HandleFunc("/section/{id:[0-9]+}", controllers.DeleteSecao).Methods("DELETE")
}
