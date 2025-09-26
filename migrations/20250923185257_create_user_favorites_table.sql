-- +goose Up
-- +goose StatementBegin
CREATE TABLE user_favorites (
    id SERIAL PRIMARY KEY,
    user_id uuid NOT NULL,
    meal_id INTEGER REFERENCES meals(id) ON DELETE CASCADE,
    added_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    UNIQUE(user_id, meal_id)
);

CREATE INDEX idx_user_favorites_user_id ON user_favorites(user_id);
CREATE INDEX idx_user_favorites_meal_id ON user_favorites(meal_id);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP INDEX IF EXISTS idx_user_favorites_user_id;
DROP INDEX IF EXISTS idx_user_favorites_meal_id;
DROP TABLE IF EXISTS user_favorites;
-- +goose StatementEnd
