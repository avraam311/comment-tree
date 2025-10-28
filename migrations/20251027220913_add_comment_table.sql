-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS comment (
    id SERIAL PRIMARY KEY,
    text TEXT NOT NULL,
    parent_id INT NOT NULL
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS comment;
-- +goose StatementEnd
