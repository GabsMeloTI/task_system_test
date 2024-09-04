package routes

import (
	"awesomeProject/controllers"
	"awesomeProject/service"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func SectionRoutes(e *echo.Echo, db *gorm.DB) {
	sectionService := service.NewSectionService(db)
	sectionController := controllers.NewSectionController(sectionService)

	e.GET("/section", sectionController.GetSections)
	e.GET("/section/:id", sectionController.GetSectionByID)
	e.POST("/section", sectionController.CreateSection)
	e.PUT("/section/:id", sectionController.UpdateSection)
	e.DELETE("/section/:id", sectionController.DeleteSection)
}
