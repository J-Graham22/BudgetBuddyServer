package db

import (
	"database/sql"
	"errors"

	"github.com/gin-gonic/gin"
)

func ExtractDBFromContext(c *gin.Context) (*sql.DB, error) {
	db, exists := c.Get("db")
	if !exists {
		err := errors.New("database connection not found")
		return nil, err
	}

	// Convert the interface{} type to *sql.DB
	dbConn, ok := db.(*sql.DB)
	if !ok {
		err := errors.New("failed to convert database connection")
		return nil, err
	}

	return dbConn, nil
}