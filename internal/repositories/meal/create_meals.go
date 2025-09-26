package meal_repo

import (
	"context"
	"fmt"
	"log/slog"

	sq "github.com/Masterminds/squirrel"
	"github.com/gomeal/logger/pkg/logger"
	business "github.com/gomeal/meal-api/internal/services/model"
)

// CreateMeals saves a list of business meals to the database within a transaction.
// Each meal is processed individually to ensure data consistency and simplicity.
func (r *repoImpl) CreateMeals(ctx context.Context, meals []business.Meal) ([]business.Meal, error) {
	var result []business.Meal

	txErr := r.transactor.WithinTranasction(ctx, func(ctx context.Context) error {
		for _, meal := range meals {
			savedMeal, err := r.saveSingleMeal(ctx, meal)
			if err != nil {
				return err
			}
			result = append(result, savedMeal)
		}
		return nil
	})

	if txErr != nil {
		logger.Error(ctx, fmt.Sprintf("[%T].CreateMeals transaction failed", r), slog.Any("error", txErr))
		return nil, fmt.Errorf("transaction failed: %w", txErr)
	}

	return result, nil
}

// saveSingleMeal processes one business meal: ensures related entities exist, saves the meal, and links ingredients.
func (r *repoImpl) saveSingleMeal(ctx context.Context, meal business.Meal) (business.Meal, error) {
	categoryID, err := r.ensureCategory(ctx, meal.Category.Name)
	if err != nil {
		return business.Meal{}, fmt.Errorf("failed to ensure category: %w", err)
	}

	cuisineID, err := r.ensureCuisine(ctx, meal.Cuisine.Name)
	if err != nil {
		return business.Meal{}, fmt.Errorf("failed to ensure cuisine: %w", err)
	}

	mealID, err := r.saveMeal(ctx, meal, categoryID, cuisineID)
	if err != nil {
		return business.Meal{}, fmt.Errorf("failed to save meal: %w", err)
	}

	ingredientIDs, err := r.saveIngredientsForMeal(ctx, mealID, meal.Ingredients)
	if err != nil {
		return business.Meal{}, fmt.Errorf("failed to save ingredients: %w", err)
	}

	return r.buildUpdatedMeal(meal, mealID, categoryID, cuisineID, ingredientIDs), nil
}

// ensureCategory inserts or updates a category and returns its database ID.
func (r *repoImpl) ensureCategory(ctx context.Context, name string) (int64, error) {
	query, args, err := sq.Insert("categories").
		PlaceholderFormat(sq.Dollar).
		Columns("name").
		Values(name).
		Suffix("ON CONFLICT (name) DO UPDATE SET updated_at = ?", r.nowTimer.Now()).
		Suffix("RETURNING id").
		ToSql()

	if err != nil {
		logger.Error(ctx, fmt.Sprintf("[%T].ensureCategory failed to build query", r), slog.Any("error", err))
		return 0, err
	}

	var id int64
	if err := r.db.QueryRow(ctx, query, args...).Scan(&id); err != nil {
		logger.Error(ctx, fmt.Sprintf("[%T].ensureCategory failed to execute", r), slog.Any("error", err))
		return 0, err
	}

	return id, nil
}

// ensureCuisine inserts or updates a cuisine and returns its database ID.
func (r *repoImpl) ensureCuisine(ctx context.Context, name string) (int64, error) {
	query, args, err := sq.Insert("cuisines").
		PlaceholderFormat(sq.Dollar).
		Columns("name").
		Values(name).
		Suffix("ON CONFLICT (name) DO UPDATE SET updated_at = ?", r.nowTimer.Now()).
		Suffix("RETURNING id").
		ToSql()

	if err != nil {
		logger.Error(ctx, fmt.Sprintf("[%T].ensureCuisine failed to build query", r), slog.Any("error", err))
		return 0, err
	}

	var id int64
	if err := r.db.QueryRow(ctx, query, args...).Scan(&id); err != nil {
		logger.Error(ctx, fmt.Sprintf("[%T].ensureCuisine failed to execute", r), slog.Any("error", err))
		return 0, err
	}

	return id, nil
}

