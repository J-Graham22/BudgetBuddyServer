-- name: AddUserHouseholdPair :exec
insert into UserHouseholdMapping(household_id, user_id)
values (sqlc.arg(household_id), sqlc.arg(user_id));
