package main

import (
	"fmt"
	"io/ioutil")


func main() {
	file, err := ioutil.ReadFile("cutegirl.jpg")
	if err != nil {
		fmt.Println("err:", err)
	} else {
		fmt.Println(file)
	}
}