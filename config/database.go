package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func getDbConnectionString() string {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	DB_HOST := os.Getenv("DB_HOST")
	DB_PORT := os.Getenv("DB_PORT")
	DB_USER := os.Getenv("DB_USER")
	DB_PASSWORD := os.Getenv("DB_PASSWORD")
	DB_NAME := os.Getenv("DB_NAME")

	var DATABASE_CONNECTION_STRING string = "postgresql://" + DB_USER + ":" + DB_PASSWORD + "@" + DB_HOST + ":" + DB_PORT + "/" + DB_NAME
	return DATABASE_CONNECTION_STRING
}

func ConnectPSQLDB() (*gorm.DB, error) {
	db, err := gorm.Open(postgres.Open(getDbConnectionString()), &gorm.Config{})
	return db, err
}