// saveMeal inserts or updates a meal and returns its database ID.
func (r *repoImpl) saveMeal(ctx context.Context, meal business.Meal, categoryID, cuisineID int64) (int64, error) {
	onConflictClause := `ON CONFLICT (external_id) DO UPDATE SET
		name = EXCLUDED.name,
		category_id = EXCLUDED.category_id,
		cuisine_id = EXCLUDED.cuisine_id,
		instructions = EXCLUDED.instructions,
		image_url = EXCLUDED.image_url,
		tags = EXCLUDED.tags,
		youtube_url = EXCLUDED.youtube_url,
		recipe_url = EXCLUDED.recipe_url,
		updated_at = ?`

	query, args, err := sq.Insert("meals").
		PlaceholderFormat(sq.Dollar).
		Columns("external_id", "name", "category_id", "cuisine_id", "instructions", "image_url", "tags", "youtube_url", "recipe_url").
		Values(meal.ExternalID, meal.Name, categoryID, cuisineID, meal.Instructions, meal.ImageURL, meal.Tags, meal.YouTubeURL, meal.RecipeURL).
		Suffix(onConflictClause, r.nowTimer.Now()).
		Suffix("RETURNING id").
		ToSql()

	if err != nil {
		logger.Error(ctx, fmt.Sprintf("[%T].saveMeal failed to build query", r), slog.Any("error", err))
		return 0, err
	}

	var id int64
	if err := r.db.QueryRow(ctx, query, args...).Scan(&id); err != nil {
		logger.Error(ctx, fmt.Sprintf("[%T].saveMeal failed to execute", r), slog.Any("error", err))
		return 0, err
	}

	return id, nil
}

// saveIngredientsForMeal ensures all ingredients exist and creates meal-ingredient links.
func (r *repoImpl) saveIngredientsForMeal(ctx context.Context, mealID int64, ingredients []business.MealIngredient) ([]int64, error) {
	var ingredientIDs []int64

	for i, ingredient := range ingredients {
		ingredientID, err := r.ensureIngredient(ctx, ingredient.Name)
		if err != nil {
			return nil, fmt.Errorf("failed to ensure ingredient: %w", err)
		}

		if err := r.linkMealIngredient(ctx, mealID, ingredientID, ingredient.Measure, i+1); err != nil {
			return nil, fmt.Errorf("failed to link meal ingredient: %w", err)
		}

		ingredientIDs = append(ingredientIDs, ingredientID)
	}

	return ingredientIDs, nil
}

// ensureIngredient inserts or updates an ingredient and returns its database ID.
func (r *repoImpl) ensureIngredient(ctx context.Context, name string) (int64, error) {
	query, args, err := sq.Insert("ingredients").
		PlaceholderFormat(sq.Dollar).
		Columns("name").
		Values(name).
		Suffix("ON CONFLICT (name) DO UPDATE SET updated_at = ?", r.nowTimer.Now()).
		Suffix("RETURNING id").
		ToSql()

	if err != nil {
		logger.Error(ctx, fmt.Sprintf("[%T].ensureIngredient failed to build query", r), slog.Any("error", err))
		return 0, err
	}

	var id int64
	if err := r.db.QueryRow(ctx, query, args...).Scan(&id); err != nil {
		logger.Error(ctx, fmt.Sprintf("[%T].ensureIngredient failed to execute", r), slog.Any("error", err))
		return 0, err
	}

	return id, nil
}

// linkMealIngredient creates or updates a meal-ingredient association.
func (r *repoImpl) linkMealIngredient(ctx context.Context, mealID, ingredientID int64, measure string, position int) error {
	onConflictClause := `ON CONFLICT (meal_id, ingredient_id) DO UPDATE SET
		measure = EXCLUDED.measure,
		position = EXCLUDED.position,
		updated_at = ?`

	query, args, err := sq.Insert("meal_ingredients").
		PlaceholderFormat(sq.Dollar).
		Columns("meal_id", "ingredient_id", "measure", "position").
		Values(mealID, ingredientID, measure, position).
		Suffix(onConflictClause, r.nowTimer.Now()).
		ToSql()

	if err != nil {
		logger.Error(ctx, fmt.Sprintf("[%T].linkMealIngredient failed to build query", r), slog.Any("error", err))
		return err
	}

	if _, err := r.db.Exec(ctx, query, args...); err != nil {
		logger.Error(ctx, fmt.Sprintf("[%T].linkMealIngredient failed to execute", r), slog.Any("error", err))
		return err
	}

	return nil
}

// buildUpdatedMeal creates a business meal with populated database IDs.
func (r *repoImpl) buildUpdatedMeal(originalMeal business.Meal, mealID, categoryID, cuisineID int64, ingredientIDs []int64) business.Meal {
	updatedMeal := originalMeal
	updatedMeal.ID = mealID
	updatedMeal.Category.ID = categoryID
	updatedMeal.Cuisine.ID = cuisineID

	for i, ingredientID := range ingredientIDs {
		if i < len(updatedMeal.Ingredients) {
			updatedMeal.Ingredients[i].ID = ingredientID
		}
	}

	return updatedMeal
}
