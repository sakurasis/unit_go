package main

import (
	"fmt"
	"time"
)

func thread(str string) {
	for i := 0; i < 3; i++ {
		fmt.Println(str, ":", i)
	}
}

func main() {
	thread("thread1")

	go thread("routine")

	go func(msg string) {
		fmt.Println(msg)
	}("go,go,go...")

	time.Sleep(time.Second)
	fmt.Println("done...")
}