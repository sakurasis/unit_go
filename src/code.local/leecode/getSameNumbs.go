//go:build ignore
// +build ignore

package main

import (
	"fmt"
	"os"
	"runtime/trace"
	"sort"
)

func main() {
	//创建trace文件
	f, err := os.Create("trace.out")
	if err != nil {
		panic(err)
	}

	defer f.Close()

	//启动trace goroutine
	err = trace.Start(f)
	if err != nil {
		panic(err)
	}
	defer trace.Stop()
	arr1 := []int{1, 2, 2, 1, 4}
	arr2 := []int{2, 2, 4}
	same := getSameTwo(arr1, arr2)
	fmt.Printf("%+v\n", same)
}

func getSameTwo(arr1 []int, arr2 []int) []int {
	i, j, k := 0, 0, 0
	if arr1 == nil || arr2 == nil {
		return nil
	}
	sort.Ints(arr1)
	sort.Ints(arr2)
	for i < len(arr1) && j < len(arr2) {
		if arr1[i] > arr2[j] {
			j++
		} else if arr1[i] < arr2[j] {
			i++
		} else {
			arr1[k] = arr1[i]
			i++
			j++
			k++
		}
	}
	return arr1[:k]
}
