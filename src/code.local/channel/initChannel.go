package main

import (
	"fmt"
	"time"
)

func slowFunc(c chan string) {
	time.Sleep(time.Second * 2)
	c <- "the channel A finished"
}

func main() {
	c := make(chan string)
	go slowFunc(c)

	msg := <-c
	fmt.Println("the message:", msg)
}
