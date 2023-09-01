package main

import (
	"fmt"
	"reflect"
)

type User struct {
	Name string `alia:"name" info:"userName"`
	Code string `alia:"userCode"`
}

func getAllTag(arg interface{}) {
	e := reflect.TypeOf(arg).Elem()
	for i := 0; i < e.NumField(); i++ {
		alia := e.Field(i).Tag.Get("alia")
		info := e.Field(i).Tag.Get("info")
		fmt.Println("alia:", alia, "\tinfo:", info)
	}
}

func main() {
	var u User
	getAllTag(&u)
}
