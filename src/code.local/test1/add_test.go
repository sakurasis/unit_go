/*
author:admin
createTime:
*/
package main

import "testing"

func TestAdd(t *testing.T) {
	sum := Add(1, 2)
	if sum == 3 {
		t.Log("the result is ok")
	} else {
		t.Fatal("the result is wrong")
	}
}

func Add(i int, i2 int) int {
	return i + i2
}

func BenchmarkRemoveEles(b *testing.B) {
	b.ResetTimer()
	numbs := []int{3, 2, 2, 3}
	for i := 0; i < b.N; i++ {
		removeElements(numbs, 3)
	}
}

// removeElements remove the element which input the number,
// back to a new array that exclude it.
func removeElements(arr []int, k int) []int {
	for i := 0; i < len(arr); {
		if arr[i] == k {
			arr = append(arr[:i], arr[i+1:]...)
		} else {
			i++
		}
	}
	return arr
}
