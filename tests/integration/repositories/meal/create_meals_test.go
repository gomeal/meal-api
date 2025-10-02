package meal_repo_test

import (
	"context"
	"testing"

	business "github.com/gomeal/meal-api/internal/services/model"
	db_test_tools "github.com/gomeal/meal-api/tests/integration/tools/db"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func (s *Suite) TestCreateMeals() {
	ctx := context.Background()

	s.T().Run("valid meals create request", func(t *testing.T) {
		require.NoError(t, db_test_tools.TruncateAll(ctx))

		meals := []business.Meal{
			{
				ExternalID: 52817,
				Name:       "Beef Burger",
				Category: business.MealCategory{
					Name: "Fast food",
				},
				Cuisine: business.MealCuisine{
					Name: "American",
				},
				Instructions: "Do a burger idk...",
				ImageURL:     "https://burger.image/image.jpg",
				Tags:         []string{"spicy", "beef", "onions"},
				YouTubeURL:   "https://www.youtube.com/watch?v=tuDbSVyClzI",
				Ingredients: []business.MealIngredient{
					{
						Name:    "Beef",
						Measure: "1 kotleta",
					},
					{
						Name:    "Bulka",
						Measure: "2 pieces",
					},
				},
			},
			{
				ExternalID: 63414,
				Name:       "Chicken Burger",
				Category: business.MealCategory{
					Name: "Not a fast food",
				},
				Cuisine: business.MealCuisine{
					Name: "Russian",
				},
				Instructions: "Do another burger idk...",
				ImageURL:     "https://burger.image/image2.jpg",
				Tags:         []string{"spicy", "beef", "onions"},
				YouTubeURL:   "https://www.youtube.com/watch?v=BIG1h2vG-Qg",
				Ingredients: []business.MealIngredient{
					{
						Name:    "Chicken Beef",
						Measure: "1 kotleta",
					},
					{
						Name:    "Bulka",
						Measure: "2 pieces",
					},
				},
			},
		}

		createdMeals, err := s.mealRepository.CreateMeals(ctx, meals)
		require.NoError(t, err)

		assert.Equal(t, 2, len(createdMeals))
		expectedMeals := []business.Meal{
			{
				ID:         1,
				ExternalID: 52817,
				Name:       "Beef Burger",
				Category: business.MealCategory{
					ID:   1,
					Name: "Fast food",
				},
				Cuisine: business.MealCuisine{
					ID:   1,
					Name: "American",
				},
				Instructions: "Do a burger idk...",
				ImageURL:     "https://burger.image/image.jpg",
				Tags:         []string{"spicy", "beef", "onions"},
				YouTubeURL:   "https://www.youtube.com/watch?v=tuDbSVyClzI",
				Ingredients: []business.MealIngredient{
					{
						ID:      1,
						Name:    "Beef",
						Measure: "1 kotleta",
					},
					{
						ID:      2,
						Name:    "Bulka",
						Measure: "2 pieces",
					},
				},
			},
			{
				ID:         2,
				ExternalID: 63414,
				Name:       "Chicken Burger",
				Category: business.MealCategory{
					ID:   2,
					Name: "Not a fast food",
				},
				Cuisine: business.MealCuisine{
					ID:   2,
					Name: "Russian",
				},
				Instructions: "Do another burger idk...",
				ImageURL:     "https://burger.image/image2.jpg",
				Tags:         []string{"spicy", "beef", "onions"},
				YouTubeURL:   "https://www.youtube.com/watch?v=BIG1h2vG-Qg",
				Ingredients: []business.MealIngredient{
					{
						ID:      3,
						Name:    "Chicken Beef",
						Measure: "1 kotleta",
					},
					{
						ID:      2,
						Name:    "Bulka",
						Measure: "2 pieces",
					},
				},
			},
		}

		assert.Equal(t, expectedMeals, createdMeals)
	})
}
