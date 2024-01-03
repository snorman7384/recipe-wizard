// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.20.0
// source: ingredients.sql

package database

import (
	"context"
	"database/sql"
	"time"
)

const createIngredient = `-- name: CreateIngredient :execresult
INSERT INTO ingredients(created_at, updated_at, name, description, amount, units, standard_amount, standard_units, recipe_id)
VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?)
`

type CreateIngredientParams struct {
	CreatedAt      time.Time
	UpdatedAt      time.Time
	Name           string
	Description    sql.NullString
	Amount         float64
	Units          string
	StandardAmount float64
	StandardUnits  string
	RecipeID       int64
}

func (q *Queries) CreateIngredient(ctx context.Context, arg CreateIngredientParams) (sql.Result, error) {
	return q.db.ExecContext(ctx, createIngredient,
		arg.CreatedAt,
		arg.UpdatedAt,
		arg.Name,
		arg.Description,
		arg.Amount,
		arg.Units,
		arg.StandardAmount,
		arg.StandardUnits,
		arg.RecipeID,
	)
}

const getIngredient = `-- name: GetIngredient :one
SELECT id, created_at, updated_at, name, description, recipe_id, amount, units, standard_amount, standard_units FROM ingredients
WHERE id = ?
`

func (q *Queries) GetIngredient(ctx context.Context, id int64) (Ingredient, error) {
	row := q.db.QueryRowContext(ctx, getIngredient, id)
	var i Ingredient
	err := row.Scan(
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
	)
	return i, err
}

const getIngredientsForRecipe = `-- name: GetIngredientsForRecipe :many
SELECT id, created_at, updated_at, name, description, recipe_id, amount, units, standard_amount, standard_units FROM ingredients
WHERE recipe_id = ?
ORDER BY id
`

func (q *Queries) GetIngredientsForRecipe(ctx context.Context, recipeID int64) ([]Ingredient, error) {
	rows, err := q.db.QueryContext(ctx, getIngredientsForRecipe, recipeID)
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
