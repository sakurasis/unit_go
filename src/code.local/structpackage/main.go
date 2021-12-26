package main

import "main/employee"

func main() {
	var e = employee.Employee{FirstName: "lee", LastName: "lin", Age: 18, Gender: true}
	e.ToString()
}
