package main

import "fmt"

func main() {
	var repsentSingler = make(map[string]string)
	repsentSingler["country music"] = "Taylor Swift"
	repsentSingler["pop music"] = "Jay Zhou"
	repsentSingler["Rock Stone"] = "Michael JackSon"
	repsentSingler["Blue Jazz"] = "Mars"
	fmt.Println(repsentSingler)

	delete(repsentSingler, "pop music")
	fmt.Println(repsentSingler)
}