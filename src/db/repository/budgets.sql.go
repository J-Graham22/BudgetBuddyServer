// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: budgets.sql

package repository

import (
	"context"
)

const getAllBudgets = `-- name: GetAllBudgets :many
select id, name, start_date, end_date, household_id from Budgets
`

func (q *Queries) GetAllBudgets(ctx context.Context) ([]Budget, error) {
	rows, err := q.db.QueryContext(ctx, getAllBudgets)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Budget
	for rows.Next() {
		var i Budget
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.StartDate,
			&i.EndDate,
			&i.HouseholdID,
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
