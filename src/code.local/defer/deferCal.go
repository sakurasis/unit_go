package main

import "fmt"

func calc(index string, a, b int) int {
	ret := a + b
	_, err := fmt.Println(index, a, b, ret)
	if err != nil {
		return 0
	}
	return ret
}

func main() {
	x := 1
	y := 2
	defer calc("AA", x, calc("A", x, y))
	x = 10
	defer calc("BB", x, calc("B", x, y))
	y = 20
}
