-- name: GetAllUsers :many
select * from Users;

-- name: AddUser :exec
insert into Users(username, name, email)
values ($1, $2, $3);
