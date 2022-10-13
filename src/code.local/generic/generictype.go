package main

import (
	"fmt"
)

type Slice[T int | float32 | float64 | string] []T

func main() {
	var a Slice[int] = []int{1, 2, 3}
	fmt.Printf("Type Name: %v", a)
}
