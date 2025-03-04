-- name: GetAllHouseholds :many
select * from Households;

-- name: AddHousehold :one
insert into Households (name)
values (sqlc.arg(name)) returning *;
