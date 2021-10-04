package main

import "fmt"

func main() {
	n := "zhangsan"
	a := 13
	g := true
	convertType(n)
	convertType(a)
	convertType(g)
}

func convertType(t interface{}) {
	fmt.Printf("val:%v type:%T\n", t, t)
}
