package main

import "fmt"

type Student struct {
	id   int32
	name string
	age  int8
}

func NewInstance(id int32, name string, age int8) *Student {
	return &Student{
		id:   id,
		name: name,
		age:  age,
	}
}

func (s Student) todoList() {
	fmt.Printf("%s's age is %d\n", s.name, s.age)
}

func (s *Student) setName(newName string) {
	s.name = newName
}

func (s Student) setName2(newName string) {
	s.name = newName
}

func (s *Student) getName() string {
	return s.name
}

func main() {
	stu := NewInstance(132, "edan", 18)
	stu.todoList()
	fmt.Println("the name is", stu.name)
	stu.setName("miya")
	miya := stu.getName()
	fmt.Printf("the name change to be %s now\n", stu.name)
	stu.setName2("tompson")
	tompson := stu.getName()
	stu.setName("smith")
	smith := stu.getName()
	fmt.Printf("the name was changed ? %v\n", miya != tompson)

	fmt.Printf("the next name was change? %v , the name is %s\n", smith != miya, smith)
}
