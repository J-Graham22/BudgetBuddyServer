-- name: AddCategory :exec
insert into Categories (name, description, household_id, parent_category_id)
values (sqlc.arg(name), sqlc.arg(description), sqlc.arg(household_id), sqlc.arg(parent_category_id));
