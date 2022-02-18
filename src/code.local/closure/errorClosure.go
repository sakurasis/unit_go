package main

import (
	"errors"
	"fmt"
)

func testError(a, b int) (i int, err error) {
	defer func(err error) { fmt.Printf("second defer err %v\n", err) }(err)
	defer fmt.Printf("first defer err %v\n", err)
	defer func() { fmt.Printf("third defer err %v\n", err) }()
	if b == 0 {
		err = errors.New("divided by zero!")
		return
	}

	i = a / b
	return
}

/*

 */
func main() {
	testError(2, 0)
}
