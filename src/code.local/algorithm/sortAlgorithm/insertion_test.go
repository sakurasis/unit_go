/*
author:admin
createTime:
*/
package main

import "testing"

var arr = []int{6, 5, 3, 1, 8, 7, 2, 4}

func BenchmarkInsertionSort(b *testing.B) {
	for i := 0; i < b.N; i++ {
		arr = InsertionSort(arr)
	}
}

func BenchmarkInsertion2Sort(b *testing.B) {
	for i := 0; i < b.N; i++ {
		arr = insertionSort1(arr)
	}
}

func BenchmarkBubbleSort(b *testing.B) {
	for i := 0; i < b.N; i++ {
		arr = bubbleSort(arr)
	}
}

func BenchmarkSelectionSort(b *testing.B) {
	for i := 0; i < b.N; i++ {
		selectionSort(arr)
	}
}
