/*
author:admin
createTime: 2022/01/28 14:50
*/
package main

import "fmt"

func main() {
	numbs := []int{3, 2, 2, 3}
	elements := removeElements(numbs, 3)
	fmt.Printf("%v", elements)
}

// removeElements remove the element which input the number,
// back to a new array that exclude it.
func removeElements(arr []int, k int) []int {
	for i := 0; i < len(arr); {
		if arr[i] == k {
			arr = append(arr[:i], arr[i+1:]...)
		} else {
			i++
		}
	}
	return arr
}
