-- name: GetAllBudgets :many
select * from Budgets;

-- name: AddBudget :exec
insert into Budgets(name, start_date, end_date, household_id)
values (sqlc.arg(name), sqlc.arg(start_date), sqlc.arg(end_date), sqlc.arg(household_id));
