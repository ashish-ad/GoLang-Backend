package models

import (
	"log"
	"gorm.io/gorm"
	"gorm.io/driver/mysql"
)

var DB *gorm.DB

func ConnectDatabase() {
	// dataBaseConnection := "host=localhost user=root dbname=go_lang_test port=3306 sslmode=disable"
	  // refer https://github.com/go-sql-driver/mysql#dsn-data-source-name for details
  dsn := "root:nirXyj-capva6-byvbyb@tcp(127.0.0.1:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local"
  dataBase, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Printf("Failed to connect to database!")
		panic(err)
	}

	dataBase.AutoMigrate(&Post{}) // register Post model

	DB = dataBase
}