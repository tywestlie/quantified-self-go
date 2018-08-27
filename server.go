package main

import (
  "log"
  "net/http"
  "fmt"
  "os"
  "database/sql"
  "github.com/gorilla/mux"
  "github.com/lib/pq"
)

// type App struct {
//   DB  *sql.DB
// }

var database *sql.DB

func initializeDB() {
  database, err = sql.Open("postgres", dbname())
  if err != nil {
    log.Fatal(err)
  }
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

func main() {
  initializeDB()

  port := getPort()
  router := mux.NewRouter()

  router.HandleFunc("/", root).Methods("GET")
  router.HandleFunc("/api/v1/foods/{id}", getFood).Methods("GET")


  log.Fatal(http.ListenAndServe(port, router))
}
