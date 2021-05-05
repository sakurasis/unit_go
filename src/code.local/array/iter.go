package main

import "fmt"

func main() {
	arr := [...]int{99: -1}
	for _, v := range arr {
		fmt.Printf("%d ",v)
	}
	fmt.Printf("the array's length is %d\n", len(arr)) 
}