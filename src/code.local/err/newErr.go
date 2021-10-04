package main

import (
	"errors"
	"fmt"
)

func main() {
	err := errors.New("throw new Errors")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("=================================")
	panic("stop now, it can't execute next code again.")
	name, job := "Michael Jackson", "Singler"
	errStr := fmt.Errorf("The %v is a famous %v", name, job)
	if errStr != nil {
		fmt.Println(errStr)
	}
}
