-- name: GetUserByEmail :one
SELECT * FROM users
WHERE email = $1 LIMIT 1;

-- name: UpdateUserByOTP :exec
UPDATE users
SET isverified = TRUE
WHERE email = $1 AND otp = $2;

-- name: CreateUser :one
INSERT INTO users (name, email, password, isverified, otp)
VALUES ($1, $2, $3, false, $4)
RETURNING *;
