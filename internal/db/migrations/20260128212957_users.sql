-- +goose Up
CREATE TABLE users(
    id SERIAL PRIMARY KEY,
    username TEXT,
    chat_id BIGINT,
    last_message_id BIGINT,
    login TEXT,
    password TEXT
);
-- +goose StatementBegin
SELECT 'up SQL query';
-- +goose StatementEnd

-- +goose Down
DROP TABLE IF EXISTS users;
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
