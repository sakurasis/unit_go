package main

import (
	"fmt"
	"reflect"
)

type User struct {
	Id     int64
	Name   string
	Gender bool
}

func (u *User) GetUserString() string {
	return "user:{" + fmt.Sprintf("id:%d,", u.Id) + fmt.Sprintf("name:%s,", u.Name) +
		fmt.Sprintf("gender:%v}", u.Gender)
}

func main() {
	var price float32 = 3.1415
	//typeFloat32 := reflect.float32
	reflectType(price)
	typeName := reflect.TypeOf(price)
	fmt.Println(typeName.Kind())
	// fmt.Println(typeFloat32 == typeName)
	fmt.Println("--------------------------------------")
	user := User{0, "zhangsan", true}
	reflectFieldsAndMethods(user)
}

func reflectType(x interface{}) {
	v := reflect.TypeOf(x)
	fmt.Printf("type:%v\n", v)
}

func reflectFieldsAndMethods(args interface{}) {
	t := reflect.TypeOf(args)
	fmt.Println("input type is:", t.Name())
	v := reflect.ValueOf(args)
	fmt.Println("input value is:", v)

	//
	for i := 0; i < t.NumField(); i++ {
		f := t.Field(i)
		name := v.Field(i).Interface()
		fmt.Printf("名称：%s\t类型: %v\t值: %v\n", f.Name, f.Type, name)
	}
	for i := 0; i < t.NumMethod(); i++ {
		m := t.Method(i)
		fmt.Printf("name:%s, type: %v\n", m.Name, m.Type)
	}
}
