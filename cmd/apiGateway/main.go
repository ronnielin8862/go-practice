package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/ronnielin8862/go-practice/pkg/userService"

	"github.com/gorilla/mux"
)

var people []userService.Person
var Address []userService.Address

func GetPerson(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	for _, item := range people {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(people)
}

func GetPeople(w http.ResponseWriter, req *http.Request) {
	json.NewEncoder(w).Encode(people)
}

func PostPerson(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	var person userService.Person
	_ = json.NewDecoder(req.Body).Decode(&person)
	person.ID = params["id"]
	people = append(people, person)
	json.NewEncoder(w).Encode(people)
}

func DeletePerson(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	for index, item := range people {
		if item.ID == params["id"] {
			people = append(people[:index], people[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(people)
}

func main() {
	router := mux.NewRouter()
	people = append(people, userService.Person{ID: "1", Firstname: "xi", Lastname: "dada", Address: &userService.Address{City: "Shenyang", Province: "Liaoning"}})
	people = append(people, userService.Person{ID: "2", Firstname: "li", Lastname: "xiansheng", Address: &userService.Address{City: "Changchun", Province: "Jinlin"}})
	router.HandleFunc("/people", GetPeople).Methods("GET")
	router.HandleFunc("/people/{id}", GetPerson).Methods("GET")
	router.HandleFunc("/people/{id}", PostPerson).Methods("POST")
	router.HandleFunc("/people/{id}", DeletePerson).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":12345", router))
}
