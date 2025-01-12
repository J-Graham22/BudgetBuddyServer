-- name: GetAllBudgets :many
select * from Budgets 

-- name: AddBudget :exec
insert into Budgets (name, start_date, end_date, household_id)
values ($1, $2, $3, $4);
