package main

import "fmt"

func f1() int {
	x := 5
	defer func() {
		x++
	}()
	return x
}

//
//func f2() (x int) {
//	fmt.Println("x:", x)
//	defer func() {
//		x++
//	}()
//	fmt.Println("defer x:", x)
//	return 5
//}
//
//func f3() (y int) {
//	x := 5
//	defer func() {
//		x++
//	}()
//	return x
//}
//func f4() (x int) {
//	fmt.Println("x:", x)
//	defer func(x int) {
//		x++
//	}(x)
//	fmt.Println("defer x:", x)
//	return 5
//}
func main() {
	fmt.Println(f1()) // 5
	//fmt.Println(f2()) // 6
	//fmt.Println(f3()) // 5
	//fmt.Println(f4())
}