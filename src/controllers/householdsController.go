package controllers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"

	"github.com/J-Graham22/BudgetBuddyServer/src/db"
)

type HouseholdsResponse struct {
	Households []db.Household `json:"households"`
}

func GetHouseholds(c *gin.Context) {
	dbConn, err := db.ExtractDBFromContext(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	var households []db.Household
	result := dbConn.Find(&households)

	fmt.Println(result.Error);
	fmt.Println(households)
	

	if(result.Error != nil) {
		fmt.Println(result.Error)
		return
	}

	var householdResponse HouseholdsResponse = HouseholdsResponse{Households: households}

	c.IndentedJSON(http.StatusOK, householdResponse)
}
