package routes

import (
	"awesomeProject/controllers"
	"github.com/gorilla/mux"
	_ "net/http"
)

// registra as rotas relacionadas a usuários
func ProjectRoutes(r *mux.Router) {
	// retorna todos os projeto
	r.HandleFunc("/project", controllers.GetProjects).Methods("GET")

	// retorna projeto específico por id
	r.HandleFunc("/project/{id:[0-9]+}", controllers.GetProjectByID).Methods("GET")

	// criar um novo projeto
	r.HandleFunc("/project", controllers.CreateProject).Methods("POST")

	// atualiza projeto existente por id
	r.HandleFunc("/project/{id:[0-9]+}", controllers.UpdateProject).Methods("PUT")

	// deletar um projeto existente por id
	r.HandleFunc("/project/{id:[0-9]+}", controllers.DeleteProject).Methods("DELETE")
}
