-- +goose Up
-- +goose StatementBegin
CREATE TABLE meals (
    id SERIAL PRIMARY KEY,
    external_id INTEGER UNIQUE, -- idMeal из TheMealDB API
    name VARCHAR(255) NOT NULL, -- strMeal
    instructions TEXT NOT NULL, -- strInstructions
    image_url TEXT, -- strMealThumb
    category_id INTEGER REFERENCES categories(id),
    cuisine_id INTEGER REFERENCES cuisines(id),
    tags TEXT ARRAY, -- strTags
    youtube_url TEXT, -- strYoutube
    recipe_url TEXT, -- strSource
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX idx_meals_external_id ON meals(external_id);
CREATE INDEX idx_meals_category_id ON meals(category_id);
CREATE INDEX idx_meals_cuisine_id ON meals(cuisine_id);
CREATE INDEX idx_meals_name ON meals(name);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP INDEX IF EXISTS idx_meals_external_id;
DROP INDEX IF EXISTS idx_meals_category_id;
DROP INDEX IF EXISTS idx_meals_cuisine_id;
DROP INDEX IF EXISTS idx_meals_name;
DROP TABLE IF EXISTS meals;
-- +goose StatementEnd
