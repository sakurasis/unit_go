/*
author:admin
createTime:
*/
package main

import "testing"

func TestGoroutine(t *testing.T) {
	c := make(chan int)
	c <- 666
	//<-c
}

func TestBeforeGoroutine(t *testing.T) {
	c := make(chan int)
	c <- 666
	go func() {
		<-c
	}()
}

func TestCallbackEach(t *testing.T) {
	c1, c2 := make(chan int), make(chan int)
	go func() {
		for {
			select {
			case <-c1:
				c2 <- 10
			}
		}
	}()
	for {
		select {
		case <-c2:
			c1 <- 10
		}
	}
}
