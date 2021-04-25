package main

import "fmt"

func main() {
	fmt.Println(test())
	fmt.Println(foo(100))
}
func test() (ret int) {
	defer func() {
		ret++
	}()
	return 1
}

func foo(n int) int {
	defer func() {n++}()             
	return n
}