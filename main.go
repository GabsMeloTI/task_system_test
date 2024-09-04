package main

import (
	"awesomeProject/configs"
	"awesomeProject/db"
	_ "awesomeProject/docs"
	"awesomeProject/routes"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"log"
	"net/http"
)

func main() {
	configs.InitS3Client()
	db.InitGorm()

	if db.DB == nil {
		log.Fatal("Conexão com o banco de dados falhou. Variável DB é nil.")
	}

	db.MigrateTables()

	e := echo.New()

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{echo.GET, echo.POST, echo.PUT, echo.DELETE},
		AllowHeaders: []string{"Content-Type", "Authorization"},
	}))
	routes.RegisterRoutes(e, db.DB)

	e.GET("/swagger/*", echoSwagger.WrapHandler)
	log.Println("Iniciando o servidor na porta 8000...")
	log.Fatal(http.ListenAndServe(":8000", corsHandler(r)))
}
