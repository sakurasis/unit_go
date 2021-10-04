package main

import (
	"fmt"
	"strconv"
)

// Movie -- it describes some properties of the movie
type Movie struct {
	Name   string
	Rating float32
}

func (mv *Movie) summary() string {
	rate := strconv.FormatFloat(float64(mv.Rating), 'f', 3, 32)
	return mv.Name + ", " + rate
}

func main() {
	mv := Movie{
		Name:   "big bucket",
		Rating: 2.454,
	}
	fmt.Println(mv.summary())
}
