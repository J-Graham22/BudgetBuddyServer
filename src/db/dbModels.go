package db

import (
	"time"
)

type Transaction struct {
	ID              int       `json:"id"`
	Description     string    `json:"description"`
	Amount          float64   `json:"amount"`
	TransactionTime time.Time `json:"transaction_time"`
	SubcategoryID   string    `json:"subcategory_id"`
}

type PeriodBudget struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	StartDate time.Time `json:"start_date"`
	EndDate   time.Time `json:"end_date"`
}

type Household struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type User struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	Password  string `json:"password"`
}
