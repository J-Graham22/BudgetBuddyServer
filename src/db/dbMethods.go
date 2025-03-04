package db

import (
  "context"
  "log"
  "os"

  "github.com/jackc/pgx/v5"
)

func PrepareContext() (context.Context, *pgx.Conn, error) {
  ctx := context.Background()

  dbConn, err := pgx.Connect(context.Background(), os.Getenv("DB_URL"))
  if err != nil {
    os.Exit(1)
  }

  log.Println("Connected")
  
  return ctx, dbConn, nil 
}
