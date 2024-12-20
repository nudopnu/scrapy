// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: params.sql

package database

import (
	"context"
)

const createSearchParams = `-- name: CreateSearchParams :one
INSERT INTO params (keyword, location_id, distance)
VALUES (
    $1,
    $2,
    $3
) RETURNING id, keyword, location_id, distance, created_at, udpated_at
`

type CreateSearchParamsParams struct {
	Keyword    string
	LocationID string
	Distance   int32
}

func (q *Queries) CreateSearchParams(ctx context.Context, arg CreateSearchParamsParams) (Param, error) {
	row := q.db.QueryRowContext(ctx, createSearchParams, arg.Keyword, arg.LocationID, arg.Distance)
	var i Param
	err := row.Scan(
		&i.ID,
		&i.Keyword,
		&i.LocationID,
		&i.Distance,
		&i.CreatedAt,
		&i.UdpatedAt,
	)
	return i, err
}
