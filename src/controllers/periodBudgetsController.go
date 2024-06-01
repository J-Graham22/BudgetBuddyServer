package controllers

import (
	"fmt"
	"net/http"

	"database/sql"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"

	"github.com/J-Graham22/BudgetBuddyServer/src/db"
)

type BudgetResponse struct {
	Budgets []db.PeriodBudget `json:"budgets"`
}

func GetPeriodBudgets(c *gin.Context) {
	dbConn, err := db.ExtractDBFromContext(c)
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

func GetPeriodBudgetById(c *gin.Context) {
	id := c.Param("id")

	dbConn, err := db.ExtractDBFromContext(c)
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

	var pb db.PeriodBudget 

	for rows.Next() {
		if err := rows.Scan(&pb.ID, &pb.StartDate, &pb.EndDate); err != nil {
			c.IndentedJSON(http.StatusNotFound, gin.H{"message": "budget not found"})
			panic(err)
		}
	}

	c.IndentedJSON(http.StatusOK, pb)
}

func AddPeriodBudget(c *gin.Context) {
	dbConn, err := db.ExtractDBFromContext(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	var pb db.PeriodBudget

	if err := c.BindJSON(&pb); err != nil {
		fmt.Println(err)
		return
	}

	tx, err := dbConn.BeginTx(c, nil)
	if err != nil {
		panic(err)
	}

	defer tx.Rollback()

	result, err := tx.ExecContext(c, "insert into period_budget (name, start_date, end_date) values ($1, $2, $3) returning id",
		pb.Name, pb.StartDate, pb.EndDate)

	if err != nil {
		panic(err)
	}

	fmt.Println(result)

	if err = tx.Commit(); err != nil {
		panic(err)
	}

	c.IndentedJSON(http.StatusCreated, pb)
}

//!local methods
func scanForPeriodBudgets(rows *sql.Rows) BudgetResponse {

	var periodBudgets []db.PeriodBudget

	if err := rows.Err(); err != nil {
		panic(err)
	}

	for rows.Next() {
		var pb db.PeriodBudget 
		if err := rows.Scan(&pb.ID, &pb.Name, &pb.StartDate, &pb.EndDate); err != nil {
			panic(err)
		}
		periodBudgets = append(periodBudgets, pb)
	}

	var budgetResponse BudgetResponse = BudgetResponse{Budgets: periodBudgets}

	return budgetResponse 
}
