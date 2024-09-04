package routes

import (
	"awesomeProject/controllers"
	"awesomeProject/service"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func LabelRoutes(e *echo.Echo, db *gorm.DB) {
	labelService := service.NewLabelService(db)
	labelController := &controllers.LabelController{
		Service: labelService,
	}

	e.GET("/label", labelController.GetLabels)
	e.GET("/label/:id", labelController.GetLabelByID)
	e.POST("/label", labelController.CreateLabel)
	e.PUT("/label/:id", labelController.UpdateLabel)
	e.DELETE("/label/:id", labelController.DeleteLabel)
}
