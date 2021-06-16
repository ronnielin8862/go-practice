package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)


func main() {
	router := mux.NewRouter()
	people = append(people, Person{ID: "1", Firstname: "xi", Lastname: "dada", Address: &Address{City: "Shenyang", Province: "Liaoning"}})
	people = append(people, Person{ID: "2", Firstname: "li", Lastname: "xiansheng", Address: &Address{City: "Changchun", Province: "Jinlin"}})
	router.HandleFunc("/people", GetPeople).Methods("GET")
	router.HandleFunc("/people/{id}", GetPerson).Methods("GET")
	router.HandleFunc("/people/{id}", PostPerson).Methods("POST")
	router.HandleFunc("/people/{id}", DeletePerson).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":12345", router))
}
