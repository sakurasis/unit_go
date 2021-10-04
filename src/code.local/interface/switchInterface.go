package main

import "fmt"

func checkType(x interface{}) {

	switch v := x.(type) {
	case string:
		fmt.Printf("the type is a string, value is %v\n", v)
	case int:
		fmt.Printf("the type is a int, value is %v\n", v)
	case bool:
		fmt.Printf("the type is a bool, value is %v\n", v)
	default:
		fmt.Printf("the unknown type.")
	}
}

func main() {
	m := "zhangsan"
	checkType(m)
	n := 123456
	checkType(n)
	k := true
	checkType(k)
}
