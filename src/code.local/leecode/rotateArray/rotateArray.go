/*
author:admin
createTime:
*/
package main

import "fmt"

func main() {
	var k = 2
	arr := []int{1, 3, 4, 5, 7, 2, 3, 8}

	rotateArr(arr)
	fmt.Printf("%v\n", arr)
	rotateArr(arr[:k%len(arr)])
	fmt.Printf("%v\n", arr)
	rotateArr(arr[k%len(arr):])
	fmt.Printf("%v\n", arr)
}

// rotateArr rotate the array to get the newest array.
func rotateArr(arr []int) {
	for i := 0; i < len(arr)/2; i++ {
		arr[i], arr[len(arr)-i-1] = arr[len(arr)-i-1], arr[i]
	}
}
