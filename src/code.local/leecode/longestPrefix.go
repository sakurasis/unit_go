package main

import "strings"

func main() {
	slices := []string{"analyse", "analysis"}
	prefix := longestPrefix(slices)
	println(prefix)
}

//longestPrefix  get the longest prefix of string
func longestPrefix(strs []string) string {
	if len(strs) <= 0 {
		return ""
	}

	prefix := strs[0]
	for _, v := range strs {
		for strings.Index(v, prefix) != 0 {
			if len(prefix) == 0 {
				return ""
			}
			prefix = prefix[:(len(prefix) - 1)]
		}
	}
	return prefix
}
