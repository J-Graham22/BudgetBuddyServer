package main

import (
	"fmt"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"

	"github.com/J-Graham22/BudgetBuddyServer/src/controllers"
	"github.com/J-Graham22/BudgetBuddyServer/src/db"
)

type App struct {
	DB *gorm.DB
}

func main() {
	fmt.Println("Hello World!")
	/*
		connStr := "postgresql://budget-db_owner:gI4jtl8CDbhB@ep-green-dream-a485fkmn.us-east-1.aws.neon.tech/budget-db?sslmode=require"
		db, err := sql.Open("postgres", connStr)
		if err != nil {
			panic(err)
		}
		defer db.Close()
	*/
	dsn := "postgresql://budget-db_owner:gI4jtl8CDbhB@ep-green-dream-a485fkmn.us-east-1.aws.neon.tech/budget-db?sslmode=require"
	gormDB, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	sqlDB, err := gormDB.DB()
	if err != nil {
		panic(err)
	}
	defer sqlDB.Close()

	gormDB.AutoMigrate(&db.Transaction{})
	gormDB.AutoMigrate(&db.PeriodBudget{})
	gormDB.AutoMigrate(&db.Household{})
	gormDB.AutoMigrate(&db.User{})

	router := gin.Default()

	app := &App{DB: gormDB}

	router.Use(func(c *gin.Context) {
		c.Set("db", app.DB)
		c.Next()
	})

	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173"},
		AllowMethods:     []string{"GET", "POST"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		AllowOriginFunc: func(origin string) bool {
			return origin == "https://github.com"
		},
		MaxAge: 12 * time.Hour,
	}))

	router.GET("/transactions", controllers.GetTransactions)
	router.POST("/transaction", controllers.AddTransaction)
	router.GET("/budgets", controllers.GetPeriodBudgets)
	router.GET("/budgets/:id", controllers.GetPeriodBudgetById)
	router.POST("/budget", controllers.AddPeriodBudget)

	router.Run("localhost:8080")
}
