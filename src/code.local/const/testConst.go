package main

import (
	"fmt"
	"math"
)

func main() {
	const a = 55
	fmt.Println("a is", a)
	b := math.Sqrt(4)
	fmt.Println("b is", b)

	name := "summy"
	fmt.Printf("type %T value %v\n", name, name)

	const c = math.Pi
	fmt.Printf("c is %.20f\n", c)

	const initBool = true
	type cusBool bool
	var myBool cusBool = initBool
	fmt.Println("myBool value is", myBool)
}
