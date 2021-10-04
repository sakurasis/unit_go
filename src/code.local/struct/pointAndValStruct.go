package main

import "fmt"

// Drink  describe a drink contains some properties.
type Drink struct {
	Name string
	Ice  bool
}

func main() {
	juice := Drink{
		Name: "orange juice",
		Ice:  false,
	}
	// redTea := juice
	redTea := &juice
	redTea.Ice = true

	fmt.Printf("juice:%+v\n", juice)
	fmt.Printf("red_tea:%+v\n", *redTea)
	fmt.Printf("the address of juice:%p\n", &juice)
	fmt.Printf("the address of red tea:%p\n", redTea)
}
