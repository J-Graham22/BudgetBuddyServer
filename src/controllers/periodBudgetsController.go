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

	fmt.Println(result.Error);
	fmt.Println(periodBudgets)
	

	if(result.Error != nil) {
		fmt.Println(result.Error)
		return
	}

	var budgetResponse BudgetResponse = BudgetResponse{Budgets: periodBudgets}

	c.IndentedJSON(http.StatusOK, budgetResponse)
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

	fmt.Println(pb)

	result := dbConn.Create(&pb)

	if (result.Error != nil) {
		fmt.Println(result.Error)
		return
	}	

	c.IndentedJSON(http.StatusCreated, pb)
}
