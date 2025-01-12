package controllers

import (
	"log"
	"net/http"
  "context"
  "fmt"
  "database/sql"
  "os"

	_ "github.com/J-Graham22/BudgetBuddyServer/src/db"
	"github.com/J-Graham22/BudgetBuddyServer/src/db/repository"
)

func PrepareContext() (context.Context, *sql.DB, error) {
  ctx := context.Background()

  var dbConn *sql.DB
  var err error

  dsn := fmt.Sprintf("%s:%s@tcp(localhost:3306)/local_kkb_db?parseTime=true", os.Getenv("DBUSER"), os.Getenv("DBPASS"))

  dbConn, err = sql.Open("mysql", dsn)
  if err != nil {
    log.Fatal(err)
    return nil, nil, err
  }

  pingErr := dbConn.Ping()
  if pingErr != nil {
    log.Fatal(pingErr)
    return nil, nil, err
  }

  log.Println("Connected")
  
  return ctx, dbConn, nil 
}

func GetAllTransactions(w http.ResponseWriter, r *http.Request) {
  log.Println("executing GetAllTransactions")

  ctx := context.Background()
  ctx, dbConn, err := PrepareContext()
  if err != nil {
    log.Fatal(err)
  }
  defer dbConn.Close()

  repo := repository.New(dbConn)
  transactions, err := repo.GetAllTransactions(ctx)
  if err != nil {
    log.Fatal(err)
  }

  log.Printf("length of transactions = %d", len(transactions))
  if len(transactions) > 0 {
    log.Println("YAHOO!")
  }
}

func AddTransaction() {
}

func UpdateTransaction() {
}

func GetTransactionsByBudget() {
}

func GetTransactionsForHousehold() {

}

func GetTransactionsByCategory() {

}

func GetTransactionsByUser() {

}
