package main

import (
	"flag"
	"fmt"
	"strings"
)

var n = flag.Bool("n", false, "new line")
var separ = flag.String("s", " ", "separate")

func main() {
	flag.Parse()
	fmt.Print(strings.Join(flag.Args(), *separ))
	if !*n {
		fmt.Println("--------")
	}
}