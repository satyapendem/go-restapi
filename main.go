package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type User struct {
	Name string `json: "name"`
	ID   string `json: "id"`
}

var user []User

func getAllUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}
func main() {

	user = append(user, User{ID: "1223", Name: "satya"})
	router := mux.NewRouter()

	router.HandleFunc("/api/v1/users", getAllUsers).Methods("GET")

	log.Fatal(http.ListenAndServe(":5678", router))

}
