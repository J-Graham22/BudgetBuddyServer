package controllers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"

	"github.com/J-Graham22/BudgetBuddyServer/src/db"
)

type TransactionResponse struct {
	Transactions []db.Transaction `json:"transactions"`
}

func GetTransactions(c *gin.Context) {
	dbConn, err := db.ExtractDBFromContext(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	var transactions []db.Transaction
	result := dbConn.Find(&transactions)

	if result.Error != nil {
		fmt.Println(result.Error)
		return
	}

	var transactionsResponse TransactionResponse = TransactionResponse{Transactions: transactions}

	c.IndentedJSON(http.StatusOK, transactionsResponse)
}

func AddTransaction(c *gin.Context) {
	dbConn, err := db.ExtractDBFromContext(c)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	fmt.Println("wowowowo")

	var newTransaction db.TransactionType

	if err := c.BindJSON(&newTransaction); err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(newTransaction)

	//newTransaction.TransactionTime = newTransaction.TransactionTime.Format("2006-01-02 15:04:05")

	var newTransactionRecord db.Transaction = db.Transaction{Transaction: newTransaction}

	result := dbConn.Create(&newTransactionRecord)

	if result.Error != nil {
		fmt.Println(result.Error)
		return
	}

	c.IndentedJSON(http.StatusCreated, newTransactionRecord)
}

func UpdateTransaction(c *gin.Context) {
	var transactionId string = c.Params.ByName("transaction_id")
	dbConn, err := db.ExtractDBFromContext(c)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	var updatingTransaction db.TransactionType
	var updatingTransactionRecord db.Transaction

	if err := c.BindJSON(&updatingTransaction); err != nil {
		fmt.Println(err)
		return
	}

	if result := dbConn.First(&updatingTransactionRecord, transactionId); result.Error != nil {
		fmt.Println(result.Error)
		return
	}

	updatingTransactionRecord.Transaction = updatingTransaction

	if result := dbConn.Save(&updatingTransaction); result.Error != nil {
		fmt.Println(result.Error)
		return
	}

	c.IndentedJSON(http.StatusOK, updatingTransactionRecord)
}

func GetTransactionsByBudget(c *gin.Context) {
	var budgetId string = c.Params.ByName("budget_id")

	dbConn, err := db.ExtractDBFromContext(c)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	var pb db.PeriodBudget
	result := dbConn.First(&pb, budgetId)

	if result.Error != nil {
		fmt.Println(result.Error)
		return	
	}

	var transactionsForBudget []db.Transaction

	result = dbConn.
        Table("transactions").
        Joins("JOIN period_budgets ON transactions.date BETWEEN period_budgets.start_date AND period_budgets.end_date").
        Where("period_budgets.id = ?", budgetId).
        Find(&transactionsForBudget)

	if result.Error != nil {
		fmt.Println(result.Error)
		return
	}

	c.IndentedJSON(http.StatusOK, transactionsForBudget)
}
