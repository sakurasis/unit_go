package main

import (
	"fmt"
	"log"
	"os"
	"testing"
)

type Employee struct {
	Id         float32
	FirstName  string
	MiddleName string
	LastName   string
}

func TestError(t *testing.T) {
	file, err := os.OpenFile("info.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			log.Print(err)
		}
	}(file)

	log.SetOutput(file)
	log.SetPrefix("DEBUG:")
	_, _ = getInfo(1000)
	log.Println("you can help it")
	t.Error("return the end")
}

func getInfo(id float32) (*Employee, error) {
	employee, err := callEmployee(id)
	if err != nil {
		return nil, fmt.Errorf("throw an error: %v", err)
	} else {
		log.Printf("the Employee is %#v\n", employee)
		return employee, err
	}
}

func callEmployee(id float32) (*Employee, error) {
	employee := Employee{LastName: "martin", MiddleName: "luther", FirstName: "king", Id: id}
	return &employee, nil
}
