/*
author:admin
createTime: 2022-06-16
*/
package main

import "fmt"

func insertionSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}
	for i := 0; i < len(arr); i++ {
		val := arr[i]
	}
	return arr
}

func main() {
	arr := []int{6, 5, 3, 1, 8, 7, 2, 4}
	arr = insertionSort(arr)
	for _, v := range arr {
		fmt.Printf("%d\t", v)
	}
}
