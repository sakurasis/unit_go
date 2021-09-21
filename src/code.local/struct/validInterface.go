package main

import "fmt"

func main() {
	// 定义一个空接口x
	var car interface{}
	lincoin := "Hello lincoin"
	car = lincoin
	fmt.Printf("type:%T value:%v\n", car, car)
	price := 1000
	car = price
	fmt.Printf("type:%T value:%v\n", car, car)
	driven := true
	car = driven
	fmt.Printf("type:%T value:%v\n", car, car)
}