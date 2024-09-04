package routes

import (
	"awesomeProject/controllers"
	"awesomeProject/service"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func ProjectRoutes(e *echo.Echo, db *gorm.DB) {
	projectService := service.NewProjectService(db)
	projectController := controllers.NewProjectController(projectService)

	e.GET("/project", projectController.GetProjects)
	e.GET("/project/:id", projectController.GetProjectByID)
	e.POST("/project", projectController.CreateProject)
	e.PUT("/project/:id", projectController.UpdateProject)
	e.DELETE("/project/:id", projectController.DeleteProject)
}
