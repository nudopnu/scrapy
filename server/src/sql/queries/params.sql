-- name: CreateSearchParams :one
INSERT INTO params (keyword, location_id, distance)
VALUES (
    $1,
    $2,
    $3
) RETURNING *;