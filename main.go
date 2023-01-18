package main

import (
	"encoding/json"
	"fmt"
	"net/http"
 "github.com/gorilla/mux"
)

var customers = map[string]uint32{
	"Yoan":       1,
	"Daniel":     2,
	"Kelly":      3,
	"Antonio":    4,
}

func index(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(customers)
}

func getMembers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(customers)
}

func deleteMember(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	id := mux.Vars(r)["id"]

	if _, ok := customers[id]; ok {
		delete(customers, id)
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(customers)
	} else {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(customers)
	}
}

func main() {
	http.HandleFunc("/", index)
router := mux.NewRouter().StrictSlash(true)
router.HandleFunc("/members", getMembers).Methods("GET")
router.HandleFunc("/deleteMember/{id}", deleteMember).Methods("DELETE")
fmt.Println("Server is starting on port 3000...")
        http.ListenAndServe(":3000", nil)
}
