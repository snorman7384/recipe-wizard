// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.20.0
// source: grocery_lists.sql

package database

import (
	"context"
	"database/sql"
	"time"
)

const createGroceryList = `-- name: CreateGroceryList :execresult
INSERT INTO grocery_lists (created_at, updated_at, name, owner_id)
VALUES (?, ?, ?, ?)
`

type CreateGroceryListParams struct {
	CreatedAt time.Time
	UpdatedAt time.Time
	Name      string
	OwnerID   int64
}

func (q *Queries) CreateGroceryList(ctx context.Context, arg CreateGroceryListParams) (sql.Result, error) {
	return q.db.ExecContext(ctx, createGroceryList,
		arg.CreatedAt,
		arg.UpdatedAt,
		arg.Name,
		arg.OwnerID,
	)
}

const getGroceryList = `-- name: GetGroceryList :one
SELECT id, created_at, updated_at, name, owner_id FROM grocery_lists
WHERE id = ?
`

func (q *Queries) GetGroceryList(ctx context.Context, id int64) (GroceryList, error) {
	row := q.db.QueryRowContext(ctx, getGroceryList, id)
	var i GroceryList
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.Name,
		&i.OwnerID,
	)
	return i, err
}

const getGroceryListsForUser = `-- name: GetGroceryListsForUser :many
SELECT id, created_at, updated_at, name, owner_id FROM grocery_lists
WHERE owner_id = ?
`

func (q *Queries) GetGroceryListsForUser(ctx context.Context, ownerID int64) ([]GroceryList, error) {
	rows, err := q.db.QueryContext(ctx, getGroceryListsForUser, ownerID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GroceryList
	for rows.Next() {
		var i GroceryList
		if err := rows.Scan(
			&i.ID,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.Name,
			&i.OwnerID,
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

const getIngredientsInGroceryList = `-- name: GetIngredientsInGroceryList :many
SELECT i.id, i.created_at, i.updated_at, i.name, i.description, i.recipe_id, i.amount, i.units, i.standard_amount, i.standard_units FROM ingredients i
JOIN recipes r ON r.id = i.recipe_id
JOIN recipe_instances ri ON r.id = ri.recipe_id
WHERE ri.grocery_list_id = ?
`

func (q *Queries) GetIngredientsInGroceryList(ctx context.Context, groceryListID int64) ([]Ingredient, error) {
	rows, err := q.db.QueryContext(ctx, getIngredientsInGroceryList, groceryListID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Ingredient
	for rows.Next() {
		var i Ingredient
		if err := rows.Scan(
			&i.ID,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.Name,
			&i.Description,
			&i.RecipeID,
			&i.Amount,
			&i.Units,
			&i.StandardAmount,
			&i.StandardUnits,
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
