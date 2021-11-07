package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Person struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
}
type PersonShort struct {
	FirstName string `json:"First"`
}

func main() {
	http.HandleFunc("/encode", encodeFunc)
	http.HandleFunc("/decode", decodeFunc)

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}

func encodeFunc(w http.ResponseWriter, r *http.Request) {
	p1 := Person{FirstName: "Alex", LastName: "J"}

	err := json.NewEncoder(w).Encode(p1)
	if err != nil {
		log.Println("Encoded bad data", err)
	}
}

func decodeFunc(w http.ResponseWriter, r *http.Request) {
	shortPeople := []PersonShort{}
	err := json.NewDecoder(r.Body).Decode(&shortPeople)
	if err != nil {
		log.Println("Bad data sent", err)
	}
	fmt.Println(shortPeople)
}
