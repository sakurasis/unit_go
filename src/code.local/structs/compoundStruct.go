package structs

import "fmt"

type student struct {
	name string
	age  int
}

func main() {
	m := make(map[string]*student)
	stus := []student{
		{name: "大王八", age: 9000},
		{name: "小王子", age: 18},
		{name: "娜扎", age: 23},
	}

	for index, stu := range stus {
		fmt.Printf("%v\n", stu.name)
		fmt.Printf("%v\n", stu)
		m[stu.name] = &stus[index]
	}

	fmt.Printf("m:%#v\n", m)
	for k, v := range m {
		fmt.Println(k, "=>", v)
	}
}
