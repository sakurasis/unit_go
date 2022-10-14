package main

import "fmt"

func main() {
	arr := []int{6, 5, 3, 1, 8, 7, 2, 4}
	arr = shellsort(arr)
	fmt.Println(arr)
}

// shellsort split into some several arrays with same length.
func shellsort(arr []int) []int {
	// interval value
	val := 1
	// change the value dynamically
	if val < len(arr)/3 {
		val *= 3
		val++
	}
	// the limited condition of loop.
	for val > 0 {
		for i, v := range arr {
			preIndex := i - 1
			for preIndex >= 0 && arr[preIndex] > v {
				arr[preIndex+1] = arr[preIndex]
				preIndex -= 1
			}
			arr[preIndex+1] = v
		}
		// return the original value.
		val /= 3
	}
	return arr
}
