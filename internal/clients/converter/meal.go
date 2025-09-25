package transport_converter

import (
	"strconv"
	"strings"

	transport "github.com/gomeal/meal-api/internal/clients/model"
	business "github.com/gomeal/meal-api/internal/services/model"
	"github.com/gomeal/meal-api/internal/utils"
	"github.com/samber/lo"
)

func TransportMealToBusinessMeal(meal transport.Meal) business.Meal {
	var (
		ingridients = meal.GetIngridients()
		measures    = meal.GetMeasures()
	)

	return business.Meal{
		ExternalID: int64(utils.Must(strconv.Atoi(meal.IDMeal))),
		Name:       meal.StrMeal,
		Category: business.MealCategory{
			Name: meal.StrCategory,
		},
		Cuisine: business.MealCuisine{
			Name: meal.StrArea,
		},
		Instructions: meal.StrInstructions,
		ImageURL:     meal.StrMealThumb,
		Tags:         strings.Split(meal.StrTags, ","),
		YouTubeURL:   meal.StrYoutube,
		Ingridients: lo.Filter(lo.Map(ingridients, func(ingridientName string, idx int) business.MealIngridient {
			return business.MealIngridient{
				Name:    ingridientName,
				Measure: measures[idx],
			}
		}), func(ingridient business.MealIngridient, idx int) bool {
			return len(ingridient.Name) > 0
		}),
		RecipeURL: meal.StrSource,
	}
}
