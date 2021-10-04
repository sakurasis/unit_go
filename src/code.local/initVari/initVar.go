package main

import (
	"./trans"
	"fmt"
	"math"
	"os"
)

func main() {
	// var (
	// 	HOME = os.Getenv("PATH")
	// )
	host, err := os.Hostname()
	// fmt.Println("the go home is", HOME)
	if err != nil {
		fmt.Printf("a error: %s\n", err)
	} else {
		fmt.Println("the host is", host)
	}
	// arctan(1) = pi/4
	fmt.Printf("%.10f", 4*math.Atan(1))

	pi2 := 2 * trans.Pi
	fmt.Printf("pi is %.3f", pi2)
}
