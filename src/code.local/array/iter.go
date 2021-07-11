package main

import "fmt"

func main() {
	arr := [...]int{99: -1}
	var sum int
	var num = [...]int{1, 3, 5, 7, 8}
	for _, v := range arr {
		fmt.Printf("%d ", v)
	}
	for _, v1 := range num {
		fmt.Printf("%d\t", v1)
		sum += v1
	}
	fmt.Println("the sum is ", sum)
	fmt.Printf("the array's length is %d\n", len(arr))
}
