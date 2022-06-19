/*
author:admin
createTime:
*/
package main

import "fmt"

// quickSort the sorted algorithm that separate it into two partitions,
// get the sorted with recurrence method.
// l the index of the first element
// r the index of the last element
func quickSort(arr []int, l int, r int) {
	if l >= r {
		return
	}
	p := getStandardIndex(arr, l, r)
	quickSort(arr, l, p-1)
	quickSort(arr, p+1, r)
}

// getStandardIndex get the pivot's index, actually.
func getStandardIndex(arr []int, l int, r int) int {
	// todo
	pivot := arr[r]
	idx := l
	for i := l; i < r; i++ {
		if arr[i] < pivot {
			arr[idx], arr[i] = arr[i], arr[idx]
			idx++
		}
	}
	arr[idx], arr[r] = pivot, arr[idx]
	return idx
}

func main() {
	arr := []int{6, 5, 3, 1, 8, 7, 2, 4}
	quickSort(arr, 0, len(arr)-1)
	for _, v := range arr {
		fmt.Printf("%d\t", v)
	}
}
