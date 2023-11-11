package main

import (
	"fmt"
	"encoding/json"
	"log"
)

type Movie struct {
	Name string
	Year int `json:"blabla"` // дескриптор поля, определяет имя поля для json
	Actors []string `json:",omitempty"` // дескриптор поля, не выводит пустое поле в json
}

var movies = []Movie{{Name: "Titanic", Year: 1980, Actors: []string{"M", "D"}}, {Name: "fdf", Year: 1233, Actors: []string{}}}

var names []struct{
	Name string
}

func main() {
	fmt.Printf("%v\n", movies)
	data, err := json.Marshal(movies)
	if err != nil {
		log.Fatalf("%s\n", err)
	}
	fmt.Printf("%s\n", data)
	data, err = json.MarshalIndent(movies, "", "  ")
	if err != nil {
		log.Fatalf("%s\n", err)
	}
	fmt.Printf("%s\n", data)

	if err := json.Unmarshal(data, &names); err != nil {
		log.Fatalf("%s\n", err)
	}
	fmt.Println(names)
}
