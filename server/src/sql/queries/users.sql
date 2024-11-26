-- name: RegisterUser :one
INSERT INTO users (username, hashed_password, role)
VALUES (
    $1,
    $2,
    'user'
) RETURNING *;

-- name: RegisterAdmin :one
INSERT INTO users (username, hashed_password, role)
VALUES (
    $1,
    $2,
    'admin'
) RETURNING *;

-- name: GetUserByUsername :one
SELECT * FROM users
WHERE username = $1;

-- name: GetUserById :one
SELECT * FROM users
WHERE id = $1;

-- name: GetAdmin :one
SELECT * FROM users
WHERE role = 'admin';

-- name: ListUsers :many
SELECT * FROM users;

-- name: Reset :exec
DELETE FROM users;