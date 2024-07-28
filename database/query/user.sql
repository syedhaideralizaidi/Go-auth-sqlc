-- name: CreateUser :one
INSERT INTO users (
    email,
    username,
    phone_number,
    password,
    role
) VALUES (
             $1, $2, $3, $4, $5
         ) RETURNING *;

-- name: GetUser :one
SELECT * FROM users
WHERE id = $1 LIMIT 1;