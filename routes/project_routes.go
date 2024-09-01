package routes

import (
	"awesomeProject/controllers"
	"github.com/gorilla/mux"
	_ "net/http"
)

// registra as rotas relacionadas a usuários
func ProjectRoutes(r *mux.Router) {
	// retorna todos os projeto
	r.HandleFunc("/project", controllers.GetProjetos).Methods("GET")

	// retorna projeto específico por id
	r.HandleFunc("/project/{id:[0-9]+}", controllers.GetProjetoId).Methods("GET")

	// criar um novo projeto
	r.HandleFunc("/project", controllers.CreateProjeto).Methods("POST")

	// atualiza projeto existente por id
	r.HandleFunc("/project/{id:[0-9]+}", controllers.UpdateProjeto).Methods("PUT")

	// deletar um projeto existente por id
	r.HandleFunc("/project/{id:[0-9]+}", controllers.DeleteProjeto).Methods("DELETE")
}
