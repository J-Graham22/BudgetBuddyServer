-- name: AddUserHouseholdPair :exec
insert into UserHouseholdMappings(household_id, user_id)
values ($1, $2);
