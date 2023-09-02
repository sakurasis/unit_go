package main

import (
	"fmt"
	"time"
)

func pingChan1(c chan string) {
	time.Sleep(time.Second * 2)
	c <- "ping channel1"
}

func pingChan2(c chan string) {
	time.Sleep(time.Second * 1)
	c <- "ping channel2"
}

func main() {
	channel1 := make(chan string)
	channel2 := make(chan string)

	go pingChan1(channel1)
	go pingChan2(channel2)

	select {
	case info1 := <-channel1:
		fmt.Printf("%s received.", info1)
	case info2 := <-channel2:
		fmt.Printf("%s received.", info2)
	// set the over time to print message in console.
	case <-time.After(5 * time.Second):
		fmt.Println("no message can be received.")
	}
	fmt.Println("-------------------------------")
	s := make(chan int)
	e := make(chan int)

	go func() {
		for i := 0; i < 10 ; i ++ {
			fmt.Println(<-s)
		}
		e <- 0
	}()

	fibonacii(s, e)
}

func fibonacii(s, e chan int) {
	x, y := 1, 1
	for {
		select {
		case s <- x :
			x  = y
			y = x + y

		case <-e :
			return
		}
	}
}
