package main

import "fmt"

func main() {
	var elements = make([]string, 4)
	elements[0] = "halo"
	elements[1] = "hello"
	elements[2] = "hi"
	elements[3] = "hei"
	fmt.Printf("length:%d,plain text:%s\n", len(elements), elements)
	// delete the element that index is the second element.
	elements = append(elements[:2], elements[2+1:]...)
	fmt.Printf("length:%d,plain text:%s\n", len(elements), elements)
	newElements := make([]string, 2)
	// copy another slice.
	copy(newElements, elements[1:])
	fmt.Println("the new element:", newElements)
}
