package controllers

import (
	"fmt"
	"net/http"

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

	var periodBudgets []db.PeriodBudget
	result := dbConn.Find(&periodBudgets)

	if(result.Error != nil) {
		fmt.Println(result.Error)
		return
	}

	var budgetResponse BudgetResponse = BudgetResponse{Budgets: periodBudgets}

	c.IndentedJSON(http.StatusOK, budgetResponse)

	/*rows, err := dbConn.Query("select * from period_budgets")
	if err != nil {
		panic(err)
	}

	var budgetResponse BudgetResponse = scanForPeriodBudgets(rows)
	c.IndentedJSON(http.StatusOK, budgetResponse)*/
}

func GetPeriodBudgetById(c *gin.Context) {
	id := c.Param("id")

	dbConn, err := db.ExtractDBFromContext(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	var periodBudget db.PeriodBudget

	result := dbConn.First(&periodBudget, id)
	if result.Error != nil {
		fmt.Println(result.Error)
		return
	}

	var budgetResponse BudgetResponse = BudgetResponse{Budgets: []db.PeriodBudget{periodBudget}}

	c.IndentedJSON(http.StatusOK, budgetResponse)

	/*rows, err := dbConn.Query("select * from period_budgets where id = $1", id)
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

	c.IndentedJSON(http.StatusOK, pb)*/
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

	result := dbConn.Create(&pb)

	if (result.Error != nil) {
		fmt.Println(result.Error)
		return
	}	

	/*tx, err := dbConn.BeginTx(c, nil)
	if err != nil {
		panic(err)
	}

	defer tx.Rollback()

	result, err := tx.ExecContext(c, "insert into period_budgets (name, start_date, end_date) values ($1, $2, $3) returning id",
		pb.Name, pb.StartDate, pb.EndDate)

	if err != nil {
		panic(err)
	}

	fmt.Println(result)

	if err = tx.Commit(); err != nil {
		panic(err)
	}*/

	c.IndentedJSON(http.StatusCreated, pb)
}
