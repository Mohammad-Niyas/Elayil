package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"kerala-food-finder/config"
	"kerala-food-finder/models"
	"kerala-food-finder/routes"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	config.ConnectDatabase()

	// AutoMigrate here!
	config.DB.AutoMigrate(
		&models.Restaurant{},
		&models.Dish{},
		&models.Reel{},
		&models.Review{},
		&models.Save{},
	)

	router := gin.Default()
	routes.SetupRoutes(router)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Println("Server starting on port:", port)
	router.Run(":" + port)
}