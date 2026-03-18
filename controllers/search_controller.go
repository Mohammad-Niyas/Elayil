package controllers

import (
	"net/http"
	"kerala-food-finder/config"
	"kerala-food-finder/models"
	"github.com/gin-gonic/gin"
)

func Search(c *gin.Context) {
	query := c.Query("q")

	if query == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Search query is required",
		})
		return
	}

	var dishes []models.Dish

	searchTerm := "%" + query + "%"

	config.DB.
		Preload("Restaurant").
		Joins("JOIN restaurants ON restaurants.id = dishes.restaurant_id").
		Where(
			"dishes.name ILIKE ? OR restaurants.name ILIKE ? OR restaurants.city ILIKE ? OR restaurants.area ILIKE ?",
			searchTerm, searchTerm, searchTerm, searchTerm,
		).
		Find(&dishes)

	c.JSON(http.StatusOK, gin.H{
		"data":  dishes,
		"count": len(dishes),
	})
}