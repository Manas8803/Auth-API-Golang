-- name: GetUser :one
SELECT * FROM users
WHERE email = $1 LIMIT 1;

-- name: CheckUserExists :one
SELECT EXISTS (SELECT 1 FROM users WHERE email = $1) AS exists;

-- name: CreateUser :one
INSERT INTO users (name, email, password)
VALUES ($1, $2, $3)
RETURNING *;

