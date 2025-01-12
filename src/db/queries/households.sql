-- name: GetAllHouseholds :many
select * from Households 

-- name: AddHousehold :exec
insert into Households(name)
values ($1);
