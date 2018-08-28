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
}

func getMeal(w http.ResponseWriter, r *http.Request) {

}
func getMeals(w http.ResponseWriter, r *http.Request) {

}
