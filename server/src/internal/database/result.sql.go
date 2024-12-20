// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: result.sql

package database

import (
	"context"
	"database/sql"
	"encoding/json"
	"time"

	"github.com/lib/pq"
)

const bulkCreateResults = `-- name: BulkCreateResults :many
INSERT INTO results (params_id, ad_id, status)
VALUES (
    unnest($1:: INT ARRAY),
    unnest($2:: TEXT ARRAY),
    'new'
)
ON CONFLICT (params_id, ad_id) DO NOTHING
RETURNING id, params_id, ad_id, status, created_at, updated_at
`

type BulkCreateResultsParams struct {
	Column1 []int32
	Column2 []string
}

func (q *Queries) BulkCreateResults(ctx context.Context, arg BulkCreateResultsParams) ([]Result, error) {
	rows, err := q.db.QueryContext(ctx, bulkCreateResults, pq.Array(arg.Column1), pq.Array(arg.Column2))
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Result
	for rows.Next() {
		var i Result
		if err := rows.Scan(
			&i.ID,
			&i.ParamsID,
			&i.AdID,
			&i.Status,
			&i.CreatedAt,
			&i.UpdatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const createResult = `-- name: CreateResult :one
INSERT INTO results (params_id, ad_id, status)
VALUES (
    $1, 
    $2,
    $3
) RETURNING id, params_id, ad_id, status, created_at, updated_at
`

type CreateResultParams struct {
	ParamsID int32
	AdID     string
	Status   string
}

func (q *Queries) CreateResult(ctx context.Context, arg CreateResultParams) (Result, error) {
	row := q.db.QueryRowContext(ctx, createResult, arg.ParamsID, arg.AdID, arg.Status)
	var i Result
	err := row.Scan(
		&i.ID,
		&i.ParamsID,
		&i.AdID,
		&i.Status,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const getResultsByParamId = `-- name: GetResultsByParamId :many
SELECT results.id, results.params_id, results.ad_id, results.status, results.created_at, results.updated_at from results
WHERE params_id = $1
`

func (q *Queries) GetResultsByParamId(ctx context.Context, paramsID int32) ([]Result, error) {
	rows, err := q.db.QueryContext(ctx, getResultsByParamId, paramsID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Result
	for rows.Next() {
		var i Result
		if err := rows.Scan(
			&i.ID,
			&i.ParamsID,
			&i.AdID,
			&i.Status,
			&i.CreatedAt,
			&i.UpdatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const listResultsFromAgent = `-- name: ListResultsFromAgent :many
SELECT ad_id, images, id, title, price, description, location, postal_code, category_id, posted_at, link, created_at, updated_at FROM (
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
OFFSET $3
`

type ListResultsFromAgentParams struct {
	ID     int32
	Limit  int32
	Offset int32
}

type ListResultsFromAgentRow struct {
	AdID        string
	Images      json.RawMessage
	ID          string
	Title       string
	Price       sql.NullString
	Description sql.NullString
	Location    sql.NullString
	PostalCode  sql.NullString
	CategoryID  sql.NullString
	PostedAt    sql.NullString
	Link        sql.NullString
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

func (q *Queries) ListResultsFromAgent(ctx context.Context, arg ListResultsFromAgentParams) ([]ListResultsFromAgentRow, error) {
	rows, err := q.db.QueryContext(ctx, listResultsFromAgent, arg.ID, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []ListResultsFromAgentRow
	for rows.Next() {
		var i ListResultsFromAgentRow
		if err := rows.Scan(
			&i.AdID,
			&i.Images,
			&i.ID,
			&i.Title,
			&i.Price,
			&i.Description,
			&i.Location,
			&i.PostalCode,
			&i.CategoryID,
			&i.PostedAt,
			&i.Link,
			&i.CreatedAt,
			&i.UpdatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateResultExpired = `-- name: UpdateResultExpired :exec
UPDATE results SET status='expired' 
WHERE NOT ad_id = ANY($1::TEXT ARRAY)
`

func (q *Queries) UpdateResultExpired(ctx context.Context, dollar_1 []string) error {
	_, err := q.db.ExecContext(ctx, updateResultExpired, pq.Array(dollar_1))
	return err
}

const updateResultUpdated = `-- name: UpdateResultUpdated :exec
UPDATE results SET status='updated'
WHERE ad_id = ANY($1::TEXT ARRAY)
`

func (q *Queries) UpdateResultUpdated(ctx context.Context, dollar_1 []string) error {
	_, err := q.db.ExecContext(ctx, updateResultUpdated, pq.Array(dollar_1))
	return err
}
