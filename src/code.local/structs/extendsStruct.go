package structs

import "fmt"

type Person struct {
	name   string
	gender bool
	age    int8
}

type Student struct {
	className string
	roomName  string
	Person
}

type Worker struct {
	workingArea string
	Person
}

func (p *Person) work() {
	fmt.Printf("%s working in a company.\n", p.name)
}

func (s *Student) homework() {
	fmt.Printf("%s working in a school.\n", s.name)
}

func main() {
	stu1 := &Student{
		className: "classOne",
		roomName:  "room 202",
		Person: Person{
			name:   "tompson",
			gender: true,
			age:    8,
		},
	}
	stu1.work()
	stu1.homework()
}
