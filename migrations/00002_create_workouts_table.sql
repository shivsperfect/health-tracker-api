-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS workouts (
     id BIGSERIAL PRIMARY KEY,
    -- user_id
    title VARCHAR(255) NOT NULL,
    description TEXT,
    duration_minutes INT NOT NULL,
    calories_burned INT,
    created_at TIMESTAMPTZ DEFAULT NOW(),
    updated_at TIMESTAMPTZ DEFAULT NOW()
    );
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS workouts;
-- +goose StatementEnd

