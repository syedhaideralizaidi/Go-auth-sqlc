-- name: CreateUser :one
INSERT INTO users (
    email,
    username,
    phone_number,
    password,
    role,
    is_verified
) VALUES (
             $1, $2, $3, $4, $5, $6
         ) RETURNING *;

-- name: GetUserByEmail :one
SELECT * FROM users WHERE email = $1 LIMIT 1;

-- name: VerifyUser :exec
UPDATE users SET is_verified = true WHERE email = $1;

-- name: GetUser :one
SELECT * FROM users
WHERE id = $1 LIMIT 1;

-- name: ListUsers :many
SELECT * FROM users
ORDER BY id
    LIMIT $1
OFFSET $2;

-- name: UpdateUser :one
UPDATE users
SET username = $2
WHERE id = $1
    RETURNING *;

-- name: DeleteUser :exec
DELETE FROM users WHERE id = $1;