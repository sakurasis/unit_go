package main

func main() {
	r, _ := multiple(8, 9)
	println("result:", r)
}

//multiply the function is to calculate the result of two number which type is integer.
func multiple(x, y int) (r int, f bool) {
	z := x * y
	return z, true
}
