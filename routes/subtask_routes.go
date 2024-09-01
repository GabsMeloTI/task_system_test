package routes

import (
	"awesomeProject/controllers"
	"github.com/gorilla/mux"
)

// registra as rotas relacionadas a usuários
func SubtaskRoutes(r *mux.Router) {
	// retorna todos os usuários
	r.HandleFunc("/subtask", controllers.GetSubtasks).Methods("GET")

	// retorna usuário específico por id
	r.HandleFunc("/subtask/{id:[0-9]+}", controllers.GetSubtaskByID).Methods("GET")

	// criar um novo usuário
	r.HandleFunc("/subtask", controllers.CreateSubtask).Methods("POST")

	// atualiza usuário existente por id
	r.HandleFunc("/subtask/{id:[0-9]+}", controllers.UpdateSubtask).Methods("PUT")

	// deletar um usuário existente por id
	r.HandleFunc("/subtask/{id:[0-9]+}", controllers.DeleteSubtask).Methods("DELETE")
}
