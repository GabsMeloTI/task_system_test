package routes

import (
	"awesomeProject/controllers"
	"github.com/gorilla/mux"
)

// registra as rotas relacionadas a usuários
func LabelRoutes(r *mux.Router) {
	// retorna todos os usuários
	r.HandleFunc("/label", controllers.GetLabels).Methods("GET")

	// retorna usuário específico por id
	r.HandleFunc("/label/{id:[0-9]+}", controllers.GetLabelByID).Methods("GET")

	// criar um novo usuário
	r.HandleFunc("/label/{id:[0-9]+/etiquetas", controllers.CreateLabel).Methods("POST")

	// atualiza usuário existente por id
	r.HandleFunc("/label/{id:[0-9]+}", controllers.UpdateLabel).Methods("PUT")

	// deletar um usuário existente por id
	r.HandleFunc("/label/{id:[0-9]+}", controllers.DeleteLabel).Methods("DELETE")
}
