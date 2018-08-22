package main

import (
  "log"
  "net/http"
  "fmt"
  "os"
  "github.com/gorilla/mux"
)

func getPort() string {
  p := os.Getenv("PORT")
  if p != "" {
    return ":" + p
  }
  return ":3000"
}

// Food Struct (Model)

type Food struct {
  ID  string `json :"id"`
  Name  string `json :"name"`
  calories  int `json :"calories"`
}

func root(w http.ResponseWriter, r *http.Request) {
  fmt.Fprintf(w, "Quantifed Self Go Backend")
}

func main() {
  port := getPort()
  router := mux.NewRouter()

  router.HandleFunc("/", root).Methods("GET")


  log.Fatal(http.ListenAndServe(port, router))
}
