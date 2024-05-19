package main

import (
	"fmt"
	"time"

	"database/sql"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

type App struct {
	DB *sql.DB
}

func main() {
	fmt.Println("Hello World!")

	connStr := "postgresql://budget-db_owner:gI4jtl8CDbhB@ep-green-dream-a485fkmn.us-east-1.aws.neon.tech/budget-db?sslmode=require"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	router := gin.Default()

	app := &App{DB: db}

	router.Use(func(c *gin.Context) {
		c.Set("db", app.DB)
		c.Next()
	})

	router.Use(cors.New(cors.Config {
		AllowOrigins: []string{"http://localhost:5173"},
		AllowMethods: []string{"GET", "POST"},
		AllowHeaders: []string{"Origin", "Content-Type", "Accept"},
		ExposeHeaders: []string{"Content-Length"},
		AllowCredentials: true,
		AllowOriginFunc: func(origin string) bool {
			return origin == "https://github.com"
		},
		MaxAge: 12 * time.Hour,
	}))
	
	router.GET("/transactions", getTransactions)
	router.POST("/transaction", addTransaction)
	router.GET("/budgets", getPeriodBudgets)
	router.GET("/budgets/:id", getPeriodBudgetById)

	// var version string
	// if err := db.QueryRow("select version()").Scan(&version); err != nil {
	// 	panic(err)
	// }

	// fmt.Printf("version=%s\n", version)

	router.Run("localhost:8080")
}
