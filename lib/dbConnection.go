package lib

import (
	"os"
	"log"
	"gorm.io/gorm"
	"gorm.io/driver/postgres"
	"github.com/golang-backend/models"
)

var DB *gorm.DB

func ConnectDatabase() {

	dbString := os.Getenv("POSTGRESDB")
  dataBase, err := gorm.Open(postgres.Open(dbString), &gorm.Config{})

	if err != nil {
		log.Printf("Failed to connect to database!")
		panic(err)
	}

	dataBase.AutoMigrate(&models.User{}) // register Post model

	DB = dataBase
}