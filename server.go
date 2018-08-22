package main

import (
  // "encoding/json"
  "log"
  "net/http"
  "fmt"
  // "math/rand"
  // "strconv"
  "github.com/gorilla/mux"

)

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
  router := mux.NewRouter()

  router.HandleFunc("/", root).Methods("GET")

  log.Fatal(http.ListenAndServe(":3000", router))
}
