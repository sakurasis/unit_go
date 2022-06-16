package main

import (
	"fmt"
	"time"
)

func runGoroutine(name string, numb int) <-chan string {
	cha := make(chan string)
	go func() {
		for i := 0; i < 10; i++ {
			cha <- fmt.Sprintf("routine %s; sleep:%d", name, numb)
			time.Sleep(time.Duration(numb) * time.Millisecond)
		}
	}()

	return cha
}

func runChannel(ipt1, ipt2 <-chan string) <-chan string {
	channel := make(chan string)
	go func() {
		select {
		case channel <- <-ipt1:
		case channel <- <-ipt2:
		}
	}()
	return channel
}

func main() {
	channel := runChannel(runGoroutine("A", 10), runGoroutine("B", 1000))
	for i := 0; i < 10; i++ {
		fmt.Printf("%q\n", <-channel)

	}
}
