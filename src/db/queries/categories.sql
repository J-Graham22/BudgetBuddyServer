-- name: AddCategory :exec
insert into Categories (name, description, household_id, parent_category_id)
values ($1, $2, $3, $4);
