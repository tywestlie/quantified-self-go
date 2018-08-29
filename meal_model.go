package main

import (
  "log"
  "net/http"
  "encoding/json"
  "github.com/gorilla/mux"
  "fmt"
  "strconv"
)

type Meal struct {
  ID        int    `json:"id"`
  Name      string `json:"name"`
  Foods     []Food `json:"foods"`
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

func getMeal(w http.ResponseWriter, r *http.Request) {
  var meal Meal
  var food Food

  w.Header().Set("Content-Type", "application/json")
  params := mux.Vars(r)
  id, _ := strconv.Atoi(params["id"])

  database.QueryRow("SELECT * FROM meals WHERE id= $1", id).Scan(&meal.ID, &meal.Name)

  rows, _ := database.Query("SELECT foods.id, foods.name, foods.calories FROM foods INNER JOIN meal_foods ON foods.id = meal_foods.food_id WHERE meal_foods.meal_id = $1 ", id)

  for rows.Next() {
    rows.Scan(&food.ID, &food.Name, &food.Calories)
    meal.Foods = append(meal.Foods, food)
  }

  json.NewEncoder(w).Encode(meal)
}

func createMealFood(w http.ResponseWriter, r *http.Request) {
  w.Header().Set("Content-Type", "application/json")
  params := mux.Vars(r)
  meal_id, _ := strconv.Atoi(params["meal_id"])
  food_id, _ := strconv.Atoi(params["food_id"])
  id := 0
  database.QueryRow("INSERT INTO meal_foods (meal_id, food_id) VALUES ($1, $2) RETURNING id").Scan(&id)

  json.NewEncoder(w).Encode(id)
}
