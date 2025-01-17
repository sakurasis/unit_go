package main

import "fmt"

func counter(out chan<- int) {
	for i := 0; i < 100; i++ {
		out <- i + 1
	}
	close(out)
}

func squarer(out chan<- int, in <-chan int) {
	for i := range in {
		out <- i * i
	}
	close(out)
}
func printer(in <-chan int) {
	for i := range in {
		fmt.Println(i)
	}
}

func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)
	go counter(ch1)
	printer(ch1)
	go squarer(ch2, ch1)
}
