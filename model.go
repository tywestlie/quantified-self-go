package main

import (
  "net/http"
  "encoding/json"
  "github.com/gorilla/mux"
  "fmt"
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
 fmt.Println(foods)
 json.NewEncoder(w).Encode(foods)
}
