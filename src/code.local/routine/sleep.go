package main

import (
	"fmt"
	"time"
)
func sleepFunc() {
	now := time.Now()
	time.Sleep(time.Second * 2)
	next := time.Now()
	fmt.Println("finish the code. It spend", (next.Sub(now)))
}

func main() {
	fmt.Println("start running....")
	go sleepFunc()
	time.Sleep(time.Second * 3)
	fmt.Println("thread main finished.")
}