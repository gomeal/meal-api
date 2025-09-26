package domain

import (
	"database/sql"
	"time"
)

type (
	Meal struct {
		ID           int64
		ExternalID   int64
		Name         string
		CategoryID   int64
		CuisineID    int64
		Instructions string
		ImageURL     sql.NullString
		Tags         []string
		YouTubeURL   sql.NullString
		RecipeURL    sql.NullString
		CreatedAt    time.Time
		UpdatedAt    time.Time
	}

	MealCategory struct {
		ID        int64
		Name      string
		CreatedAt time.Time
		UpdatedAt time.Time
	}

	MealCuisine struct {
		ID        int64
		Name      string
		CreatedAt time.Time
		UpdatedAt time.Time
	}

	MealIngredient struct {
		ID        int64
		Name      string
		CreatedAt time.Time
		UpdatedAt time.Time
	}

	MealIngridientsLink struct {
		ID           int64
		MealID       int64
		IngridientID int64
		Measure      string
		Position     int64
		UpdatedAt    time.Time
		CreatedAt    time.Time
	}
)
