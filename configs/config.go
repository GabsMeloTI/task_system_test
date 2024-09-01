package configs

import (
	"log"

	_ "github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/spf13/viper"
)

var cfg *config

type config struct {
	API APIConfig
	DB  DBConfig
}

type APIConfig struct {
	Port string
}

type DBConfig struct {
	Host     string
	Port     string
	User     string
	Pass     string
	Database string
}

func init() {
	viper.SetDefault("api.port", "8081")
	viper.SetDefault("database.host", "localhost")
	viper.SetDefault("database.port", "5432")
	viper.SetDefault("database.user", "postgres")
	viper.SetDefault("database.pass", "091104")
	viper.SetDefault("database.database", "mydb")

	// Permite que variáveis de ambiente sobrescrevam os valores do arquivo de configuração
	viper.AutomaticEnv()
}

func Load() error {
	viper.SetConfigName("config")
	viper.SetConfigType("toml")
	viper.AddConfigPath(".")

	err := viper.ReadInConfig()
	if err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); !ok {
			return err
		}
		log.Println("Arquivo de configuração não encontrado, usando valores padrão.")
	}

	cfg = new(config)

	cfg.API = APIConfig{
		Port: viper.GetString("api.port"),
	}

	cfg.DB = DBConfig{
		Host:     viper.GetString("database.host"),
		Port:     viper.GetString("database.port"),
		User:     viper.GetString("database.user"),
		Pass:     viper.GetString("database.pass"),
		Database: viper.GetString("database.database"),
	}

	return nil
}

func (c *config) GetDB() DBConfig {
	return c.DB
}

func (c *config) GetServerPort() string {
	return c.API.Port
}

func GetDB() DBConfig {
	return cfg.GetDB()
}

func GetServerPort() string {
	return cfg.GetServerPort()
}
