package controllers

import (
	"net/http"
	"kerala-food-finder/config"
	"kerala-food-finder/models"
	"github.com/gin-gonic/gin"
)

// GET /api/trending?city=Kochi
func GetTrending(c *gin.Context) {
	var dishes []models.Dish
	city := c.Query("city")

	query := config.DB.
		Preload("Restaurant").
		Joins("JOIN restaurants ON restaurants.id = dishes.restaurant_id").
		Order("dishes.saves DESC").
		Limit(10)

	if city != "" {
		query = query.Where(
			"restaurants.city = ?", city,
		)
	}

	query.Find(&dishes)

	c.JSON(http.StatusOK, gin.H{
		"data": dishes,
	})
}