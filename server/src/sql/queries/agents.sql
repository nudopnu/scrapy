-- name: ListAgents :many
SELECT * FROM agents;

-- name: ListAgentsWithImages :many
SELECT a.*, 
       COALESCE((SELECT i.url 
        FROM images i 
        JOIN results r ON i.ad_id = r.ad_id 
        JOIN agent_params ap ON r.params_id = ap.params_id 
        WHERE ap.agent_id = a.id AND i.label = 'thumbnail' AND i.image_number = 0
        ORDER BY i.created_at DESC
        LIMIT 1), '')::TEXT AS thumbnail
FROM agents a;

-- name: CreateSearchAgent :one
INSERT INTO agents (name, user_id) 
VALUES (
    $1,
    $2
) RETURNING *;

-- name: MarkAgentUpdated :exec
UPDATE agents 
SET updated_at=CURRENT_TIMESTAMP, last_fetched_at=CURRENT_TIMESTAMP 
WHERE id=$1;

-- name: GetNextAgentToUpdate :one
SELECT * from agents
WHERE last_fetched_at + interval '3 minutes' < CURRENT_TIMESTAMP
OR last_fetched_at IS NULL
ORDER BY last_fetched_at ASC NULLS FIRST;

-- name: GetAgentByName :one
SELECT * from agents
WHERE name = $1;