package employee

import "fmt"

type Employee struct {
	FirstName string
	LastName  string
	Age       int
	Gender    bool
}

func (e Employee) ToString() {
	fmt.Printf("Employee: {firstName:%s, lastName:%s, age:%d, gender:%v}\n", e.FirstName, e.LastName, e.Age, e.Gender)
}
