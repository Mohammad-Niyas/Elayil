package controllers

import (
	"net/http"
	"kerala-food-finder/config"
	"kerala-food-finder/models"
	"github.com/gin-gonic/gin"
)

func GetAllDishes(c *gin.Context) {
	var dishes []models.Dish
	city := c.Query("city")
	category := c.Query("category")

	query := config.DB.Preload("Restaurant")

	if city != "" {
		query = query.Joins(
			"JOIN restaurants ON restaurants.id = dishes.restaurant_id",
		).Where("restaurants.city = ?", city)
	}

	if category != "" {
		query = query.Where(
			"dishes.category = ?", category,
		)
	}

	query.Find(&dishes)

	c.JSON(http.StatusOK, gin.H{
		"data": dishes,
	})
}

func GetDish(c *gin.Context) {
	var dish models.Dish
	id := c.Param("id")

	result := config.DB.
		Preload("Restaurant").
		Preload("Reviews").
		First(&dish, id)

	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Dish not found",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": dish,
	})
}

func SaveDish(c *gin.Context) {
	id := c.Param("id")

	var dish models.Dish
	result := config.DB.First(&dish, id)
	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Dish not found",
		})
		return
	}

	save := models.Save{
		DishID: dish.ID,
	}
	config.DB.Create(&save)

	config.DB.Model(&dish).Update(
		"saves", dish.Saves+1,
	)

	c.JSON(http.StatusOK, gin.H{
		"message": "Dish saved successfully!",
	})
}


func UnsaveDish(c *gin.Context) {
	id := c.Param("id")

	config.DB.Where(
		"dish_id = ?", id,
	).Delete(&models.Save{})

	var dish models.Dish
	config.DB.First(&dish, id)
	if dish.Saves > 0 {
		config.DB.Model(&dish).Update(
			"saves", dish.Saves-1,
		)
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Dish unsaved successfully!",
	})
}
