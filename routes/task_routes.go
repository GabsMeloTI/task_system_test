package routes

import (
	"awesomeProject/controllers"
	"github.com/gorilla/mux"
)

// registra as rotas relacionadas a usuários
func TaskRoutes(r *mux.Router) {
	// retorna todos os usuários
	r.HandleFunc("/task", controllers.GetTarefa).Methods("GET")

	// retorna usuário específico por id
	r.HandleFunc("/task/{id:[0-9]+}", controllers.GetTarefasId).Methods("GET")

	// criar um novo usuário
	r.HandleFunc("/task", controllers.CreateTarefa).Methods("POST")

	// criar um novo usuário
	r.HandleFunc("/task/{id:[0-9]+}/label", controllers.AddEtiquetaToTarefa).Methods("POST")

	// atualiza usuário existente por id
	r.HandleFunc("/task/{id:[0-9]+}", controllers.UpdateTarefa).Methods("PUT")

	// deletar um usuário existente por id
	r.HandleFunc("/task/{id:[0-9]+}", controllers.DeleteTarefa).Methods("DELETE")
}
