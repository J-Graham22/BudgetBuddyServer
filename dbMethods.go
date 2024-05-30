package main

import (
	"database/sql"
	"errors"

	"github.com/gin-gonic/gin"
)

func extractDBFromContext(c *gin.Context) (*sql.DB, error) {
	db, exists := c.Get("db")
	if !exists {
		err := errors.New("Database connection not found")
		return nil, err
	}

	// Convert the interface{} type to *sql.DB
	dbConn, ok := db.(*sql.DB)
	if !ok {
		err := errors.New("Failed to convert database connection")
		return nil, err
	}

	return dbConn, nil
}