package main

import (
  "log"
  "net/http"
  "encoding/json"
  // "github.com/gorilla/mux"
  "fmt"
  // "strconv"
)

type Meal struct {
  ID        int    `json:"id"`
  Name      string `json:"name"`
  Foods     []Food `json:"foods"`
}

func getMeal(w http.ResponseWriter, r *http.Request) {

}
func getMeals(w http.ResponseWriter, r *http.Request) {
  w.Header().Set("Content-Type", "application/json")
  fmt.Println("GET MEALZZ")
  rows, err := database.Query("SELECT * FROM meals")


  if err != nil{
    log.Fatal(err)
  }

  var meals []Meal
  var meal Meal
  var food Food

  for rows.Next() {
    rows.Scan(&meal.ID, &meal.Name)
    meals = append(meals, meal)
  }

  fmt.Println(meals)

  for i, m := range meals {
    fmt.Println(m.ID)
    database.QueryRow("SELECT foods.id, foods.name, foods.calories FROM foods INNER JOIN meal_foods ON foods.id = meal_foods.food_id WHERE meal_foods.meal_id = $1 ", m.ID).Scan(&food.ID, &food.Name, &food.Calories)

    fmt.Println("Should append",food, "to", m)

    meals[i].Foods = append(meals[i].Foods, food)
    fmt.Println(meals[i])
  }



  json.NewEncoder(w).Encode(meals)
}
