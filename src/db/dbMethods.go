package db

import (
	"database/sql/driver"
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"

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

func (t *TransactionType) Scan(value interface{}) error {
	str, ok := value.(string)
	if !ok {
		return errors.New("type assertion to string failed")
	}

	parts, err := CustomTypeExtraction(str, 4)
	if err != nil {
		return err
	}

	t.TransactionName = RemoveUnnecessaryQuotes(parts[0]) 
	t.Description = RemoveUnnecessaryQuotes(parts[1])

	amount, err := strconv.ParseFloat(parts[2], 64)
	if err != nil {
		return err
	}
	t.Amount = amount

	parts[3] = RemoveUnnecessaryQuotes(parts[3])
	//transactionTime, err := time.Parse("2006-01-02 15:04:05", parts[3])
	transactionTime, err := time.Parse(time.RFC822, parts[3])
	if err != nil {
		return err
	}
	t.TransactionTime = transactionTime

	return nil
}

func (t TransactionType) Value() (driver.Value, error) {
	return fmt.Sprintf("(%s,%s,%f,%s)", t.TransactionName, t.Description, t.Amount, t.TransactionTime), nil
}

func CustomTypeExtraction(typeStr string, fieldNum int) ([]string, error) {
	// Remove parentheses
	typeStr = strings.TrimPrefix(typeStr, "(")
	typeStr = strings.TrimSuffix(typeStr, ")")

	// Split the string by commas
	parts := strings.Split(typeStr, ",")

	if len(parts) != fieldNum {
		return nil, errors.New("invalid data format")
	}

	return parts, nil
}

func RemoveUnnecessaryQuotes(strValue string) string {
	return strings.Trim(strValue, "\"")
}
