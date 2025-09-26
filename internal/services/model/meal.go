package business

type (
	Meal struct {
		ID           int64
		ExternalID   int64
		Name         string
		Category     MealCategory
		Cuisine      MealCuisine
		Instructions string
		ImageURL     string
		Tags         []string
		YouTubeURL   string
		Ingredients  []MealIngredient
		RecipeURL    string
	}

	MealCategory struct {
		ID   int64
		Name string
	}

	MealCuisine struct {
		ID   int64
		Name string
	}

	MealIngredient struct {
		ID      int64
		Name    string
		Measure string
	}
)
