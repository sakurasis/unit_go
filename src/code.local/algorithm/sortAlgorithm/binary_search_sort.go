package main

import (
	"fmt"
	"sort"
)

func binarySearchSort(arr []int, key int) int {
	var low int = 0
	var high int = len(arr) - 1
	for low <= high {
		var mid int = (low + high) >> 2
		var midVal int = arr[mid]
		if key > midVal {
			low = mid + 1
		} else if key < midVal {
			high = mid - 1
		} else {
			return mid
		}
	}
	return -(low + 1)
}

func main() {
	arr := []int{6, 5, 3, 1, 8, 7, 2, 4}
	sort.Ints(arr)
	key := 2
	idx := binarySearchSort(arr, key)

	for _, v := range arr {
		fmt.Printf("%d\t", v)
	}
	if idx != -1 {
		fmt.Printf("Find num %d at index %d\n", key, idx)
	} else {
		fmt.Printf("Num %d not exists in nums\n", key)
	}
}
