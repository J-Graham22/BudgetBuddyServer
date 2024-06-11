package controllers

import (
	"fmt"
	"net/http"

	"database/sql"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"

	"github.com/J-Graham22/BudgetBuddyServer/src/db"
)

type TransactionResponse struct {
	Transactions []db.Transaction `json:"transactions"`
}

// ! Endpoints
func GetTransactions(c *gin.Context) {
	dbConn, err := db.ExtractDBFromContext(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	rows, err := dbConn.Query("select * from transactions")
	if err != nil {
		panic(err)
	}

	var transactionsResponse = scanForTransactions(rows)

	c.IndentedJSON(http.StatusOK, transactionsResponse)
}

func AddTransaction(c *gin.Context) {
	dbConn, err := db.ExtractDBFromContext(c)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	var newTransaction db.Transaction

	if err := c.BindJSON(&newTransaction); err != nil {
		fmt.Println(err)
		return
	}

	tx, err := dbConn.BeginTx(c, nil)
	if err != nil {
		panic(err)
	}

	defer tx.Rollback()

	result, err := tx.ExecContext(c, "insert into transactions (amount, description, transaction_time) values ($1, $2, NOW()) returning id",
		newTransaction.Amount, newTransaction.Description)

	if err != nil {
		panic(err)
	}

	fmt.Println(result)

	if err = tx.Commit(); err != nil {
		panic(err)
	}

	c.IndentedJSON(http.StatusCreated, newTransaction)
}

func GetTransactionsByBudget(c *gin.Context) {
	//var budgetId string = c.Params.ByName("budget_id")

}

// ! local methods
func scanForTransactions(rows *sql.Rows) TransactionResponse {

	var transactions []db.Transaction

	if err := rows.Err(); err != nil {
		panic(err)
	}

	for rows.Next() {
		var t db.Transaction
		if err := rows.Scan(&t.ID, &t.Description, &t.Amount, &t.TransactionTime); err != nil {
			panic(err)
		}
		transactions = append(transactions, t)
	}

	var transactionsResponse TransactionResponse = TransactionResponse{Transactions: transactions}

	return transactionsResponse
}
