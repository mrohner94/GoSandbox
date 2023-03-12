package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

// BaseURL is the base endpoint for the star wars API
const BaseURL = "https://swapi.dev/api/"

type Planet struct {
	Name string `json:"name"`
	Population string `json:"population"`
	Terrain string `json:"terrain"`
}

type Person struct {
	Name string `json:"name"`
	HomeworldURL string `json:"homeworld"`
	Homeworld Planet
}

type AllPeople struct {
	People []Person `json:"results"`
}

func (p *Person) getHomeworld() {
	res, err := http.Get(p.HomeworldURL)
	if err != nil {
		log.Print("Error fetching homeworld", err)
	}

	var bytes []byte
	if bytes, err = io.ReadAll(res.Body); err != nil {
		log.Print("Error reading response body", err)
	}

	json.Unmarshal(bytes, &p.Homeworld)
}

func getPeople(w http.ResponseWriter, r *http.Request) {
	res, err := http.Get(BaseURL + "people")

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		log.Print("Failed to request star wars people")
	}


	//fmt.Println(res)
	bytes, err := io.ReadAll(res.Body)

	if(err != nil) {
		http.Error(w, err.Error(), http.StatusBadRequest)
		log.Print("Failed to request star wars people")
	}

	var people AllPeople

	// fmt.Println(string(bytes))
	if err := json.Unmarshal(bytes, &people); err != nil {
		fmt.Println("Error parsing json", err)
	}

	fmt.Println(people)

	for _, pers := range people.People {
		pers.getHomeworld()
		fmt.Println(pers)
	}
}


func main() {
	port := ":8080"

	http.HandleFunc("/people", getPeople)
	fmt.Println("Serving on ", port)
	log.Fatal(http.ListenAndServe(port, nil))
}
