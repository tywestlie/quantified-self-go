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
  // var food Food

  for rows.Next() {
    rows.Scan(&meal.ID, &meal.Name)
    meals = append(meals, meal)
  }

  fmt.Println(meals)

  for i, m := range meals {
  meals[i].Foods = getMealFoods(m.ID)
  }

  json.NewEncoder(w).Encode(meals)
}

func getMealFoods(meal_id int) []Food {
  rows, err := database.Query("SELECT foods.id, foods.name, foods.calories FROM foods INNER JOIN meal_foods ON foods.id=meal_foods.food_id WHERE meal_foods.meal_id=$1", meal_id)
  if err != nil {
    log.Fatal(err)
  }
  var (
    food Food
    foods []Food
  )
  defer rows.Close()

  for rows.Next() {
    if err := rows.Scan(&food.ID, &food.Name, &food.Calories); err != nil {
      log.Fatal(err)
    }
    foods = append(foods, food)
  }
  return foods
}


func getMeal(w http.ResponseWriter, r *http.Request) {
  var meal Meal

  w.Header().Set("Content-Type", "application/json")
  params := mux.Vars(r)
  id, _ := strconv.Atoi(params["id"])

  database.QueryRow("SELECT * FROM meals WHERE id= $1", id).Scan(&meal.ID, &meal.Name)

  meal.Foods = getMealFoods(id)

  json.NewEncoder(w).Encode(meal)
}

func createMealFood(w http.ResponseWriter, r *http.Request) {
  w.Header().Set("Content-Type", "application/json")
  params := mux.Vars(r)

  fmt.Println("params:", params)
  meal_id, _ := strconv.Atoi(params["meal_id"])
  fmt.Println("meal id", meal_id)
  food_id, _ := strconv.Atoi(params["food_id"])
  fmt.Println("food id", food_id)
  id := 0
  database.QueryRow("INSERT INTO meal_foods (meal_id, food_id) VALUES ($1, $2) RETURNING id", meal_id, food_id).Scan(&id)

  json.NewEncoder(w).Encode(id)
}
