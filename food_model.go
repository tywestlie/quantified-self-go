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

  database.QueryRow("SELECT * FROM foods WHERE id=$1",
    params["id"]).Scan(&food.ID, &food.Name, &food.Calories)
  json.NewEncoder(w).Encode(food)
  fmt.Println(food)
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

  func updateQuery(food Food) {
    food_id := 0
    database.QueryRow("UPDATE foods SET name=$1, calories=$2 WHERE id=$3 RETURNING id", food.Name, food.Calories, food.ID).Scan(&food_id)
  }


  func updateFood(w http.ResponseWriter, r *http.Request) {
    var food Food
    var tupperWare TupperWare
    _ = json.NewDecoder(r.Body).Decode(&tupperWare)
    calories,_ := strconv.Atoi(tupperWare.NewFood.Calories)
    food = Food{Name: tupperWare.NewFood.Name, Calories: calories}
    updateQuery(food)
  }

  func deleteFood(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    var food Food

    params := mux.Vars(r)
    id, _ := strconv.Atoi(params["id"])
    database.QueryRow("DELETE FROM foods WHERE id=$1 RETURNING id,  name", id).Scan(&id, &food.Name)

    json.NewEncoder(w).Encode(id)
  }
