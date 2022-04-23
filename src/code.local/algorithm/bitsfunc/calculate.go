package main

import "fmt"

func main() {
	println(add(1, 3))
	println(sub(1, 3))
	println(sub1(1, 3))
	println(mul(2, 3))
	println(div(2, 0))
	println(div(4, 2))
}

//add the sum of x and y
func add(x, y int) int {
	if y != 0 {
		return add(x^y, (x&y)<<1)
	} else {
		return x
	}
}

// sub the value of which is y subtracted from x
func sub(x, y int) int {
	return add(x, -y)
}

func sub1(x, y int) int {
	return add(x, add(^y, 1))
}

func mul(x, y int) int {
	var n, val int
	for y != 0 {
		if (y & 1) == 1 {
			val += x << n
			y = y >> 1
			n++
		} else {
			y = y >> 1
			n++
		}
	}
	return val
}

func div(x, y int) int {
	var val int
	if y == 0 {
		fmt.Printf("error number：%d\n", y)
		return -1
	}
	for i := 31; i > 0; i-- {
		if x>>i >= y {
			val = add(val, 1<<i)
			x = sub(x, y<<i)
		}
	}
	if x^y < 0 {
		val = add(^val, 1)
	}
	return val
}
