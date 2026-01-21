package main

import (
	"fmt"
	"log"

	"github.com/DharunKumar04/task-manager-api/config"
	"github.com/DharunKumar04/task-manager-api/handlers"
	"github.com/DharunKumar04/task-manager-api/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	db, err := config.ConnectPSQLDB()
	if err != nil {
		log.Fatal("Database connection failed:", err)
	}
	fmt.Println("✅ DB Connection is Successful")

	h := handlers.NewHandler(db)
	if h == nil {
		log.Fatal("Handler initialization failed")
	}
	fmt.Println("Handler initialized with database")

	routes.SetupRoutes(router, h)

	fmt.Println("Gin Server is Starting on http://0.0.0.0:8080")
	err = router.Run("0.0.0.0:8080")
	if err != nil {
		log.Fatal("Server failed:", err)
	}
}
