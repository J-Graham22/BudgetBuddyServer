package main

import (
	"fmt"
	"net/http"
	"time"

	"database/sql"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

type period_budget struct {
	ID	string	`json:"id"`
	StartDate time.Time `json:"start_date"`
	EndDate time.Time `json:"end_date"`

}

func getPeriodBudgets(c *gin.Context) {
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

	rows, err := dbConn.Query("select * from period_budget")
	if err != nil {
		panic(err)
	}

	var period_budgets []period_budget

	for rows.Next() {
		var pb period_budget
		if err := rows.Scan(&pb.ID, &pb.StartDate, &pb.EndDate); err != nil {
			panic(err)
		}
		period_budgets = append(period_budgets, pb)
		fmt.Println(pb)
	}

	if err = rows.Err(); err != nil {
		panic(err)
	}

	c.IndentedJSON(http.StatusOK, period_budgets)
}

func getPeriodBudgetById(c *gin.Context) {
	id := c.Param("id")

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

	rows, err := dbConn.Query("select * from period_budget where id = $1", id)
	if err != nil {
		panic(err)
	}

	var pb period_budget

	for rows.Next() {
		if err := rows.Scan(&pb.ID, &pb.StartDate, &pb.EndDate); err != nil {
			c.IndentedJSON(http.StatusNotFound, gin.H{"message": "budget not found"})
			panic(err)
		}
	}

	c.IndentedJSON(http.StatusOK, pb)
}