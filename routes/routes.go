package routes

import (
	"github.com/gorilla/mux"
	"gorm.io/gorm"
	"log"
)

func RegisterRoutes(r *mux.Router, db *gorm.DB) {
	log.Println("Registrando rotas")
	UserRoutes(r, db)
	ProjectRoutes(r, db)
	SectionRoutes(r, db)
	TaskRoutes(r, db)
	SubtaskRoutes(r, db)
	LabelRoutes(r, db)
	CommentRoutes(r, db)
	log.Println("Rotas registradas com sucesso")
}
