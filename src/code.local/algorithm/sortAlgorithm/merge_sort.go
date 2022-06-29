/*
author:admin
createTime:
*/
package main

import "fmt"

// mergeSort separate the array into two spices.
func mergeSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}
	part := len(arr) / 2
	l := mergeSort(arr[0:part])
	r := mergeSort(arr[part:])
	return sortOp(l, r)
}

func sortOp(left, right []int) []int {
	i, j := 0, 0
	// the length of one side
	m, n := len(left), len(right)
	var arrays []int
	for {
		if i >= m || j >= n {
			break
		}
		// compare with two number between previous and next one.
		if left[i] <= right[j] {
			arrays = append(arrays, left[i])
			i++
		} else {
			arrays = append(arrays, right[j])
			j++
		}
	}
	// append last value on left side.
	if i != m {
		for ; i < m; i++ {
			arrays = append(arrays, left[i])
		}
	}
	// append last value on right side.
	if j != n {
		for ; j < n; j++ {
			arrays = append(arrays, right[j])
		}
	}
	return arrays
}

func main() {
	arr := []int{6, 5, 3, 1, 8, 7, 2, 4}
	arr = mergeSort(arr)
	for _, v := range arr {
		fmt.Printf("%d\t", v)
	}
}
