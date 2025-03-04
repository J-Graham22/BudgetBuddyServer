package controllers

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/J-Graham22/BudgetBuddyServer/src/db"
	"github.com/J-Graham22/BudgetBuddyServer/src/db/repository"
)

func AddTransaction() {
}

func UpdateTransaction() {
}

func GetTransactionsByBudget() {
}

func GetTransactionsForHousehold(w http.ResponseWriter, r *http.Request) {
  ctx := context.Background()
  ctx, dbConn, err := db.PrepareContext()
  if err != nil {
    w.WriteHeader(http.StatusInternalServerError)
    w.Write([]byte(fmt.Sprint("Encountered error while trying to open database - error: %s", err)))
    return
  }
  defer dbConn.Close(ctx)

  repo := repository.New(dbConn)
  transactions, err := repo.GetTransactionsByHousehold(ctx, 0)
  if err != nil {
    w.WriteHeader(http.StatusInternalServerError)
    w.Write([]byte(fmt.Sprint("Encountered error while retrieving transactions - error: %s", err)))
    return
  }

  transactionsJson, err := json.Marshal(transactions)
  if err != nil {
    w.WriteHeader(http.StatusInternalServerError)
    w.Write([]byte(fmt.Sprint("Encountered error while serializing transactions into json - error: %s", err)))
    return
  }
  w.WriteHeader(http.StatusOK)
  w.Write(transactionsJson)
}

func GetTransactionsByCategory() {

}

func GetTransactionsByUser() {

}
