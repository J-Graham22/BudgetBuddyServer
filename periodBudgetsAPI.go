package main

import (
	"fmt"
	"net/http"
	"time"

	"database/sql"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

type PeriodBudget struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	StartDate time.Time `json:"start_date"`
	EndDate   time.Time `json:"end_date"`
}

type BudgetResponse struct {
	Budgets []PeriodBudget `json:"budgets"`
}

func getPeriodBudgets(c *gin.Context) {
	dbConn, err := extractDBFromContext(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	rows, err := dbConn.Query("select * from period_budget")
	if err != nil {
		panic(err)
	}

	var budgetResponse BudgetResponse = scanForPeriodBudgets(rows)
	c.IndentedJSON(http.StatusOK, budgetResponse)
}

func getPeriodBudgetById(c *gin.Context) {
	id := c.Param("id")

	dbConn, err := extractDBFromContext(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	rows, err := dbConn.Query("select * from period_budget where id = $1", id)
	if err != nil {
		panic(err)
	}

	var budgetResponse BudgetResponse = scanForPeriodBudgets(rows)
	c.IndentedJSON(http.StatusOK, budgetResponse)

	var pb PeriodBudget 

	for rows.Next() {
		if err := rows.Scan(&pb.ID, &pb.StartDate, &pb.EndDate); err != nil {
			c.IndentedJSON(http.StatusNotFound, gin.H{"message": "budget not found"})
			panic(err)
		}
	}

	c.IndentedJSON(http.StatusOK, pb)
}

func addPeriodBudget(c *gin.Context) {
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

	var newPeriodBudget period_budget

	if err := c.BindJSON(&newPeriodBudget); err != nil {
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

	result, err := tx.ExecContext(c, "insert into period_budget (name, start_date, end_date) values ($1, $2, $3) returning id",
		newPeriodBudget.Name, newPeriodBudget.StartDate, newPeriodBudget.EndDate)

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

	c.IndentedJSON(http.StatusCreated, newPeriodBudget)
}

//!local methods
func scanForPeriodBudgets(rows *sql.Rows) BudgetResponse {

	var periodBudgets []PeriodBudget

	if err := rows.Err(); err != nil {
		panic(err)
	}

	for rows.Next() {
		var pb PeriodBudget 
		if err := rows.Scan(&pb.ID, &pb.Name, &pb.StartDate, &pb.EndDate); err != nil {
			panic(err)
		}
		periodBudgets = append(periodBudgets, pb)
	}

	var budgetResponse BudgetResponse = BudgetResponse{Budgets: periodBudgets}

	return budgetResponse 
}
