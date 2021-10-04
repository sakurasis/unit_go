package main

import (
	"fmt"
	"reflect"
)

func main() {
	var price float32 = 3.1415
	//typeFloat32 := reflect.float32
	reflectType(price)
	typeName := reflect.TypeOf(price)
	fmt.Println(typeName.Kind())
	// fmt.Println(typeFloat32 == typeName)
}

func reflectType(x interface{}) {
	v := reflect.TypeOf(x)
	fmt.Printf("type:%v\n", v)
}
