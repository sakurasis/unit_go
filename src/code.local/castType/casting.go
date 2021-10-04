package main

import "fmt"

func main() {
	var b int16 = 4

	var c int32

	c = int32(b)

	fmt.Printf("int in 32 bit is: %d\n", c)
	fmt.Printf("int in 16 bit is: %d\n", c)

	var complexNum complex64 = 5 + 10i

	d := complex(4, 3)
	fmt.Printf("the complex number is %v\n", complexNum)
	fmt.Printf("the d value is %v\n", d)

	var revertNum = ^10
	fmt.Printf("revert number is %d\n", revertNum)
}
