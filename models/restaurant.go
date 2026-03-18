package models

import (
	"kerala-food-finder/config"

	"gorm.io/gorm"
)

type Restaurant struct {
	gorm.Model
	Name     string
	City     string
	Area     string
	Location string
	Verified bool
	Dishes   []Dish
	Reels    []Reel
}


func FindOrCreateRestaurant(
	name string,
	city string,
	area string,
) Restaurant {

	var restaurant Restaurant

	result := config.DB.Where(
		"name = ? AND city = ?",
		name, city,
	).First(&restaurant)

	if result.Error != nil {
		restaurant = Restaurant{
			Name: name,
			City: city,
			Area: area,
		}
		config.DB.Create(&restaurant)
	}

	return restaurant
}