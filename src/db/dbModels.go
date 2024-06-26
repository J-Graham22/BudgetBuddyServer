package db

import (
	"time"

	"gorm.io/gorm"
)

//! Database Tables
type Transaction struct {
	gorm.Model
	Transaction TransactionType	`gorm:"type:transaction_type"`
}

type PeriodBudget struct {
	gorm.Model
	Name      string    `gorm:"name"`
	StartDate time.Time `gorm:"start_date" json:"start_date"`
	EndDate   time.Time `gorm:"end_date" json:"end_date"`
}

type Household struct {
	gorm.Model
	Name string `gorm:"name"`
}

type User struct {
	gorm.Model
	FirstName string `gorm:"first_name"`
	LastName  string `gorm:"last_name"`
	Email     string `gorm:"email"`
	Password  string `gorm:"password"`
	HouseholdIds	[]int	`gorm:"household_ids"`
}

//! Database Types
type TransactionType struct {
	TransactionName	string	`json:"transaction_name"`
	Description	string	`json:"description"`
	Amount float64	`json:"amount"`
	TransactionTime	time.Time	`json:"transaction_time"`
}