package main

import (
	"fmt"
	"time"
)

func process(ch chan string) {
	time.Sleep(3 * time.Second)
	ch <- "processing!"
}

func replicate(ch chan string) {
	time.Sleep(1 * time.Second)
	ch <- "replicating"
}

func main() {
	channel1 := make(chan string)
	channel2 := make(chan string)
	go process(channel1)
	go replicate(channel2)

	for i := 0; i < 2; i++ {
		select {
		case p1 := <-channel1:
			fmt.Println(p1)
		case r1 := <-channel2:
			fmt.Println(r1)
		}

	}
}
