package main

import (
	"awesomeProject/configs"
	"awesomeProject/db"
	_ "awesomeProject/docs"
	"awesomeProject/routes"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	httpSwagger "github.com/swaggo/http-swagger"
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

	corsHandler := handlers.CORS(
		handlers.AllowedOrigins([]string{"*"}),
		handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE"}),
		handlers.AllowedHeaders([]string{"Content-Type", "Authorization"}),
	)

	r := mux.NewRouter()
	routes.RegisterRoutes(r, db.DB)

	r.PathPrefix("/swagger/").Handler(httpSwagger.WrapHandler)
	log.Println("Iniciando o servidor na porta 8000...")
	log.Fatal(http.ListenAndServe(":8000", corsHandler(r)))
}
