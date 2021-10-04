package main

import "fmt"

func main() {
	printHW()
}

func printHW() {
	var str string = "Hello world"
	fv := func(vars string) {
		for i := 0; i < len(vars); i++ {
			fmt.Printf("%c", vars[i])
		}
		fmt.Println("")
	}
	fv(str)
	fmt.Println("It's right?")

}
