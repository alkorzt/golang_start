package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

// Person - контакт
type Person struct {
	gorm.Model
	Name  string
	Phone string
}

func allPersons(w http.ResponseWriter, r *http.Request) {
	db, err := gorm.Open("sqlite3", "test.db")
	if err != nil {
		panic("failed to connect database")
	}
	defer db.Close()

	var persons []Person
	db.Find(&persons)
	fmt.Println("{}", persons)

	json.NewEncoder(w).Encode(persons)
}

func newPerson(w http.ResponseWriter, r *http.Request) {
	fmt.Println("New Person Endpoint Hit")

	db, err := gorm.Open("sqlite3", "test.db")
	if err != nil {
		panic("failed to connect database")
	}
	defer db.Close()

	vars := mux.Vars(r)
	name := vars["name"]
	phone := vars["phone"]

	db.Create(&Person{Name: name, Phone: phone})
	fmt.Fprintf(w, "New Person Successfully Created")
}

func deletePerson(w http.ResponseWriter, r *http.Request) {
	db, err := gorm.Open("sqlite3", "test.db")
	if err != nil {
		panic("failed to connect database")
	}
	defer db.Close()

	vars := mux.Vars(r)
	name := vars["name"]

	var person Person
	db.Where("name = ?", name).Find(&person)
	db.Delete(&person)

	fmt.Fprintf(w, "Successfully Deleted Person")
}

func updatePerson(w http.ResponseWriter, r *http.Request) {
	db, err := gorm.Open("sqlite3", "test.db")
	if err != nil {
		panic("failed to connect database")
	}
	defer db.Close()

	vars := mux.Vars(r)
	name := vars["name"]
	phone := vars["phone"]

	var person Person
	db.Where("name = ?", name).Find(&person)

	person.Phone = phone

	db.Save(&phone)
	fmt.Fprintf(w, "Successfully Updated Person")
}

func initialMigration() {
	db, err := gorm.Open("sqlite3", "test.db")
	if err != nil {
		fmt.Println(err.Error())
		panic("failed to connect database")
	}
	defer db.Close()
	db.AutoMigrate(&Person{})
}

func handleRequests() {
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/persons", allPersons).Methods("GET")
	myRouter.HandleFunc("/person/{name}", deletePerson).Methods("DELETE")
	myRouter.HandleFunc("/person/{name}/{phone}", updatePerson).Methods("PUT")
	myRouter.HandleFunc("/person/{name}/{phone}", newPerson).Methods("POST")
	log.Fatal(http.ListenAndServe(":5555", myRouter))
}

func main() {
	fmt.Println("Go ORM Tutorial")
	initialMigration()
	handleRequests()
}
