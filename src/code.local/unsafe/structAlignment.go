package main

import (
	"fmt"
	"unsafe"
)

func main() {
	a := Animal{123, 123, 18}
	fmt.Println("the first offset:", unsafe.Offsetof(a.id))
	fmt.Println("the first size:", unsafe.Sizeof(a.id))
	fmt.Println("the second offset:", unsafe.Offsetof(a.num))
	fmt.Println("the second size:", unsafe.Sizeof(a.num))
	fmt.Println("the third offset:", unsafe.Offsetof(a.age))
	fmt.Println("the third size:", unsafe.Sizeof(a.age))
	fmt.Println("the total size:", unsafe.Sizeof(a))
	fmt.Println("----------- the optimized order ---------")
	d := Dog{123, 123, 18}
	fmt.Println(unsafe.Sizeof(d))
	fmt.Println("id's size:", unsafe.Sizeof(d.id))
	fmt.Println("age's offset:", unsafe.Offsetof(d.age))
	fmt.Println("age size:", unsafe.Sizeof(d.age))
	fmt.Println("num's offset:", unsafe.Offsetof(d.num))
	fmt.Println("num's size:", unsafe.Sizeof(d.num))
}

type Animal struct {
	id  uint16
	num uint64
	age uint16
}

type Dog struct {
	id  uint16
	age uint16
	num uint64
}
