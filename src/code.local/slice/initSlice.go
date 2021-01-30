package main

import "fmt"

func main() {
	a := make([]int, 2)
	a[0] = 5
	a[1] = 4
	a = append(a, 3, 2, 1)
	fmt.Println("a:", a)
}