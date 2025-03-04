// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: categories.sql

package repository

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

const addCategory = `-- name: AddCategory :exec
insert into Categories (name, description, household_id, parent_category_id)
values ($1, $2, $3, $4)
`

type AddCategoryParams struct {
	Name             string      `json:"name"`
	Description      pgtype.Text `json:"description"`
	HouseholdID      int32       `json:"household_id"`
	ParentCategoryID pgtype.Int4 `json:"parent_category_id"`
}

func (q *Queries) AddCategory(ctx context.Context, arg AddCategoryParams) error {
	_, err := q.db.Exec(ctx, addCategory,
		arg.Name,
		arg.Description,
		arg.HouseholdID,
		arg.ParentCategoryID,
	)
	return err
}
