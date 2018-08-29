package main

import (
  "fmt"
  "testing"
  "bytes"
  "strings"
  "net/http"
  "net/http/httptest"
  "github.com/gorilla/mux"
  "encoding/json"
)

func setup() *mux.Router {
  initializeDB()

  r := mux.NewRouter()
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
  req, _ := http.NewRequest("GET", "/api/v1/foods", nil)
  response := httptest.NewRecorder()
  fmt.Println(req)
  r.ServeHTTP(response, req)
  fmt.Println(response.Body)
  actual := response.Body.String()
  actual = strings.TrimRight(actual, "\r\n ")
  expected := `[{"id":1,"name":"Hotdog","calories":500},{"id":2,"name":"Burger","calories":1200}]`
  if actual != expected {
    t.Error("Get Foods - Expected:", expected, "Got:", actual)
  }
}

func TestGetFood(t * testing.T) {
  r := setup()
  database.Exec("INSERT INTO foods (name, calories) VALUES ('Hotdog', 500)")
  database.Exec("INSERT INTO foods (name, calories) VALUES ('Burger', 1200)")
  req, _ := http.NewRequest("GET", "/api/v1/foods/1", nil)
  response := httptest.NewRecorder()
  fmt.Println(req)
  r.ServeHTTP(response, req)
  fmt.Println(response.Body)
  actual := response.Body.String()
  actual = strings.TrimRight(actual, "\r\n ")
  expected := `{"id":1,"name":"Hotdog","calories":500}`
  if actual != expected {
    t.Error("Get Foods - Expected:", expected, "Got:", actual)
  }
}

func TestCreateFood(t *testing.T) {
  r := setup()
  newfood := []byte(`{"food":{"name":"burrito","calories":"1200"}}`)
  req, _ := http.NewRequest("POST", "/api/v1/foods/", bytes.NewBuffer(newfood))
  response := httptest.NewRecorder()
  r.ServeHTTP(response, req)

  var m map[string]interface{}
    json.Unmarshal(response.Body.Bytes(), &m)

    if m["name"] != "burrito" {
        t.Errorf("Expected food name to be 'burrito'. Got '%v'", m["name"])
    }

    if m["calories"] != 1200 {
        t.Errorf("Expected product price to be '1200'. Got '%v'", m["calories"])
    }
}

func TestDeleteFoods(t *testing.T) {
  r := setup()
  database.Exec("INSERT INTO foods (name, calories) VALUES ('Hotdog', 500)")
  database.Exec("INSERT INTO foods (name, calories) VALUES ('Burger', 1200)")

  req, _ := http.NewRequest("DELETE", "/api/v1/foods/1", nil)
  response := httptest.NewRecorder()
  fmt.Println(req)
  r.ServeHTTP(response, req)
  actual := response.Code
  expected := 204
  
  if actual != expected {
    t.Error("Expected:", expected, "Got:", actual)
  }
}
