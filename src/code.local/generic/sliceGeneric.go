package main

import "fmt"

func assembleSlice[T any](s []T)  {
	for _, v := range s {
		_, _ = fmt.Printf("%v\t", v)
	}
	_, _ = fmt.Print("\n")
}


func main() {
	assembleSlice[int]([]int{1,2,3,4})
	assembleSlice[string]([]string{"champion", "edg"})
}
