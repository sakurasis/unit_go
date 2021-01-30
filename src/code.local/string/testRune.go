package main

import (
	"fmt"
	"strconv"
)


func main() {
	str := `After a backslash, certain single character escapes represent 
special values
	n is a line feed or new line
	t is a tab`
	fmt.Println(str)
	// concat string
	ansStr := "Q:what's your name?\n"
	ansStr += "A:what did you say?"
	fmt.Println(ansStr)
	// casting to string
	fmt.Println("===========cast to string===============")
	num := 2020
	var msg string = "now is "
	numStr := strconv.Itoa(num)
	fmt.Println(msg + numStr)
}