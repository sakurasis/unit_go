package main

import "fmt"

func main() {
	p := new(int)
	fmt.Println(*p)

	*p = 2
	fmt.Println(p)

	var x = 1 << 1 | 1 << 5
	fmt.Printf("%08b\n", x)      
}                                                             