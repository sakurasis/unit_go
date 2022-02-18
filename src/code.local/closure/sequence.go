/*
author:admin
createTime:
*/
package main

func main() {
	sequence := getSequence()
	println(sequence())
	println(sequence())
	println(sequence())
	println(sequence())
}

func getSequence() func() int {
	var i = 0
	return func() int {
		i++
		return i
	}
}
