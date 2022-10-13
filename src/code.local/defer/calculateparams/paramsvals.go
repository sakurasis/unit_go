package main

import "fmt"

func deferInloop() {
	for i := 0; i < 5; i++ {
		defer fmt.Println(i)
	}
}
func main() {

	deferInloop()
}
