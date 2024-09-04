package routes

import (
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
	"log"
)

func RegisterRoutes(e *echo.Echo, db *gorm.DB) {
	log.Println("Registrando rotas")
	UserRoutes(e, db)
	ProjectRoutes(e, db)
	SectionRoutes(e, db)
	TaskRoutes(e, db)
	SubtaskRoutes(e, db)
	LabelRoutes(e, db)
	CommentRoutes(e, db)
	log.Println("Rotas registradas com sucesso")
}
