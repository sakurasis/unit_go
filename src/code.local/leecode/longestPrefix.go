//go:build ignore
// +build ignore

/*
author:admin
createTime:
*/
package main

import (
	"sort"
	"strings"
)

func main() {
	// slices := []string{"analytical", "analysis", "analyse"}
	println("--------------------------")
	var idxs = []string{"anmial", "analyse", "analyze"}
	prefix := getLongestPrefix(idxs)
	println(prefix)
}

// getLongestPrefix get the longest prefix string throught using
// the function which name is index import from the package strings.
func getLongestPrefix(str []string) string {
	if len(str) <= 0 {
		return ""
	}
	sort.Strings(str)
	prefix := str[0]
	// loop the slice.
	if len(prefix) == 0 {
		panic("prefix can't be empty characters.")
	}
	for _, v := range str {
		for strings.Index(v, prefix) != 0 {
			if len(prefix) == 0 {
				panic("prefix can't be empty characters.")
			}
			prefix = prefix[:len(prefix)-1]
		}
	}
	return prefix
}
