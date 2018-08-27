package main

import (
  "net/http"
  "encoding/json"
  "github.com/gorilla/mux"
  "fmt"
  "strconv"
)

type Food struct {
  ID  string `json :"id"`
  Name  string `json :"name"`
  Calories  int `json :"calories"`
}

func getFood(w http.ResponseWriter, r *http.Request) {
  var food Food
  params := mux.Vars(r)
    database.QueryRow("SELECT name, calories FROM foods WHERE id=$1",
       params["id"]).Scan(&food.Name, &food.Calories)
  json.NewEncoder(w).Encode(food)
}

func getFoods(w http.ResponseWriter, r *http.Request) {
 rows, _ := database.Query("SELECT * FROM foods;")

 fmt.Println(rows)

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
  var tupperWare TupperWare
  _ = json.NewDecoder(r.Body).Decode(&tupperWare)
  calories,_ := strconv.Atoi(tupperWare.NewFood.Calories)
  food := Food{Name: tupperWare.NewFood.Name, Calories: calories}
  query := "INSERT INTO foods (name, calories) VALUES ($1, $2) RETURNING id"
  fmt.Println(food)
  id := 0
  database.QueryRow(query, food.Name, food.Calories).Scan(&id)

  json.NewEncoder(w).Encode(id)
}
