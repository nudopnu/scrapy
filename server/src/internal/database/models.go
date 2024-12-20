// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0

package database

import (
	"database/sql"
	"time"
)

type Ad struct {
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

type Agent struct {
	ID            int32
	Name          string
	UserID        int32
	LastFetchedAt sql.NullTime
	CreatedAt     time.Time
	UpdatedAt     time.Time
}

type AgentParam struct {
	ID        int32
	AgentID   int32
	ParamsID  int32
	CreatedAt sql.NullTime
	UpdatedAt sql.NullTime
}

type Image struct {
	ID          int32
	AdID        string
	ImageNumber int32
	Label       string
	Url         string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type Location struct {
	ID         int32
	PostalCode string
	LocationID string
	CreatedAt  time.Time
	UpdatedAt  time.Time
}

type Param struct {
	ID         int32
	Keyword    string
	LocationID string
	Distance   int32
	CreatedAt  sql.NullTime
	UdpatedAt  sql.NullTime
}

type RefreshToken struct {
	Token     string
	UserID    int32
	ExpiresAt time.Time
	RevokedAt sql.NullTime
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Result struct {
	ID        int32
	ParamsID  int32
	AdID      string
	Status    string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type User struct {
	ID             int32
	Username       string
	Role           string
	HashedPassword string
	CreatedAt      time.Time
	UpdatedAt      time.Time
}
