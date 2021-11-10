package router

import (
	"math"

	"github.com/dcruselex/food/internal/rapidapi"
)

func getFilteredFoods(option int) []rapidapi.Recipe {
	max := 30
	min := 0
	if option == 1 {
		max = 60
		min = 31
	} else if option == 2 {
		max = math.MaxInt
		min = 61
	}

	var filteredFoods []rapidapi.Recipe
	for _, food := range foods {
		if food.ReadyInMinutes >= min && food.ReadyInMinutes <= max {
			filteredFoods = append(filteredFoods, food)
		}
	}

	return filteredFoods
}
