package main

import "fmt"

func recure(num int) int {
	if num == 0 {
		return 1
	}
	return num * recure(num - 1) 
}

func main() {
	fmt.Println(recure(3)) // 3! = 3*2*1
}