package controllers

import (
	"context"
	"fmt"
	"net/http"
	"strconv"

	"github.com/J-Graham22/BudgetBuddyServer/src/db"
	"github.com/J-Graham22/BudgetBuddyServer/src/db/repository"
)

func AddUserHouseholdMapping(w http.ResponseWriter, r *http.Request) {
  householdId := r.URL.Query().Get("household_id")
  if householdId == "" {
    w.WriteHeader(http.StatusBadRequest)
    w.Write([]byte("household_id not set"))
    return
  }
  householdIdInt, err := strconv.ParseInt(householdId, 10, 64)
  if err != nil {
    w.WriteHeader(http.StatusBadRequest)
    w.Write([]byte("could not parse household_id into int"))
    return
  }

  userId := r.URL.Query().Get("user_id")
  if userId == "" {
    w.WriteHeader(http.StatusBadRequest)
    w.Write([]byte("user_id not set"))
    return
  }
  userIdInt, err := strconv.ParseInt(userId, 10, 64)
  if err != nil {
    w.WriteHeader(http.StatusBadRequest)
    w.Write([]byte("could not parse user_id into int"))
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
  err = repo.AddUserHouseholdPair(
    ctx, 
    repository.AddUserHouseholdPairParams{
      HouseholdID: int32(householdIdInt),
      UserID: int32(userIdInt)},
  )

  if err != nil {
    w.WriteHeader(http.StatusInternalServerError)
    w.Write([]byte(fmt.Sprintf("Encountered exception while attempting to - error: %s", err)))
    return
  }

  w.WriteHeader(http.StatusOK)
  w.Write([]byte(fmt.Sprintf("Sucessfully mapped household %d to user %d", householdIdInt, userIdInt)))
}

// READ 
func GetUsersByHouseholdId(w http.ResponseWriter, r *http.Request) {
  //TODO
}
func GetHouseholdsByUserId(w http.ResponseWriter, r *http.Request) {
  //TODO
}

// UPDATE is unnecessary here

// DELETE
// use to remove a user from a household
func DeleteUserHouseholdMapping(w http.ResponseWriter, r *http.Request) {
  //TODO
}
// use when deleting a household
