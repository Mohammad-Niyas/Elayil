package controllers

import (
	"net/http"
	"kerala-food-finder/config"
	"kerala-food-finder/models"
	"github.com/gin-gonic/gin"
)

func GetReviews(c *gin.Context) {
	var reviews []models.Review
	dishID := c.Param("id")

	config.DB.Where(
		"dish_id = ?", dishID,
	).Find(&reviews)

	c.JSON(http.StatusOK, gin.H{
		"data": reviews,
	})
}

func CreateReview(c *gin.Context) {
	dishID := c.Param("id")

	var dish models.Dish
	result := config.DB.First(&dish, dishID)
	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Dish not found",
		})
		return
	}

	var input struct {
		UserName string `json:"user_name"`
		Rating   int    `json:"rating"`
		Taste    int    `json:"taste"`
		Value    int    `json:"value"`
		Ambience int    `json:"ambience"`
		Comment  string `json:"comment"`
		Visited  bool   `json:"visited"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid input",
		})
		return
	}

	if input.Rating < 1 || input.Rating > 5 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Rating must be between 1 and 5",
		})
		return
	}

	review := models.Review{
		DishID:   dish.ID,
		UserName: input.UserName,
		Rating:   input.Rating,
		Taste:    input.Taste,
		Value:    input.Value,
		Ambience: input.Ambience,
		Comment:  input.Comment,
		Visited:  input.Visited,
		Helpful:  0,
	}

	config.DB.Create(&review)

	c.JSON(http.StatusCreated, gin.H{
		"message": "Review posted successfully!",
		"data":    review,
	})
}

func MarkHelpful(c *gin.Context) {
	var review models.Review
	id := c.Param("id")

	result := config.DB.First(&review, id)
	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Review not found",
		})
		return
	}

	config.DB.Model(&review).Update(
		"helpful", review.Helpful+1,
	)

	c.JSON(http.StatusOK, gin.H{
		"message": "Marked as helpful!",
	})
}