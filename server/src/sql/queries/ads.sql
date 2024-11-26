-- name: CreateAd :one
INSERT INTO ads (id, title, price, description, location, postal_code, category_id, posted_at, link)
VALUES (
    $1,
    $2,
    $3,
    $4,
    $5,
    $6,
    $7,
    $8,
    $9
) RETURNING *;

-- name: BulkCreateAds :many
INSERT INTO ads (id, title, price, description, location, postal_code, category_id, posted_at, link)
VALUES (
    unnest($1::TEXT ARRAY),
    unnest($2::TEXT ARRAY),
    unnest($3::TEXT ARRAY),
    unnest($4::TEXT ARRAY),
    unnest($5::TEXT ARRAY),
    unnest($6::TEXT ARRAY),
    unnest($7::TEXT ARRAY),
    unnest($8::TEXT ARRAY),
    unnest($9::TEXT ARRAY)
)
ON CONFLICT (id) DO NOTHING
RETURNING *;

-- name: GetAdByEbayId :one
SELECT * FROM ads WHERE id = $1;

-- name: GetNumberOfDuplicates :one
SELECT COUNT(*) FROM ads
WHERE id=ANY($1::TEXT ARRAY);