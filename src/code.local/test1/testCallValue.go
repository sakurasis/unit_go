package main

func getValue1(x int) int {
	return x
}

func getValue2(x *int) int {
	println("the x value:", x)
	return *x
}

var area = func(a, b int) int {
	return a * b
}

func main() {
	v1 := getValue1(255)
	println("getValue1:", v1, "the addr:", &v1)
	v2 := getValue2(&v1)
	println("getValue2:", v2)
	a := 255
	b := 1
	println(area(a, b))
}
