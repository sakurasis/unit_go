package main

import (
	"fmt"
	"os"
)

func main() {
	var str, step string

	for i := 0; i < len(os.Args); i++ {
		str += step + os.Args[i]
		step = " "
	}
	fmt.Println(str)
}
