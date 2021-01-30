package main

import "fmt"

func main() {
	const a = 96
	var intVal int = a
	var int32Val int32 = a
	var float64Val float64 = a
	var complex64Val complex64 = a
	fmt.Println("intVal is", intVal, "\nint32Val is", int32Val, 
	"\nfloat64Val is", float64Val, "\ncomplex64Val is", complex64Val)
}