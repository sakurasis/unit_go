package main

import (
	"fmt"
	"github.com/golang/example/stringutil"
)

func main() {
	str := "Hybrid Transactional and Analytical Processing"
	reverseStr := stringutil.Reverse(str)
	fmt.Println(reverseStr)
}

