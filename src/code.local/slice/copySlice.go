package main

import "fmt"

func main() {
	s := []int{5, 6, 7, 8, 9}
	fmt.Printf("slice:%v\n", s[2:])
	fmt.Printf("slice:%v\n", s[3:])
	copy(s[2:], s[3:])
	fmt.Printf("slice:%v\n", s)
	fmt.Printf("s:%v\n", s[:len(s)-1])

	s[2] = s[len(s)-1]
}
