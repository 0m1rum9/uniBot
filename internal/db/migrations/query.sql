-- name: CreateUser :exec
INSERT INTO users(username, chat_id, last_message_id, login, password)
VALUES($1, $2, $3, $4, $5);
