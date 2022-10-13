//go:build ignore
// +build ignore

// go build ignore
package main

import (
	"fmt"
	"sort"
)

func main() {
	var ages []string
	names := map[string]int{
		"adana":  18,
		"carry":  12,
		"bob":    20,
		"venius": 30,
	}
	for name := range names {
		ages = append(ages, name)
	}
	sort.Strings(ages)

	for _, age := range ages {
		fmt.Printf("%s\t%d\n", age, names[age])
	}
}
