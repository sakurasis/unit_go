package main

import (
	"fmt"
	"unsafe"
)

func main() {
	fmt.Println(unsafe.Sizeof(bool(true)))    // 1
	fmt.Println(unsafe.Sizeof(float32(1.12))) // 4
	fmt.Println(unsafe.Sizeof(int8(123)))     // 1
	fmt.Println(unsafe.Sizeof(int16(123)))    // 2
	fmt.Println(unsafe.Sizeof(int32(123)))    // 4
	fmt.Println(unsafe.Sizeof(int64(123)))    // 8
	fmt.Println(unsafe.Sizeof(uint(123)))     // 8
	fmt.Println(unsafe.Sizeof(uint8(123)))
	fmt.Println(unsafe.Sizeof(uint16(123)))
	fmt.Println(unsafe.Sizeof(uintptr(123)))       // 8
	fmt.Println(unsafe.Sizeof(int(1)))             // 8
	fmt.Println(unsafe.Sizeof(printSize()))        // 8
	fmt.Println(unsafe.Sizeof(string("sakura")))   // 16
	fmt.Println(unsafe.Sizeof([]string{"sakura"})) // 24
}

func printSize() int {
	return 1
}
