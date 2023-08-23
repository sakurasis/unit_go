package main

import (
	"fmt"
	"github.com/google/uuid"
	"regexp"
)

func main() {
	reg := regexp.MustCompile(`[^0-9]+`)
	s, _ := uuid.NewRandom()
	all := reg.ReplaceAllString(s.String(), "")
	fmt.Println(s)
	fmt.Println("------------------")
	fmt.Println(all)
}
