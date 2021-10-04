package main

import (
	"fmt"
	"strings"
)

const N = 5

func main() {
	str := "Download and start using DBeaver Enterprise edition right now "
	// to lower format the string
	str = strings.ToLower(str)
	fmt.Println("lower func:", str)
	andStr := strings.Index(str, "and")
	fmt.Println("and's index:", andStr)
	newStr := strings.TrimSpace(str)
	fmt.Println("str:", str, "length: ", len(str))
	fmt.Println("newStr:", newStr, "length: ", len(newStr))
}
