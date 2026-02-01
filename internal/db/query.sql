-- name: CreateUser :exec
INSERT INTO users(username, chat_id, login, password)
VALUES($1, $2, $3, $4);

-- name: GetUser :one
SELECT * FROM users
WHERE chat_id=$1 LIMIT 1;
