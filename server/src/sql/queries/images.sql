-- name: BulkCreateImages :many
INSERT INTO images (ad_id, label, image_number, url)
VALUES (
    unnest($1::TEXT ARRAY),
    unnest($2::TEXT ARRAY),
    unnest($3::INT ARRAY),
    unnest($4::TEXT ARRAY)
) ON CONFLICT (url) DO NOTHING RETURNING *;

-- name: GetThumbnailsForAgent :many
SELECT * FROM (
	SELECT r.ad_id, json_agg(json_build_object('image_url', i.url, 'image_number', i.image_number)) AS images from agents a
	JOIN agent_params ap on ap.agent_id = a.id
	JOIN params p on ap.params_id = p.id
	JOIN results r on r.params_id = p.id
	JOIN images i on i.ad_id = r.ad_id
	WHERE a.id = $1 and label = 'thumbnail'
	GROUP BY r.ad_id
) as images
JOIN ads on ads.id = images.ad_id
LIMIT 3;
