-- name: ListLocations :many
SELECT * from locations;

-- name: GetLocationByPostalCode :one
SELECT * from locations
WHERE postal_code = $1;

-- name: AddLocation :one
INSERT INTO locations (postal_code, location_id)
VALUES (
    $1, $2
) RETURNING *;