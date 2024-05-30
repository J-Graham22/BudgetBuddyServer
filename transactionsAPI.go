package main

import (
	"fmt"
	"net/http"
	"time"

	"database/sql"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

type Transaction struct {
	ID              string    `json:"id"`
	Description     string    `json:"description"`
	Amount          float64   `json:"amount"`
	TransactionTime time.Time `json:"transaction_time"`
	SubcategoryID	string	`json:"subcategory_id"`
}

type TransactionResponse struct {
	Transactions []Transaction `json:"transactions"`
}

//! Endpoints
func getTransactions(c *gin.Context) {
	dbConn, err := extractDBFromContext(c)
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

func addTransaction(c *gin.Context) {
	dbConn, err := extractDBFromContext(c)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	var newTransaction Transaction 

	if err := c.BindJSON(&newTransaction); err != nil {
		fmt.Println(err)
		return
	}

	tx, err := dbConn.BeginTx(c, nil)
	if err != nil {
		panic(err)
	}

	defer tx.Rollback()

	// statement, err := dbConn.Prepare("insert into transactions (amount, description, transaction_time) values ($1, $2, NOW())")
	// if err != nil { fmt.Println(err.Error()) }

	// result, err := statement.Exec(newTransaction.Amount, newTransaction.Description)

	result, err := tx.ExecContext(c, "insert into transactions (amount, description, transaction_time) values ($1, $2, NOW()) returning id",
		newTransaction.Amount, newTransaction.Description)

	if err != nil {
		panic(err)
	}

	fmt.Println(result)

	// orderID, err := result.LastInsertId()
	// if err != nil {
	// 	panic(err)
	// }

	if err = tx.Commit(); err != nil {
		panic(err)
	}

	//fmt.Println(orderID)

	c.IndentedJSON(http.StatusCreated, newTransaction)
}

func getTransactionsByBudget(c *gin.Context) {
	//var budgetId string = c.Params.ByName("budget_id")


}

//! local methods
func scanForTransactions(rows *sql.Rows) TransactionResponse {

	var transactions []Transaction

	if err := rows.Err(); err != nil {
		panic(err)
	}

	for rows.Next() {
		var t Transaction 
		if err := rows.Scan(&t.ID, &t.Description, &t.Amount, &t.TransactionTime); err != nil {
			panic(err)
		}
		transactions = append(transactions, t)
	}

	var transactionsResponse TransactionResponse = TransactionResponse{Transactions: transactions}

	return transactionsResponse
}
