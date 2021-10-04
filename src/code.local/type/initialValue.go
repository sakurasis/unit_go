package main

import (
	"fmt"
	"math"
	"unsafe"
)

type in int

func main() {
	age, name := 18, "\"letgo\""
	fmt.Println("my age is", age, ", name is", name)
	firstNum, secondNum := 12.3, 12.31
	tempNum := math.Min(firstNum, secondNum)

	fmt.Println("min Numb is", tempNum)
	fmt.Println("*********************************")

	var a int = 18
	fmt.Printf("type of a is %T, size of a is %d\n", a, unsafe.Sizeof(a))
	var c, b in = 31, 1341223
	d := c + b

	fmt.Printf("d is %d\n", d)
}
