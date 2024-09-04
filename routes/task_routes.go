package routes

import (
	"awesomeProject/controllers"
	"awesomeProject/service"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func TaskRoutes(e *echo.Echo, db *gorm.DB) {
	taskService := service.NewTaskService(db)
	taskController := &controllers.TaskController{
		Service: taskService,
	}

	e.GET("/task", taskController.GetTasks)
	e.GET("/task/:id", taskController.GetTaskByID)
	e.POST("/task", taskController.CreateTask)
	e.POST("/task/:id/labels", taskController.AssignLabelsToTask)
	e.PUT("/task/:id", taskController.UpdateTask)
	e.DELETE("/task/:id", taskController.DeleteTask)
}
