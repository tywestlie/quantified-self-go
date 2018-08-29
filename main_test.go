package main

import (
  // "fmt"
  "testing"
  "net/http"
  "net/http/httptest"
  "github.com/gorilla/mux"
  "strings"
  // "bytes"
)

func setup() *mux.Router {
  database.Exec("TRUNCATE TABLE meal_foods RESTART IDENTITY")
  database.Exec("TRUNCATE TABLE foods RESTART IDENTITY")
  database.Exec("TRUNCATE TABLE meals RESTART IDENTITY")
  seedDB()

  r := mux.NewRouter()
  routesSetup(r)
  return r
}

func TestGetPort(t *testing.T) {
  port := getPort()
  if port != ":3000" {
    t.Error("Expected :3000, got ", port)
  }
}

func TestGetFoods(t *testing.T) {
  r := setup()
  database.Exec("INSERT INTO foods (name, calories) VALUES ('Hotdog', 500)")
  database.Exec("INSERT INTO foods (name, calories) VALUES ('Burger', 1200)")
  req, _ := http.NewRequest("GET", "/api/v1/foods/", nil)
  response := httptest.NewRecorder()
  r.ServeHTTP(response, req)
  actual := response.Body.String()
  actual = strings.TrimRight(actual, "\r\n ")
  expected := `[{"id":1,"name":"Hotdog","calories":500},{"id":2,"name":"Burger","calories":1200}]`
  if actual != expected {
    t.Error("Get Foods - Expected:", expected, "Got:", actual)
  }
}
