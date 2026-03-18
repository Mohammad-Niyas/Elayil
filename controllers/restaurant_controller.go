package controllers

import (
	"net/http"
	"kerala-food-finder/config"
	"kerala-food-finder/models"
	"github.com/gin-gonic/gin"
)

func GetAllRestaurants(c *gin.Context) {
	var restaurants []models.Restaurant
	config.DB.Find(&restaurants)
	c.JSON(http.StatusOK, gin.H{
		"data": restaurants,
	})
}


func GetRestaurant(c *gin.Context) {
	var restaurant models.Restaurant
	id := c.Param("id")

	result := config.DB.First(&restaurant, id)
	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Restaurant not found",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": restaurant,
	})
}


func GetRestaurantDishes(c *gin.Context) {
	var dishes []models.Dish
	id := c.Param("id")

	config.DB.Where(
		"restaurant_id = ?", id,
	).Find(&dishes)

	c.JSON(http.StatusOK, gin.H{
		"data": dishes,
	})
}


func GetRestaurantReels(c *gin.Context) {
	var reels []models.Reel
	id := c.Param("id")

	config.DB.Where(
		"restaurant_id = ?", id,
	).Find(&reels)

	c.JSON(http.StatusOK, gin.H{
		"data": reels,
	})
}