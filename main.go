package main

import (
	"encoding/json"
	"log"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// Record struct (Model)
type Record struct {
	ID     string  `json:"id"`
	Isbn   string  `json:"isbn"`
	Title  string  `json:"title"`
	patient *patient `json:"patient"`
}

// patient struct
type patient struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
}

// Init Records var as a slice Record struct
var Records []Record

// Get all Records
func getRecords(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(Records)
}

// Get single Record
func getRecord(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r) // Gets params
	// Loop through Records and find one with the id from the params
	for _, item := range Records {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&Record{})
}

// Add new Record
func createRecord(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var Record Record
	_ = json.NewDecoder(r.Body).Decode(&Record)
	Record.ID = strconv.Itoa(rand.Intn(100000000)) // Mock ID - not safe
	Records = append(Records, Record)
	json.NewEncoder(w).Encode(Record)
}

// Update Record
func updateRecord(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range Records {
		if item.ID == params["id"] {
			Records = append(Records[:index], Records[index+1:]...)
			var Record Record
			_ = json.NewDecoder(r.Body).Decode(&Record)
			Record.ID = params["id"]
			Records = append(Records, Record)
			json.NewEncoder(w).Encode(Record)
			return
		}
	}
}

// Delete Record
func deleteRecord(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range Records {
		if item.ID == params["id"] {
			Records = append(Records[:index], Records[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(Records)
}

// Main function
func main() {
	// Init router
	r := mux.NewRouter()

	// Hardcoded data 
	Records = append(Records, Record{ID: "1", Isbn: "438227", Title: "Record One", patient: &patient{Firstname: "smsm", Lastname: "hhh"}})
	Records = append(Records, Record{ID: "2", Isbn: "454555", Title: "Record Two", patient: &patient{Firstname: "3m", Lastname: "grges"}})

	// Route handles & endpoints
	r.HandleFunc("/Records", getRecords).Methods("GET")
	r.HandleFunc("/Records/{id}", getRecord).Methods("GET")
	r.HandleFunc("/Records", createRecord).Methods("POST")
	r.HandleFunc("/Records/{id}", updateRecord).Methods("PUT")
	r.HandleFunc("/Records/{id}", deleteRecord).Methods("DELETE")

	// Start server
	log.Fatal(http.ListenAndServe(":8000", r))
}

// Request sample
// {
// 	"isbn":"4545454",
// 	"title":"Record Three",
// 	"patient":{"firstname":"Dr4","lastname":"lool"}
// }
