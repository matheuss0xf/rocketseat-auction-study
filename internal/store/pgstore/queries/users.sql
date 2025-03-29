-- name: CreateUser :one
INSERT INTO users ("id", "user_name", "email", "password", "bio")
VALUES ($1, $2, $3, $4, $5)
RETURNING id;

-- name: GetUserById :one
SELECT ("id", "user_name", "email", "password", "bio", "created_at", "updated_at") FROM users WHERE id = $1;

-- name: GetUserByEmail :one
SELECT id, user_name, password, email, bio, created_at, updated_at FROM users WHERE email = $1;