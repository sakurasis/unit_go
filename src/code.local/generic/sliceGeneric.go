// build ignore
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

type CusError interface{
	Error() string
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

type CusInterface [T int|string] interface {
    int | string
    SetName(d T) T
    GetName() T
}

type CusInt int

type CusInteger interface {
	~int | interface{CusInt}
}

type CusStr string

func (c CusStr) SetName(n string) string {
	return n
}

type Str struct {
	Name string
}

type Reader interface {
	Read(p []byte) (n int, err error)
}

type Writer interface {
	Write(p []byte) (n int, err error)
}

type CusRead[T interface{*Reader | *Writer}] []T



func main() {

	sum := func[T int|float32|float64|string] (a, b T) T {
		return a +b;
	}
	fmt.Println(sum(1, 2))

	func1("a")
	assembleSlice([]int{1,2,3,4})
	assembleSlice([]string{"champion", "edg"})
}

