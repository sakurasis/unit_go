package main

import "fmt"

func send(ch chan<- string, msg string) {
	fmt.Printf("Sending: %#v\n", msg)
	ch <- msg
}

func receive(ch <-chan string) {
	fmt.Printf("Receiving: %#v\n", <-ch)
}

func main() {
	ch := make(chan string, 1)
	send(ch, "hello, gophers.")
	receive(ch)
}
