package main

import (
	"fmt"
	"sort"
)

// an algorithm of the bucket's sort
func main() {
	arr := []int{6, 5, 3, 1, 8, 7, 2, 4}
	bucketsSort(arr, 3)
	fmt.Println(arr)
}

func bucketsSort(arr []int, bucketSize int) {
	min, max := arr[0], arr[0]
	for _, v := range arr {
		if v < min {
			min = v
		}
		if v > max {
			max = v
		}
	}
	// setElementsToBucket put every elements to the bucket.
	bucketCount := (max-min)/bucketSize + 1
	twoDimArr := make([][]int, bucketCount)
	for i := 0; i < len(arr); i++ {
		idx := (arr[i] - min) / bucketSize
		twoDimArr[idx] = append(twoDimArr[idx], arr[i])
	}
	// the value's sort of some child sequences.
	idx := 0
	for _, v := range twoDimArr {
		sort.Ints(v)
		for _, k := range v {
			arr[idx] = k
			idx++
		}
	}
}
