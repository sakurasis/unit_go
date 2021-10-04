package main

import "fmt"

type Student struct {
	Id     int
	Name   string
	Gender bool
}

type Beaner interface {
	GetName() string
	SetName(string)
}

func (s Student) GetName() string { // 1
	return s.Name
}

func (s *Student) SetName(name string) { // 2
	s.Name = name
}

func main() {
	// stu := Student{1, "lisi", true} // 3 类型：main.Student
	// stu.SetName("zhangsan")
	// fmt.Println(stu.GetName(), stu)

	var stu Beaner = &Student{1, "lisi", true}
	stu.SetName("zhangsan")
	fmt.Println(stu.GetName(), stu)
}
