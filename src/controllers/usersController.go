package controllers

import (
	"github.com/J-Graham22/BudgetBuddyServer/src/db"
)

type UserResponse struct {
	Users []db.User `json:"users"`
}