package main

import (
	"fmt"
	"time"
)


func printMsg(c chan string) {
	for msg := range c {
		fmt.Println(msg)
	}
}

func main() {
	msgs := make(chan string, 2)
	msgs <- "welcome to"
	msgs <- "gopher's world"
	close(msgs)
	fmt.Println("sent two messages")
	time.Sleep(time.Second * 3)
	printMsg(msgs)
}