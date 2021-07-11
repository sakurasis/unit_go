package main

import "fmt"

func main() {
	var num1 = [...]int{1, 3, 5, 7, 8}
	getNumEighth(num1)
}

func getNumEighth(arr [5]int) {
	for i := 0; i < len(arr); i++ {
		for j := len(arr) - 1; j >= i; j-- {
			sum := arr[i] + arr[j]
			if sum == 8 && i != j {
				fmt.Printf("[%d, %d]\t", i, j)
			}
		}
	}
}
