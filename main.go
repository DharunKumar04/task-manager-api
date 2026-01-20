package main

import (
	"fmt"
	"log"

	"github.com/DharunKumar04/task-manager-api/config"
	"gorm.io/gorm"
)

var DB *gorm.DB

func main() {
	var err error
	DB, err = config.ConnectPSQLDB()
	if err != nil {
		log.Fatal("Database connection failed:", err)
	}
	fmt.Println("✅ DB Connection is Successful")
}
