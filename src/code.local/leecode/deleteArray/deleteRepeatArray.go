/*
author:admin
createTime:
*/
package main

func main() {
	var arr = []int{1, 9, 9}
	deleteRepeatElement(arr)
}

// deleteRepeatElement remove same Elements from the array.
func deleteRepeatElement(arr []int) {
	for i := len(arr) - 1; i >= 0; i-- {
		n := (arr[i] + 1) / 10
		if n > 0 {
			// 进一位

		} else {
			// 只进行加1操作

		}
	}
}
