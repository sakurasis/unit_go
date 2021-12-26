package main

import (
	"fmt"
	"time"
)

var quit = make(chan bool)

// the function is that fibonacci sequence can use by calculating
// the sum of the previous number and the next number
func fib(ch chan int) {
	x, y := 1, 1
	for {
		select {
		case ch <- x:
			x, y = y, x+y
		case <-quit:
			fmt.Println("Exiting the progress.")
			return
		}
	}
}

func main() {
	now := time.Now()
	var com string
	data := make(chan int)

	for {
		num := <-data
		fmt.Println(num)
		_, err := fmt.Scanf("%s", &com)
		if err != nil {
			return
		}

		if com == "quit" {
			quit <- true
			break
		}
	}
	between := time.Since(now)
	fmt.Printf("It tooks %v seconds in codes.\n", between.Seconds())
}
