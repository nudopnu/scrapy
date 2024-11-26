-- name: AddSearchParamToAgent :one
INSERT into agent_params (agent_id, params_id)
VALUES (
    $1, 
    $2
) RETURNING *;

-- name: GetSearchParamsBySearchAgent :many
SELECT p.* from agent_params ap
JOIN agents a ON a.id = ap.agent_id
JOIN params p ON p.id = ap.params_id
WHERE a.id = $1;