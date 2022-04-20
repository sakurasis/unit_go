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
		sum := 0
		for i := 0; i < len(values); i++ {
			sum += values[i].(int)
		}
		return sum
	}
	add5 := addCurry.Curry(5)

	v := add5(8, 3, 4)

	if v != 13 {
		fmt.Println("期望13，实际", v)
	} else {
		fmt.Println("real value: ", v)
	}

}
