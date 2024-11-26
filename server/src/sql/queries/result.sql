-- name: ListResultsFromAgent :many
SELECT * FROM (
	SELECT r.ad_id, json_agg(json_build_object('image_url', i.url, 'image_number', i.image_number)) AS images from agents a
	JOIN agent_params ap on ap.agent_id = a.id
	JOIN params p on ap.params_id = p.id
	JOIN results r on r.params_id = p.id
	JOIN images i on i.ad_id = r.ad_id
	WHERE a.id = $1 AND i.label = 'extraLarge'
	GROUP BY r.ad_id
) as images
JOIN ads on ads.id = images.ad_id
ORDER BY ads.created_at DESC
LIMIT $2
OFFSET $3;

-- name: GetResultsByParamId :many
SELECT results.* from results
WHERE params_id = $1;

-- name: CreateResult :one
INSERT INTO results (params_id, ad_id, status)
VALUES (
    $1, 
    $2,
    $3
) RETURNING *;

-- name: UpdateResultUpdated :exec
UPDATE results SET status='updated'
WHERE ad_id = ANY($1::TEXT ARRAY);

-- name: UpdateResultExpired :exec
UPDATE results SET status='expired' 
WHERE NOT ad_id = ANY($1::TEXT ARRAY);

-- name: BulkCreateResults :many
INSERT INTO results (params_id, ad_id, status)
VALUES (
    unnest($1:: INT ARRAY),
    unnest($2:: TEXT ARRAY),
    'new'
)
ON CONFLICT (params_id, ad_id) DO NOTHING
RETURNING *;