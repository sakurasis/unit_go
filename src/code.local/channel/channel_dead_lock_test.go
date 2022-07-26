package main

import (
	"fmt"
	"testing"
)

func TestUnbufferedChannel(t *testing.T) {
	data := make(chan int, 1)

	data <- 6
	go func() {
		fmt.Println(<-data)
	}()

}

func TestChannelNil(t *testing.T) {
	var data chan int

	<-data
}

func TestBufferedChannel(t *testing.T) {
	c := make(chan int, 1)

	<-c
}
