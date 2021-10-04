package main

import "fmt"

// Book the object of Book which describe it's properties
// @name the book's name
// @price the price
// @count the number of book
type Book struct {
	Name   string
	Price  float32
	Count  int
	Author Author
}

// Author a object which named author,it describe a person
// whom wrote a book and become his(or her) career
type Author struct {
	Name      string
	Gender    int8
	Age       int8
	Represent string
}

func main() {

	fmt.Println("============the first initition Method=================")
	var fiction Book = Book{
		Name:  "The Three-Body Problem",
		Price: 35.55,
		Count: 30000,
	}
	fmt.Printf("fiction:%+v\n", fiction)
	fmt.Println("============the second initition Method=================")
	var cartoon Book
	cartoon.Name = "Gallop!"
	cartoon.Price = 85.0
	cartoon.Count = 9999999
	fmt.Printf("the cartoon:%+v\n", cartoon)
	fmt.Println("============nesting struct=================")
	book1 := Book{
		Name:  "the Dark forest",
		Price: 12.35,
		Count: 999999999,
		Author: Author{
			Name:      "CIXI LIU",
			Gender:    1,
			Age:       18,
			Represent: "The Three-Body Problem",
		},
	}
	fmt.Printf("the book detail:%+v\n", book1)

}
