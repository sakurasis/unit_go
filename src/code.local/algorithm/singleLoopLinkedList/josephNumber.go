package main

import "fmt"

func main() {
	fmt.Println("the last:", joseph(11, 3))
}

func joseph(n, m int) int {
	if n == 1 {
		return 0
	}
	return (joseph(n-1, m) + m) % n
}
