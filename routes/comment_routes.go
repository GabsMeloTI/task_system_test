package routes

import (
	"awesomeProject/controllers"
	"github.com/gorilla/mux"
)

// registra as rotas relacionadas a usuários
func CommentRoutes(r *mux.Router) {
	// retorna todos os usuários
	r.HandleFunc("/comment", controllers.GetComentario).Methods("GET")

	// retorna usuário específico por id
	r.HandleFunc("/comment/{id:[0-9]+}", controllers.GetComentariosId).Methods("GET")

	// criar um novo usuário
	r.HandleFunc("/comment", controllers.CreateComentario).Methods("POST")

	// atualiza usuário existente por id
	r.HandleFunc("/comment/{id:[0-9]+}", controllers.UpdateComentario).Methods("PUT")

	// deletar um usuário existente por id
	r.HandleFunc("/comment/{id:[0-9]+}", controllers.DeleteComentario).Methods("DELETE")
}
