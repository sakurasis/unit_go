package main

import (
	"fmt"
	"time"
	"strconv"
)


func pingChan(c chan string) {
	t :=time.NewTicker(1 * time.Millisecond)
	for i:=0;;i ++{
		c <- ("ping " + strconv.Itoa(i) + " times")
		<- t.C
	}
}

func readChannel(msg <-chan string) {
	info := <-msg
	fmt.Println(info)
}

func writeChannel(msg chan<- string) {
	msg <- "Hello gophers"
}

func readAndWriteChannel(msg chan string) {
	info := <-msg
	fmt.Println(info)
	msg <- "hello gophers!"
}

func main() {
	msgs := make(chan string)
	writeChannel(msgs)
	readChannel(msgs)
	// go pingChan(msgs)
	// for {
	// 	msg := <-msgs
	// 	fmt.Println(msg)
	// }
}