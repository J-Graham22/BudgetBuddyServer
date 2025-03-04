package controllers

import (
	"encoding/json"
	"net/http"
  "fmt"
  "context"

	"github.com/J-Graham22/BudgetBuddyServer/src/db"
	"github.com/J-Graham22/BudgetBuddyServer/src/db/repository"
	"golang.org/x/crypto/bcrypt"
)

func AddUser(w http.ResponseWriter, r *http.Request) {
  decoder := json.NewDecoder(r.Body)
  var user repository.User
  err := decoder.Decode(&user)
  if err != nil {
    w.WriteHeader(http.StatusBadRequest)
    w.Write([]byte("could not deserialize json body into user"))
    return
  }

  saltedAndHashedPassword, err := bcrypt.GenerateFromPassword(user.Password, bcrypt.DefaultCost)
  if err != nil {
    w.WriteHeader(http.StatusBadRequest)
    w.Write([]byte(fmt.Sprintf("Unable to process password - error: %s", err)))
    return
  }

  ctx := context.Background()
  ctx, dbConn, err := db.PrepareContext()
  if err != nil {
    w.WriteHeader(http.StatusInternalServerError)
    w.Write([]byte(fmt.Sprintf("Unable to open database - error: %s", err)))
    return
  }
  defer dbConn.Close(ctx)

  repo := repository.New(dbConn)
  err = repo.AddUser(
    ctx, 
    repository.AddUserParams{
      Name: user.Name,
      Email: user.Email,
      Password: []byte(saltedAndHashedPassword),
    },
  )

  if err != nil {
    w.WriteHeader(http.StatusInternalServerError)
    w.Write([]byte(fmt.Sprintf("Unable to add user - error: %s", err)))
    return
  }

  w.WriteHeader(http.StatusOK)
  w.Write([]byte(fmt.Sprintf("Added user %s with email %s", user.Name, user.Email)))
}
