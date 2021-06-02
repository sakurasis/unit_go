package main

import "fmt"

func main() {
	p := new(int)
	fmt.Println(*p)

	*p = 2
	fmt.Println(p)

	var x = 1 << 1 | 1 << 5
	fmt.Printf("%08b\n", x) 
	
	
	fmt.Println("**************************")

	var num1 float64 = 100000000000000.0 / 3

	var num2 float32 = 100.0 / 3 

	fmt.Printf("%v\n", num1)
	fmt.Printf("%v\n", num2)
	fmt.Println("-------------------------")
	fmt.Printf("%.3f\n", num1)
	fmt.Printf("%.3f\n", num2)
	fmt.Println("-------------------------")
	fmt.Printf("%4.2f\n", num1)
	fmt.Printf("%4.2f\n", num2)
	fmt.Println("-------------------------")


}                                                             