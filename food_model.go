package main

import (
  "log"
  "net/http"
  "encoding/json"
  "github.com/gorilla/mux"
  "fmt"
  "strconv"
)

type Food struct {
  ID        int    `json:"id"`
  Name      string `json:"name"`
  Calories  int    `json:"calories"`
}

func getFood(w http.ResponseWriter, r *http.Request) {
  w.Header().Set("Content-Type", "application/json")
  var food Food
  params := mux.Vars(r)

  database.QueryRow("SELECT name, calories FROM foods WHERE id=$1",
    params["id"]).Scan(&food.Name, &food.Calories)
    json.NewEncoder(w).Encode(food)
  }

  func getFoods(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    rows, err := database.Query("SELECT * FROM foods")
    fmt.Println(rows)

    if err != nil{
      log.Fatal(err)
    }

    foods := []Food {}

    for rows.Next() {
      var food Food
      rows.Scan(&food.ID, &food.Name, &food.Calories)
      foods = append(foods, food)
    }

    json.NewEncoder(w).Encode(foods)
  }

  type TupperWare struct {
    NewFood NewFood `json:"food"`
  }

  type NewFood struct {
    Name string `json:"name"`
    Calories string `json:"calories"`
  }

  func createFood(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    var tupperWare TupperWare
    _ = json.NewDecoder(r.Body).Decode(&tupperWare)

    fmt.Printf("Params: %#v\n", tupperWare)

    calories,_ := strconv.Atoi(tupperWare.NewFood.Calories)
    food := Food{Name: tupperWare.NewFood.Name, Calories: calories}
    query := "INSERT INTO foods (name, calories) VALUES ($1, $2) RETURNING id"
    fmt.Println(food.Name, food.Calories)
    id := 0
    err := database.QueryRow(query, food.Name, food.Calories).Scan(&id)
    if err != nil {
      log.Fatal(err)
    }
    fmt.Println("Added food", id)

    json.NewEncoder(w).Encode(food)
  }

  // func updateFood(w http.ResponseWriter, r *http.Request) {
  //   vars := mux.Vars(r)
  //   id, err := strcov.Atoi(vars["id"])
  //   if err != nil {
  //      respondWithError(w, http.StatusBadRequest, "Invalid food ID")
  //      return
  //  }
  //
  //  var food Food
  //  decoder := json.NewDecoder(r.Body)
  //
  //  food.ID = id
  // }

  func deleteFood(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    var food Food

    params := mux.Vars(r)
    id, _ := strconv.Atoi(params["id"])
    database.QueryRow("DELETE FROM foods WHERE id=$1 RETURNING id,  name", id).Scan(&id, &food.Name)

    json.NewEncoder(w).Encode(id)
    }
