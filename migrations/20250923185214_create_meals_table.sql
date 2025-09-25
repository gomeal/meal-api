-- +goose Up
-- +goose StatementBegin
CREATE TABLE meals (
    id SERIAL PRIMARY KEY,
    external_id VARCHAR(20) UNIQUE, -- idMeal из TheMealDB API
    name VARCHAR(255) NOT NULL, -- strMeal
    instructions TEXT, -- strInstructions
    image_url TEXT, -- strMealThumb
    category_id INTEGER REFERENCES categories(id),
    cuisine_id INTEGER REFERENCES cuisines(id),
    tags TEXT, -- strTags (comma-separated, можно потом вынести в отдельную таблицу)
    youtube_url TEXT, -- strYoutube
    source_url TEXT, -- strSource
    cooking_time_minutes INTEGER DEFAULT 0,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX idx_meals_external_id ON meals(external_id);
CREATE INDEX idx_meals_category_id ON meals(category_id);
CREATE INDEX idx_meals_cuisine_id ON meals(cuisine_id);
CREATE INDEX idx_meals_name ON meals(name);
CREATE INDEX idx_meals_cooking_time ON meals(cooking_time_minutes);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP INDEX IF EXISTS idx_meals_external_id;
DROP INDEX IF EXISTS idx_meals_category_id;
DROP INDEX IF EXISTS idx_meals_cuisine_id;
DROP INDEX IF EXISTS idx_meals_name;
DROP INDEX IF EXISTS idx_meals_cooking_time;
DROP TABLE IF EXISTS meals;
-- +goose StatementEnd
