package data

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var Db *gorm.DB
var DB *sql.DB

func ConnectDB() error {
	fmt.Println("any connectdb ")
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	DB_USER := os.Getenv("DB_USER")
	DB_PASSWORD := os.Getenv("DB_PASS")
	DB_HOST := os.Getenv("DB_HOST")
	DB_PORT := os.Getenv("DB_PORT")
	DB_NAME := os.Getenv("DB_NAME")

	fmt.Printf("DB_USER: %s,DB_PASSWORD: %s ,DB_HOST: %s, DB_PORT: %s, DB_NAME: %s\n", DB_USER, DB_PASSWORD, DB_HOST, DB_PORT, DB_NAME)

	connectionString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		DB_USER, DB_PASSWORD, DB_HOST, DB_PORT, DB_NAME)

	Db, err = gorm.Open(mysql.Open(connectionString), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to the database ----->: ", err)
	}

	DB, err = sql.Open("mysql", connectionString)
	if err != nil {
		log.Fatal("Failed to initialize sql.DB ----->: ", err)
	}

	// Db.AutoMigrate(&models.User{})
	return nil
}
