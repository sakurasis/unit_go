package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type Movie struct {
	Title string
	Year int `json:"released"`
	Color bool `json:"color,omitempty"`
	Actors []string
}

var movies = []Movie {
	{Title: "Cool Hand Luke", Year: 1967, Color: true,
		Actors: []string{"Paul Newman"}},
		{Title: "marshal", Year: 1850, Color: false,
			Actors: []string{"steve jackson","jacqueline Bisset"}},
}

func main() {
	marshal, err := json.MarshalIndent(movies, "", "    ")
	if err != nil {
		log.Fatalf("Json marshaling failed: %s", err)
	}
	fmt.Printf("%s\n", marshal)
}
