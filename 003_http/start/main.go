package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
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

func homeIndex(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "it work...")
}

func handleRequest() {
	http.HandleFunc("/", homeIndex)
	http.HandleFunc("/contacts", allContacts)
	log.Fatal(http.ListenAndServe(":5555", nil))
}

func main() {
	handleRequest()
}
