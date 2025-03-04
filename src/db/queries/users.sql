-- name: GetAllUsers :many
select * from Users;

-- name: AddUser :exec
insert into Users(username, name, email, password)
values (sqlc.arg(username), sqlc.arg(name), sqlc.arg(email), sqlc.arg(password));
