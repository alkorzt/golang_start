package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// Person - контакт
type Person struct {
	Name         string `json:"Name"`
	Surname      string `json:"Surname"`
	Patronomical string `json:"Patronomical"`
}

// PersonBook -  контактная книга
type PersonBook []Person

func allContacts(w http.ResponseWriter, r *http.Request) {
	contacts := PersonBook{
		Person{Name: "Ivan", Surname: "Smirnov", Patronomical: "Petrovich"},
	}
	json.NewEncoder(w).Encode(contacts)
}

func testPostContacts(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "test...")
}

func homeIndex(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "it work...")
}

func handleRequest() {
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/", homeIndex)
	myRouter.HandleFunc("/contacts", allContacts).Methods("GET")
	myRouter.HandleFunc("/contacts", testPostContacts).Methods("POST")
	log.Fatal(http.ListenAndServe(":5555", myRouter))
}

func main() {
	handleRequest()
}
