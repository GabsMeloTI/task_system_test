package routes

import (
	"awesomeProject/controllers"
	"awesomeProject/service"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func SubtaskRoutes(e *echo.Echo, db *gorm.DB) {
	subtaskService := service.NewSubtaskService(db)
	subtaskController := &controllers.SubtaskController{
		Service: (*service.SubtaskService)(subtaskService),
	}

	e.GET("/subtask", subtaskController.GetSubtasks)
	e.GET("/subtask/:id", subtaskController.GetSubtaskByID)
	e.POST("/subtask", subtaskController.CreateSubtask)
	e.PUT("/subtask/:id", subtaskController.UpdateSubtask)
	e.DELETE("/subtask/:id", subtaskController.DeleteSubtask)
}
