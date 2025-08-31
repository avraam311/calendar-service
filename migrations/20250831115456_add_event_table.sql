-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS events (
    user_id INT PRIMARY KEY,
    event TEXT,
    date TIMESTAMPTZ,
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS events;
-- +goose StatementEnd
