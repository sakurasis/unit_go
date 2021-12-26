package main

func GetAndIncrease() func() int {
	v := 0
	return func() int {
		v++
		return v
	}
}

func main() {
	val := GetAndIncrease()
	println(val()) // 1
	println(val()) // 2
}
