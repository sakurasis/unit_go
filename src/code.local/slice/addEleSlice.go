package main

import "fmt"

func main() {
	// add some elements to slice.
	var fruit = make([]string, 5)
	fruit = append(fruit, "apple")
	fruit = append(fruit, "banana", "pear", "cherry", "peach")
	
	for idx, v := range fruit {
		if idx == 0 {
			fmt.Printf("fruit:")
		}
		fmt.Printf("%s ", v)
	}
	fmt.Printf("\n")
}
