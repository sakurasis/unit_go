package main

import (
	"fmt"
	"reflect"
)

func assembleSlice[T comparable](s []T)  {
	for _, v := range s {
		_, _ = fmt.Printf("%v\t", v)
	}
	_, _ = fmt.Print("\n")
}

func Sum(a,b interface{}) interface{} {
	if reflect.TypeOf(a).Kind() != reflect.TypeOf(b).Kind() {
        return nil
    }
    switch reflect.TypeOf(a).Kind() {
	case reflect.Int:
		return reflect.ValueOf(a).Int() + reflect.ValueOf(b).Int()
	case reflect.Float64:
		return reflect.ValueOf(a).Float() + reflect.ValueOf(b).Float()
	case reflect.String:
		return reflect.ValueOf(a).String() + reflect.ValueOf(b).String()
	default:
		return nil
	}
}

type Float interface {
	~float32 | ~float64
}

type Signed interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 
}

type Unsigned interface {
	~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 
}

type Integer interface {
	Signed | Unsigned
}

type Ordered interface {
	Integer | Float | ~string
}

func func1[T Ordered] (a T){
	fmt.Printf("current value is : %v\n", reflect.TypeOf(a))
}

func main() {
	func1("a")
	a :=Sum(1,2)
	b := Sum(1.1,2.1)
	c := Sum("1","2")
	fmt.Printf("a:%v\t b:%v c:%v\n", a, b, c)
	assembleSlice([]int{1,2,3,4})
	assembleSlice([]string{"champion", "edg"})
}
