package routes

import (
	"awesomeProject/controllers"
	"awesomeProject/service"
	"gorm.io/gorm"
)

func UserRoutes(e *echo.Echo, db *gorm.DB) {
	userService := service.NewUserService(db)
	userController := &controllers.UserController{
		Service: userService,
	}

	e.GET("/user", userController.GetUsers)
	e.GET("/user/:id", userController.GetUserByID)
	e.POST("/user", userController.CreateUser)
	e.PUT("/user/:id", userController.UpdateUser)
	e.PATCH("/user/:id/image", userController.UpdateUserImage)
	e.DELETE("/user/:id", userController.DeleteUser)
	e.POST("/register", userController.RegisterUser)
	e.POST("/login", userController.Login)
}
