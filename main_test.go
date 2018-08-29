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

  database.Exec("DELETE FROM meal_foods")
  database.Exec("ALTER SEQUENCE meal_foods_id_seq RESTART WITH 1")
  database.Exec("DELETE FROM meals")
  database.Exec("ALTER SEQUENCE meals_id_seq RESTART WITH 1")
  database.Exec("DELETE FROM foods")
  database.Exec("ALTER SEQUENCE foods_id_seq RESTART WITH 1")
  database.Exec(`INSERT INTO meals (id, name)
  VALUES (1, 'Breakfast'),
         (2, 'Snack'),
         (3, 'Lunch'),
         (4, 'Dinner')`)

  r := mux.NewRouter()

  routes(r)
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

  database.Exec("INSERT INTO foods (name, calories) VALUES ('Gourd', 500)")
  database.Exec("INSERT INTO foods (name, calories) VALUES ('Beans', 1200)")
  req, _ := http.NewRequest("GET", "/api/v1/foods/", nil)
  response := httptest.NewRecorder()
  r.ServeHTTP(response, req)
  actual := response.Body.String()
  actual = strings.TrimRight(actual, "\r\n ")
  expected := `[{"id":1,"name":"Gourd","calories":500},{"id":2,"name":"Beans","calories":1200}]`
  if actual != expected {
    t.Error("Get Foods - Expected:", expected, "Got:", actual)
  }
}

func TestGetFood(t * testing.T) {
  r := setup()
  database.Exec("INSERT INTO foods (name, calories) VALUES ('Gourd', 500)")
  database.Exec("INSERT INTO foods (name, calories) VALUES ('Beans', 1200)")
  req, _ := http.NewRequest("GET", "/api/v1/foods/1", nil)
  response := httptest.NewRecorder()
  fmt.Println(req)
  r.ServeHTTP(response, req)
  fmt.Println(response.Body)
  actual := response.Body.String()
  actual = strings.TrimRight(actual, "\r\n ")
  expected := `{"id":1,"name":"Gourd","calories":500}`
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

    if m["calories"] != 1200.00 {
        t.Errorf("Expected calories to be '1200'. Got '%v'", m["calories"])
    }
}

func TestDeleteFoods(t *testing.T) {
  r := setup()

  database.Exec("INSERT INTO foods (name, calories) VALUES ('Gourd', 500)")
  database.Exec("INSERT INTO foods (name, calories) VALUES ('Beans', 1200)")

  req, _ := http.NewRequest("DELETE", "/api/v1/foods/1", nil)
  response := httptest.NewRecorder()
  fmt.Println(req)
  r.ServeHTTP(response, req)
  actual := response.Code
  expected := 200

  if actual != expected {
    t.Error("Expected:", expected, "Got:", actual)
  }
}

func TestGetMeals(t *testing.T) {
    r := setup()

    database.Exec("INSERT INTO foods (name, calories) VALUES ('Gourd', 500)")
    database.Exec("INSERT INTO foods (name, calories) VALUES ('Beans', 1200)")
    database.Exec("INSERT INTO meal_foods (meal_id, food_id) VALUES ('1','2')")
    database.Exec("INSERT INTO meal_foods (meal_id, food_id) VALUES ('2','1')")

    req, _ := http.NewRequest("GET", "/api/v1/meals/", nil)
    response := httptest.NewRecorder()
    r.ServeHTTP(response, req)
    actual := response.Body.String()
    actual = strings.TrimRight(actual, "\r\n ")
    expected := `[{"id":1,"name":"Breakfast","foods":[{"id":2,"name":"Beans","calories":1200}]},{"id":2,"name":"Snack","foods":[{"id":1,"name":"Gourd","calories":500}]},{"id":3,"name":"Lunch","foods":null},{"id":4,"name":"Dinner","foods":null}]`

    if actual != expected {
      t.Error("Get Foods - Expected:", expected, "Got:", actual)
    }
}

func TestGetMeal(t *testing.T) {
  r := setup()

  database.Exec("INSERT INTO foods (name, calories) VALUES ('Gourd', 500)")
  database.Exec("INSERT INTO foods (name, calories) VALUES ('Beans', 1200)")
  database.Exec("INSERT INTO meal_foods (meal_id, food_id) VALUES ('1','2')")
  database.Exec("INSERT INTO meal_foods (meal_id, food_id) VALUES ('2','1')")

  req, _ := http.NewRequest("GET", "/api/v1/meals/1/foods/", nil)
  response := httptest.NewRecorder()
  r.ServeHTTP(response, req)
  actual := response.Body.String()
  actual = strings.TrimRight(actual, "\r\n ")
  expected := `{"id":1,"name":"Breakfast","foods":[{"id":2,"name":"Beans","calories":1200}]}`

  if actual != expected {
    t.Error("Get Foods - Expected:", expected, "Got:", actual)
  }
}

func TestPostMeal(t *testing.T) {
  r := setup()

  database.Exec("INSERT INTO foods (name, calories) VALUES ('Gourd', 500)")
  database.Exec("INSERT INTO foods (name, calories) VALUES ('Beans', 1200)")

  req, _ := http.NewRequest("POST", "/api/v1/meals/1/foods/1", nil)
  response := httptest.NewRecorder()
  r.ServeHTTP(response, req)
  actual := response.Body.String()
  actual = strings.TrimRight(actual, "\r\n ")
  expected := `1`
  if actual != expected {
    t.Error("Get Foods - Expected:", expected, "Got:", actual)
  }
}

func TestDeleteMeal(t *testing.T) {
  r := setup()

  database.Exec("INSERT INTO foods (name, calories) VALUES ('Gourd', 500)")
  database.Exec("INSERT INTO foods (name, calories) VALUES ('Beans', 1200)")
  database.Exec("INSERT INTO meal_foods (meal_id, food_id) VALUES ('1','2')")
  database.Exec("INSERT INTO meal_foods (meal_id, food_id) VALUES ('2','1')")

  req, _ := http.NewRequest("DELETE", "/api/v1/meals/1/foods/2", nil)
  response := httptest.NewRecorder()
  r.ServeHTTP(response, req)
  actual := response.Body.String()
  actual = strings.TrimRight(actual, "\r\n ")
  expected := `1`

  if actual != expected {
    t.Error("Get Foods - Expected:", expected, "Got:", actual)
  }
}
