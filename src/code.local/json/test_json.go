package main

import (
	"fmt"
	"encoding/json"
)

type Movie struct {
	Name string `json:"name"`
	Year int `json:"year"`
	Score float64 `json:"score"`
	Actors []string `json:"actors"`
}

func main() {
	m := Movie{"avtar2", 2022, 9.88, []string{"lily", "tomperson"}}
	str, err := json.Marshal(m)
	if err != nil {
		panic(err)
	}
	
	fmt.Printf("json:%s\n",str)

	var mv Movie 

	jsonstr := "{\"name\":\"avtar\",\"year\":2002,\"score\":9.00,\"actors\":[\"lily\",\"tomperson\"]}"

	err = json.Unmarshal([]byte(jsonstr), &mv)
	if err != nil {
		panic(err)
	}
	fmt.Println(mv)
}