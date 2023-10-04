// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.20.0

package database

import (
	"database/sql"
	"time"
)

type Ingredient struct {
	ID          int64
	CreatedAt   time.Time
	UpdatedAt   time.Time
	Name        string
	Description sql.NullString
	RecipeID    int64
}

type Recipe struct {
	ID          int64
	CreatedAt   time.Time
	UpdatedAt   time.Time
	Name        string
	Description sql.NullString
	Url         sql.NullString
	PrepTime    sql.NullString
	CookTime    sql.NullString
	TotalTime   sql.NullString
}

type User struct {
	ID             int64
	CreatedAt      time.Time
	UpdatedAt      time.Time
	Username       string
	HashedPassword string
	FirstName      sql.NullString
	LastName       sql.NullString
}
