-- name: GetAllTransactions :many
select * from Transactions;

-- name: GetTransactionsByHousehold :many
select * from Transactions
where household_id = sqlc.arg(household_id); 
