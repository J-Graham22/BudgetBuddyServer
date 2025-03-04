package controllers

import (
	"context"
	"encoding/json"
	_ "errors"
	"fmt"
	_ "fmt"
	"net/http"

	"github.com/J-Graham22/BudgetBuddyServer/src/db"
	"github.com/J-Graham22/BudgetBuddyServer/src/db/repository"
	"github.com/jackc/pgx/v5/pgtype"
)

// CREATE
func GetHouseholds(w http.ResponseWriter, r *http.Request) {
  ctx := context.Background()
  ctx, dbConn, err := db.PrepareContext()
  if err != nil {
    w.WriteHeader(http.StatusInternalServerError)
    w.Write([]byte(fmt.Sprintf("Unable to open database - err: %s", err)))
    return
  }
  defer dbConn.Close(ctx)

  repo := repository.New(dbConn)
  households, err := repo.GetAllHouseholds(ctx)
  if err != nil {
    w.WriteHeader(http.StatusInternalServerError)
    w.Write([]byte(fmt.Sprintf("Encountered error when getting households - err: %s", err)))
    return
  }

  w.WriteHeader(http.StatusOK)
  householdsJson, err := json.Marshal(households)
  if err != nil {
    w.WriteHeader(http.StatusInternalServerError)
    w.Write([]byte(fmt.Sprintf("Encountered error when serializing households into JSON - err: %s", err)))
    return
  }
  w.Write([]byte(householdsJson))
}

func AddHousehold(w http.ResponseWriter, r *http.Request) {
  householdName := r.PathValue("name")

  if householdName == "" {
    w.WriteHeader(http.StatusBadRequest)
    w.Write([]byte("Name of new household not provided in request"))
    return
  }

  ctx := context.Background()
  ctx, dbConn, err := db.PrepareContext()
  if err != nil {
    w.WriteHeader(http.StatusInternalServerError)
    w.Write([]byte(fmt.Sprintf("Unable to open database - err: %s", err)))
    return
  }
  defer dbConn.Close(ctx)

  repo := repository.New(dbConn)
  household, err := repo.AddHousehold(ctx, householdName)
  if err != nil {
    w.WriteHeader(http.StatusInternalServerError)
    w.Write([]byte(fmt.Sprintf("Encountered error when adding household - err: %s", err)))
    return
  }

  //add the default categories
  var categoryDescription pgtype.Text
  //1. Need 
  categoryDescription.Scan("Use this category for any required expenses that arise, such as groceries and gas")
  err = repo.AddCategory(
    ctx,
    repository.AddCategoryParams{
      Name: "Needs",  
      Description: categoryDescription,
      HouseholdID: household.ID,
    })
  if err != nil {
    w.WriteHeader(http.StatusBadRequest)
    w.Write([]byte(fmt.Sprintf("Encountered error when adding default categories - err: %s", err)))
    return
  }

  //2. Wants
  categoryDescription.Scan("Use this category for any discretionary expenses that are for enjoyment like going out to eat")
  err = repo.AddCategory(
    ctx,
    repository.AddCategoryParams{
      Name: "Wants",  
      Description: categoryDescription,
      HouseholdID: household.ID,
    })
  if err != nil {
    w.WriteHeader(http.StatusBadRequest)
    w.Write([]byte(fmt.Sprintln("Encountered error when adding default categories - err: %s", err)))
    return
  }

  //3. Unexpected 
  categoryDescription.Scan("Use this category for necessary but unexpected expenses, like a copay after getting sick or irregular car maintenance")
  err = repo.AddCategory(
    ctx,
    repository.AddCategoryParams{
      Name: "Unexpected",  
      Description: categoryDescription,
      HouseholdID: household.ID,
    })
  if err != nil {
    w.WriteHeader(http.StatusBadRequest)
    w.Write([]byte(fmt.Sprintln("Encountered error when adding default categories - err: %s", err)))
    return
  }

  //4. Cultural 
  categoryDescription.Scan("Use this category for expenses concerning personal cultural learning, such as education or admission to a museum")
  err = repo.AddCategory(
    ctx,
    repository.AddCategoryParams{
      Name: "Cultural",  
      Description: categoryDescription,
      HouseholdID: household.ID,
    })
  if err != nil {
    w.WriteHeader(http.StatusBadRequest)
    w.Write([]byte(fmt.Sprintln("Encountered error when adding default categories - err: %s", err)))
    return
  }

  w.WriteHeader(http.StatusOK)
  w.Write([]byte(fmt.Sprintln("Successfully added the %s household and added default categories", householdName)))
}
