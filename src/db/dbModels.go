package db

import (
	"time"

	"gorm.io/gorm"
)

type Transaction struct {
	gorm.Model
	Description     string    `json:"description"`
	Amount          float64   `json:"amount"`
	TransactionTime time.Time `json:"transaction_time"`
	SubcategoryID   string    `json:"subcategory_id"`
}

type PeriodBudget struct {
	gorm.Model
	Name      string    `json:"name"`
	StartDate time.Time `json:"start_date"`
	EndDate   time.Time `json:"end_date"`
}

type Household struct {
	gorm.Model
	Name string `json:"name"`
}

type User struct {
	gorm.Model
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	HouseholdIds	[]int	`json:"household_ids"`
}
