package main

import (
  "net/http"
  "encoding/json"
  "github.com/gorilla/mux"
  // "errors"
)

type Food struct {
  ID  string `json :"id"`
  Name  string `json :"name"`
  Calories  int `json :"calories"`
}

func getFood(w http.ResponseWriter, r *http.Request) {
  var f Food
  params := mux.Vars(r)
    database.QueryRow("SELECT name, calories FROM foods WHERE id=$1",
       params["id"]).Scan(&f.Name, &f.Calories)
  json.NewEncoder(w).Encode(f)
}
