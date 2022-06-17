/*
author:admin
createTime: 2022-06-16
*/
package main

import "fmt"

// InsertionSort insert the special value,
// and then compare with the next value from the array.
func InsertionSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}
	var temp int
	for i := 1; i < len(arr); i++ {
		temp = arr[i]
		j := i - 1
		// the condition that compare two numbers between the one with the next.
		for j >= 0 && arr[j] > temp {
			arr[j+1] = arr[j]
			j--
		}
		arr[j+1] = temp
	}
	return arr
}

// insertionSort1 the different way to get the result.
func insertionSort1(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}
	for i := 0; i < len(arr); i++ {
		var temp = arr[i]
		j := i - 1
		for ; j >= 0 && arr[j] > temp; j-- {
			arr[j+1] = arr[j]
		}
		// 1.the solution which the index is zero that can't go into the second loop.
		// 2.the solution which the index is bigger than zero, it will into the loop, and get the newest number.
		arr[j+1] = temp
	}
	return arr
}

func main() {
	arr := []int{6, 5, 3, 1, 8, 7, 2, 4}
	arr = insertionSort1(arr)
	for _, v := range arr {
		fmt.Printf("%d\t", v)
	}
}
