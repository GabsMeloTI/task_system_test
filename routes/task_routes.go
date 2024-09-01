package routes

import (
	"awesomeProject/controllers"
	"github.com/gorilla/mux"
)

// registra as rotas relacionadas a usuários
func TaskRoutes(r *mux.Router) {
	// retorna todos os usuários
	r.HandleFunc("/task", controllers.GetTasks).Methods("GET")

	// retorna usuário específico por id
	r.HandleFunc("/task/{id:[0-9]+}", controllers.GetTaskByID).Methods("GET")

	// criar um novo usuário
	r.HandleFunc("/task", controllers.CreateTask).Methods("POST")

	// criar um novo usuário
	r.HandleFunc("/task/{id:[0-9]+}/label", controllers.AssignLabelsToTask).Methods("POST")

	// atualiza usuário existente por id
	r.HandleFunc("/task/{id:[0-9]+}", controllers.UpdateTask).Methods("PUT")

	// deletar um usuário existente por id
	r.HandleFunc("/task/{id:[0-9]+}", controllers.DeleteTask).Methods("DELETE")
}
