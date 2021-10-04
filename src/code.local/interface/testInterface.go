package main

import "fmt"

type TestObj struct{}

func checkNil(t interface{}) bool {
	return t == nil
}

func main() {
	var s *TestObj
	fmt.Println(s == nil)
	fmt.Println(checkNil(s))
}
