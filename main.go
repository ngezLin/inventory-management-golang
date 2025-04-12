package main

import (
	"inventory-management/config"
	"inventory-management/models"
	"inventory-management/routes"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	db := config.ConnectDatabase()

	db.AutoMigrate(&models.Product{}, &models.Inventory{}, &models.Order{})

	r := gin.Default()

	routes.SetupRoutes(r, db)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	r.Static("/uploads", "./uploads")
	r.Run(":" + port)
}
