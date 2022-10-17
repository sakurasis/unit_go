package main

import "fmt"

func main() {
	arr := []int{6, 5, 3, 1, 8, 7, 2, 4}
	heapSort(arr)
	fmt.Println(arr)
}

func heapSort(arr []int) {
	// build the maxmium heap
	buildMaxHeap(arr)
	// the array's length
	l := len(arr)
	for i := 0; i < l; i++ {
		swap(arr, 0, l-1-i)
		heapify(arr, 0, l-1-i)
	}
}

// buildMaxHeap build the max heap.
func buildMaxHeap(arr []int) {
	for i := (len(arr) - 2) / 2; i >= 0; i-- {
		heapify(arr, i, len(arr))
	}
}

// heapify build a struct of heap.
func heapify(arr []int, index, len int) {
	lnode := 2*index + 1
	rnode := 2*index + 2
	pnode := index

	// left child's node > parent's node
	if lnode < len && arr[lnode] > arr[pnode] {
		pnode = lnode
	}
	// right child's node > parent's node.
	if rnode < len && arr[rnode] > arr[pnode] {
		pnode = rnode
	}
	// the position of parent's node is not the index of current node
	// the largest value give to the parent's node
	if pnode != index {
		swap(arr, index, pnode)
		heapify(arr, pnode, len)
	}
}

// Swap two adjacent values
func swap(arr []int, i, j int) {
	arr[i], arr[j] = arr[j], arr[i]
}
