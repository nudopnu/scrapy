-- name: GetRefreshToken :one
SELECT * from refresh_tokens 
WHERE token = $1;

-- name: CreateRefreshToken :one
INSERT INTO refresh_tokens (token, user_id, expires_at)
VALUES (
    $1,
    $2,
    $3
) RETURNING *;