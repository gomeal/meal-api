-- +goose Up
-- +goose StatementBegin
CREATE TABLE meal_ingredients (
    id SERIAL PRIMARY KEY,
    meal_id INTEGER REFERENCES meals(id) ON DELETE CASCADE,
    ingredient_id INTEGER REFERENCES ingredients(id),
    measure VARCHAR(100) NOT NULL, -- strMeasure1, strMeasure2, etc.
    position INTEGER NOT NULL, -- порядок в рецепте (1, 2, 3...)
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    UNIQUE(meal_id, ingredient_id)
);

CREATE INDEX idx_meal_ingredients_meal_id ON meal_ingredients(meal_id);
CREATE INDEX idx_meal_ingredients_ingredient_id ON meal_ingredients(ingredient_id);
CREATE INDEX idx_meal_ingredients_position ON meal_ingredients(meal_id, position);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP INDEX IF EXISTS idx_meal_ingredients_meal_id;
DROP INDEX IF EXISTS idx_meal_ingredients_ingredient_id;
DROP INDEX IF EXISTS idx_meal_ingredients_position;
DROP TABLE IF EXISTS meal_ingredients;
-- +goose StatementEnd
