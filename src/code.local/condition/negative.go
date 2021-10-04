package main

import (
	"fmt"
)

func main() {
	val := 0
	for val <= 0 {
		fmt.Print("Enter the number which is a integer number:")
		fmt.Scanf("%d", &val)

		switch {
		case val > 0:
			fmt.Println("You entered:", val)
		case val == 0:
			fmt.Println("0 is neither negative nor positive")
		case val < 0:
			panic(val)
		default:
			break
		}
	}

}

func checkNumb(i int) int {
	return 1 + (i >> 31) - (-i >> 31)
}
