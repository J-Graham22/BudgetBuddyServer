package main

import (
	"fmt"
	"net/http"
	"time"

	"database/sql"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

type transaction struct {
	ID              string    `json:"id"`
	Description     string    `json:"description"`
	Amount          float64   `json:"amount"`
	TransactionTime time.Time `json:"transaction_time"`
}

func getTransactions(c *gin.Context) {
	// Retrieve the database connection from the context
	db, exists := c.Get("db")
	if !exists {
		c.JSON(500, gin.H{"error": "Database connection not found"})
		return
	}

	// Convert the interface{} type to *sql.DB
	dbConn, ok := db.(*sql.DB)
	if !ok {
		c.JSON(500, gin.H{"error": "Failed to convert database connection"})
		return
	}

	rows, err := dbConn.Query("select * from transactions")
	if err != nil {
		panic(err)
	}

	var transactions []transaction

	for rows.Next() {
		var t transaction
		if err := rows.Scan(&t.ID, &t.Description, &t.Amount, &t.TransactionTime); err != nil {
			panic(err)
		}
		transactions = append(transactions, t)
		fmt.Println(t)
	}

	if err = rows.Err(); err != nil {
		panic(err)
	}

	c.IndentedJSON(http.StatusOK, transactions)
}

func addTransaction(c *gin.Context) {
	db, exists := c.Get("db")
	if !exists {
		c.JSON(500, gin.H{"error": "Database connection not found"})
		return
	}

	// Convert the interface{} type to *sql.DB
	dbConn, ok := db.(*sql.DB)
	if !ok {
		c.JSON(500, gin.H{"error": "Failed to convert database connection"})
		return
	}

	var newTransaction transaction

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
