/*
author:admin
createTime: 2022-06-17
*/
package main

import "fmt"

func selectionSort(arr []int) {
	if len(arr) <= 1 {
		return
	}
	for i := 0; i < len(arr)-1; i++ {
		idx := 0
		for j := 0; j < len(arr)-i; j++ {
			// get the index of the biggest number.
			if arr[j] > arr[idx] {
				idx = j
			}
		}
		temp := arr[idx]
		// the last value
		arr[idx] = arr[len(arr)-1-i]
		arr[len(arr)-1-i] = temp
	}
}

func main() {
	arr := []int{6, 5, 3, 1, 8, 7, 2, 4}
	selectionSort(arr)
	for _, v := range arr {
		fmt.Printf("%d\t", v)
	}
}
