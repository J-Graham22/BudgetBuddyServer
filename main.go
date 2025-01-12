package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	_ "time"

	_ "github.com/gin-contrib/cors"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"

	"github.com/J-Graham22/BudgetBuddyServer/src/controllers"
	_ "github.com/J-Graham22/BudgetBuddyServer/src/db"
	"github.com/J-Graham22/BudgetBuddyServer/src/db/repository"
)

func main() {
  ctx := context.Background()

	log.Println("Hello World!")

  var dbConn *sql.DB
  var err error

  dsn := fmt.Sprintf("%s:%s@tcp(localhost:3306)/local_kkb_db?parseTime=true", os.Getenv("DBUSER"), os.Getenv("DBPASS"))

  dbConn, err = sql.Open("mysql", dsn)
  if err != nil {
    log.Fatal(err)
  }
  defer dbConn.Close()

  pingErr := dbConn.Ping()
  if pingErr != nil {
    log.Fatal(pingErr)
  }

  log.Println("Connected")


  repo := repository.New(dbConn)
  transactions, err := repo.GetAllTransactions(ctx)
  if err != nil {
    log.Fatal(err)
  }

  log.Printf("length of transactions = %d", len(transactions))
  if len(transactions) > 0 {
    log.Println("YAHOO!")
  }

  router := http.NewServeMux()
  /*router.HandleFunc("POST /item{id}", func(w http.ResponseWriter, r *http.Request) {
    id := r.PathValue("id")
    w.Write([]byte("received request for item: " + id))
  })*/

  router.HandleFunc("GET /transactions", controllers.GetAllTransactions)

  server := http.Server{
    Addr: ":8080",
    Handler: router,
  }

  log.Println("Starting server on port :8080")
  server.ListenAndServe()
}
