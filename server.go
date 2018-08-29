package main

import (
  "log"
  "net/http"
  "fmt"
  "os"
  "database/sql"
  "github.com/gorilla/mux"
  "github.com/gorilla/handlers"
  "github.com/lib/pq"
)

var database *sql.DB

func initializeDB(){
  db, err := sql.Open("postgres", dbname())
  if err != nil {
    log.Fatal(err)
  }
  database = db
  migrateDB()
  seedDB()
}

func dbname() string {
url := os.Getenv("DATABASE_URL")
if url != "" {
  connection, _ := pq.ParseURL(url)
  connection += " sslmode=require"
  return connection
}
return "dbname=qs_go_dev sslmode=disable"
}

func getPort() string {
  p := os.Getenv("PORT")
  if p != "" {
    return ":" + p
  }
  return ":3000"
}

func root(w http.ResponseWriter, r *http.Request) {
  fmt.Fprintf(w, "Quantifed Self Go Backend")
}


func routesSetup(r *mux.Router) {
  r.HandleFunc("/", root).Methods("GET")
  r.HandleFunc("/api/v1/foods/", createFood).Methods("POST")
  r.HandleFunc("/api/v1/foods/", getFoods).Methods("GET")
  r.HandleFunc("/api/v1/foods/{id}", getFood).Methods("GET")
  r.HandleFunc("/api/v1/foods/{id}", deleteFood).Methods("DELETE")
  r.HandleFunc("/api/v1/meals/", getMeals).Methods("GET")
  r.HandleFunc("/api/v1/meals/{id}/foods/", getMeal).Methods("GET")
  r.HandleFunc("/api/v1/meals/{meal_id}/foods/{food_id}", createMealFood).Methods("POST")
  r.HandleFunc("/api/v1/meals/{meal_id}/foods/{food_id}", deleteMealFood).Methods("DELETE")
}

func main() {
  initializeDB()

  r := mux.NewRouter()
  routesSetup(r)
  port := getPort()

  log.Fatal(http.ListenAndServe(port, handlers.CORS(
  handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"}),
  handlers.AllowedMethods([]string{"GET", "POST", "PUT", "HEAD", "DELETE", "OPTIONS"}),
  handlers.AllowedOrigins([]string{"*"}))(r)))
}
