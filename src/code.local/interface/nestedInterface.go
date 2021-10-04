package main

import "fmt"

type Eater interface {
	Eat()
	Sleep()
}

type Humaner interface {
	Eater
}

type Student struct {
	Name string
}

func (s *Student) Eat() {
	fmt.Println(s.Name + " eating at the moment.")
}

// func (s *Student) Sleep() {
// 	fmt.Println(s.Name + " go to bed at night.")
// }

func main() {
	var s Humaner = &Student{"Anna"}
	s.Eat()
	//s.Sleep()
}
