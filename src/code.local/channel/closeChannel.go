package main

import (
	"fmt"
	"time")


func sendMsg(c chan string) {
	t := time.NewTicker(1 * time.Second)
	for {
		c <- "sending a message"
		<- t.C
	}
}

func stopTime(flag chan bool) {
	time.Sleep(2 * time.Second)
	fmt.Println("stop it......")
	flag <- true
}

func main() {
	info := make(chan string)
	flagStop := make(chan bool)
	go sendMsg(info)
	go func() {
		time.Sleep(2 * time.Second)
		fmt.Println("stop it......")
		flagStop <- true
	} ()
	
	for {
		select {
		case <- flagStop: 
			return
		case msg := <-info:
			fmt.Println(msg)
		}
	}
}