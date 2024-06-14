package db

import (
	"errors"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func ExtractDBFromContext(c *gin.Context) (*gorm.DB, error) {
	db, exists := c.Get("db")
	if !exists {
		err := errors.New("database connection not found")
		return nil, err
	}

	// Convert the interface{} type to *sql.DB
	dbConn, ok := db.(*gorm.DB)
	if !ok {
		err := errors.New("failed to convert database connection to *gorm.DB")
		return nil, err
	}

	return dbConn, nil
}