package routes

import (
	_ "awesomeProject/controllers"
	"github.com/gorilla/mux"
)

// registra todas as rotas da aplicação
func RegisterRoutes(r *mux.Router) {
	UserRoutes(r)
	ProjectRoutes(r)
	SectionRoutes(r)
	TaskRoutes(r)
	SubtaskRoutes(r)
	LabelRoutes(r)
	CommentRoutes(r)
}
