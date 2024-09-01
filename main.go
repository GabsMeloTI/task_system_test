package main

import (
	"awesomeProject/configs"
	"awesomeProject/db"
	"awesomeProject/models"
	"awesomeProject/routes"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	if err := configs.Load(); err != nil {
		log.Fatalf("Erro ao carregar configurações: %v", err)
	}

	db.InitGorm()

	if err := db.DB.AutoMigrate(&models.Usuario{}, &models.Projeto{}, &models.Secao{},
		&models.Tarefa{}, &models.Subtarefa{}, &models.Comentario{}, &models.Etiqueta{}).Error; err != nil {
		log.Fatalf("Erro ao migrar as tabelas: %v", err)
	}

	r := mux.NewRouter()
	routes.RegisterRoutes(r)

	log.Println("Iniciando o servidor na porta 8080...")
	log.Fatal(http.ListenAndServe(":8080", r))
}
