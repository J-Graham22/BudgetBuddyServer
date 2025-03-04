package main

import (
	"log"
	"net/http"
	_ "time"

  "github.com/rs/cors"
	_ "github.com/gin-contrib/cors"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"

	"github.com/J-Graham22/BudgetBuddyServer/src/controllers"
	_ "github.com/J-Graham22/BudgetBuddyServer/src/db"
	_ "github.com/J-Graham22/BudgetBuddyServer/src/db/repository"
)

func main() {
	log.Println("Hello World!")

  router := http.NewServeMux()


  //households
  router.HandleFunc("POST /households/{name}", controllers.AddHousehold)

  //users
  router.HandleFunc("POST /users", controllers.AddUser)

  //categories
  router.HandleFunc("POST /categories/{id}", controllers.AddCategory)
  router.HandleFunc("GET /categories/{id}", controllers.GetCategoriesForHousehold)

  //transactions
  router.HandleFunc("GET /transactions/{id}", controllers.GetTransactionsForHousehold)

  _cors := cors.Options{
      AllowedMethods: []string{"POST", "OPTIONS"},
      AllowedOrigins: []string{"http://localhost:8080"},
  }
  handler := cors.New(_cors).Handler(router)

  server := http.Server{
    Addr: ":8080",
    Handler: handler,
  }

  log.Println("Starting server on port :8080")
  server.ListenAndServe()
}
