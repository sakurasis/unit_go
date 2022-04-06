package main

import (
	"fmt"
	"main/curry"
)

func add(x, y int) int {
	return x + y
}

func main() {

	var addCurry curry.Function = func(values ...interface{}) interface{} {
		return add(values[0].(int), values[1].(int))
	}
	add5 := addCurry.Curry(5)

	v := add5(8)

	if v != 13 {
		fmt.Println("期望13，实际", v)
	} else {
		fmt.Println("real value: ", v)
	}

}
