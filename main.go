package main

import (
	"fmt"
	"net/http"
	"time"

	"database/sql"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

type album struct {
	ID	string	`json:"id"`
	Title	string	`json:"title"`
	Artist	string	`json:"artist"`
	Price	float64	`json:"price"`
}

type transaction struct {
	ID string
	description string
	amount float64
	transaction_time time.Time
}

var albums = []album {
	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
    {ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
    {ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
}

func main() {
	fmt.Println("Hello World!");

	connStr := "postgresql://budget-db_owner:gI4jtl8CDbhB@ep-green-dream-a485fkmn.us-east-1.aws.neon.tech/budget-db?sslmode=require"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	rows, err := db.Query("select * from transactions")
	if err != nil {
		panic(err)
	}

	var transactions []transaction

	for rows.Next() {
		var t transaction
		if err := rows.Scan(&t.ID, &t.description, &t.amount, &t.transaction_time); err != nil {
			panic(err)
		}
		transactions = append(transactions, t)
	}

	if err = rows.Err(); err != nil {
		panic(err)
	}
	for i := 0; i < len(transactions); i++ {
		fmt.Print(transactions[i])
	}

	var version string
	if err := db.QueryRow("select version()").Scan(&version); err != nil {
		panic(err)
	}

	fmt.Printf("version=%s\n", version)

	return

	router := gin.Default();
	router.GET("/albums", getAlbums);
	router.GET("/albums/:id", getAlbumById)
	router.POST("/albums", postAlbums);

	router.Run("localhost:8080");
}

func getAlbums(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, albums);
}

func getAlbumById(c *gin.Context) {
	id := c.Param("id")

	for _, a := range albums {
		if a.ID == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}

	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
}

func postAlbums(c *gin.Context) {
	var newAlbum album;

	if err := c.BindJSON(&newAlbum); err != nil {
		fmt.Println(err)
		return
	}

	albums = append(albums, newAlbum)
	c.IndentedJSON(http.StatusCreated, newAlbum)
}