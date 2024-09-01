package db

import (
	"awesomeProject/configs"
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
)

var DB *gorm.DB

func OpenConnection() (*gorm.DB, error) {
	conf := configs.GetDB()

	sc := "host=%s port=%s user=%s password=%s dbname=%s sslmode=disable"
	dsn := fmt.Sprintf(sc, conf.Host, conf.Port, conf.User, conf.Pass, conf.Database)

	db, err := gorm.Open("postgres", dsn)
	if err != nil {
		return nil, err
	}

	return db, nil
}

func InitGorm() {
	var err error
	DB, err = OpenConnection()
	if err != nil {
		panic("failed to connect database")
	}
}
