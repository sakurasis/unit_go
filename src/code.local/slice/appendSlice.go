package main

import "fmt"

func main() {
	var a = make([]string, 5, 10)
	for i := 0; i < len(a); i++ {
		fmt.Printf("add:%p, len:%v, cap:%v\n", &a[i], len(a), cap(a))
	}
	for i := 0; i < 10; i++ {
		a = append(a, fmt.Sprintf("%v", i))
		fmt.Printf("add:%p, len:%v, cap:%v\n", &a[i], len(a), cap(a))
	}

	fmt.Println(a)

}
