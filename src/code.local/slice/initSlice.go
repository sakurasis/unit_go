package main

import (
	"fmt"
)

func main() {
	s := []int{1, 3, 5, 6, 7, 8}
	fmt.Printf("before:%v,add:%p\n", s, s)
	s1 := s[:]
	fmt.Printf("after:%v,add:%p\n", s1, s1)
	s2 := make([]int, 10, 10)
	copy(s2, s)
	fmt.Printf("copy:%v,add:%p\n", s2, s2)
}
