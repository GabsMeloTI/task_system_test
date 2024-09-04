package routes

import (
	"awesomeProject/controllers"
	"awesomeProject/service"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func CommentRoutes(e *echo.Echo, db *gorm.DB) {
	commentService := service.NewCommentService(db)
	commentController := &controllers.CommentController{
		Service: commentService,
	}

	e.GET("/comment", commentController.GetComments)
	e.GET("/comment/:id", commentController.GetCommentByID)
	e.POST("/comment", commentController.CreateComment)
	e.PUT("/comment/:id", commentController.UpdateComment)
	e.DELETE("/comment/:id", commentController.DeleteComment)
}
