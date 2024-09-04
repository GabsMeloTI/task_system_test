package db

import (
	"awesomeProject/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

var DB *gorm.DB

func OpenConnection() (*gorm.DB, error) {
	dsn := "host=go_db port=5432 dbname='PostgreSQL 16' user=postgres password=12345678 connect_timeout=10 sslmode=prefer"
	log.Printf("DSN: %s", dsn)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Printf("Erro ao abrir conexão com o banco de dados: %v", err)
		return nil, err
	}
	log.Println("Conexão com o banco de dados aberta com sucesso")
	return db, nil
}

func InitGorm() {
	log.Println("Iniciando a conexão com o banco de dados...")
	var err error
	DB, err = OpenConnection()
	if err != nil {
		log.Fatalf("Erro ao conectar ao banco de dados: %v", err)
	}
	if DB == nil {
		log.Fatalf("A instância do banco de dados é nil após a inicialização")
	}
	log.Println("Conexão com o banco de dados estabelecida com sucesso")
}

func MigrateTables() {
	log.Println("Migrando as tabelas...")
	if DB != nil {
		err := DB.AutoMigrate(&models.User{}, &models.Project{}, &models.Section{}, &models.Task{}, &models.Subtask{}, &models.Label{}, &models.Comment{})
		if err != nil {
			log.Fatalf("Erro ao migrar as tabelas: %v", err)
		}
	} else {
		log.Fatal("Banco de dados não inicializado")
	}
}
