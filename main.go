package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

type contact struct {
	ID        int    `json:"id"`
	LastName  string `json:"lastName"`
	FirstName string `json:"firtName"`
	Email     string `json:"email"`
	State     string `json:"state"`
	Phone     string `json:"phone"`
}

var contacts []contact

func logHandler(w http.ResponseWriter, r *http.Request) {
	log.Println(r.Method, r.URL)
}

func getContacts(w http.ResponseWriter, r *http.Request) {
	logHandler(w, r)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(&contacts)
}

func getContactsByState(w http.ResponseWriter, r *http.Request) {
	logHandler(w, r)
	w.Header().Set("Content-Type", "application/json")
	state := mux.Vars(r)["state"]
	var list []contact
	if state != "" {
		for _, contact := range contacts {
			if contact.State == state {
				list = append(list, contact)
			}
		}
	}
	json.NewEncoder(w).Encode(&list)
}

func loadData() {
	for i := 1; i <= 100; i++ {

		state := "FL"
		if i%2 == 0 {
			state = "CA"
		}
		contact := contact{ID: i, LastName: fmt.Sprintf("Doe%02d", i), FirstName: "John", Email: fmt.Sprintf("jdoe%02d@company.com", i), Phone: "(305)999-99999", State: state}
		contacts = append(contacts, contact)
	}
}

func main() {
	loadData()

	router := mux.NewRouter()
	router.HandleFunc("/api/contacts/{state}", getContactsByState).Methods("GET")
	router.HandleFunc("/api/contacts", getContacts).Methods("GET")

	port := os.Getenv("API_PORT")
	if port == "" {
		port = "80"
	}
	fmt.Println("Running on port", port)
	http.ListenAndServe(":"+port, router)
}
