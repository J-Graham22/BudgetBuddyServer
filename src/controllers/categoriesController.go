package controllers

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/J-Graham22/BudgetBuddyServer/src/db"
	"github.com/J-Graham22/BudgetBuddyServer/src/db/repository"
)

// CREATE
func AddCategory(w http.ResponseWriter, r *http.Request) {
  householdId := r.PathValue("id")

  if householdId == "" {
    w.WriteHeader(http.StatusBadRequest)
    w.Write([]byte("Name of new household not provided in request"))
    return
  }
  householdIdInt, err := strconv.ParseInt(householdId, 10, 64)
  if err != nil {
    w.WriteHeader(http.StatusBadRequest)
    w.Write([]byte("could not parse household_id into int"))
    return
  }

  decoder := json.NewDecoder(r.Body)
  var category repository.Category
  err = decoder.Decode(&category)
  if err != nil {
    w.WriteHeader(http.StatusBadRequest)
    w.Write([]byte("could not deserialize json body into category"))
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
  err = repo.AddCategory(ctx, repository.AddCategoryParams{
    Name: category.Name,
    HouseholdID: int32(householdIdInt),
    Description: category.Description,
    ParentCategoryID: category.ParentCategoryID,
  })
  if err != nil {
    w.WriteHeader(http.StatusInternalServerError)
    w.Write([]byte(fmt.Sprintf("Encountered error while trying to add category - err: %s", err)))
    return
  }

  w.WriteHeader(http.StatusOK)
  w.Write([]byte(fmt.Sprintf("Successfully added category %s for household %d", category.Name, householdIdInt)))
}


// READ 
func GetCategoriesForHousehold(w http.ResponseWriter, r *http.Request) {

}

// UPDATE 
func UpdateCategoryName(w http.ResponseWriter, r *http.Request) {

}

// DELETE
func DeleteCategory(w http.ResponseWriter, r *http.Request) {
  //TODO

  //determine behavior for when a category is deleted
  // -- for all transactions that match that category 
  // ---- if the category has a parent category, make that the new category
  // ---- if the category has no parent, leave the transactions uncategorized
}
