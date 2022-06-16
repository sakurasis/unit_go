/*
author:admin
createTime: 2022-06-16 16:30
*/
package main

import "fmt"

func bubbleSort(arr []int) []int {

	if len(arr) <= 1 {
		return arr
	}

	for i := 0; i < len(arr)-1; i++ {
		for j := 0; j < len(arr)-i-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
	return arr
}

func main() {
	arr := []int{6, 5, 3, 1, 8, 7, 2, 4}
	arr = bubbleSort(arr)
	for _, v := range arr {
		fmt.Printf("%d\t", v)
	}

}
