package main

import (
	"fmt"
)

func main() {
	test1()
}

func test1() {
	for i := 0; i < 10; i++ {
		g := func(x int) { fmt.Printf("%d\t", x) }
		g(i)
	}
	fmt.Printf("\n")
	fmt.Println("loop end!!!")
}
