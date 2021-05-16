package main

import "fmt"

func main() {
	s := []int {5,6,7,8,9}
	fmt.Printf("slice:%t\n", s[2:])
	fmt.Printf("slice:%t\n", s[3:])
	copy(s[2:], s[3:])
	fmt.Printf("slice:%t\n", s)
	fmt.Printf("s:%t\n", s[:len(s)-1])

	s[2] = s[len(s)-1]
}