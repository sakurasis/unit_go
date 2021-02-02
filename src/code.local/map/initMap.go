package main

import "fmt"

func main() {
	var representative = make(map[string]string)
	representative["country music"] = "Taylor Swift"
	representative["pop music"] = "Jay Zhou"
	representative["Rock Stone"] = "Michael JackSon"
	representative["Blue Jazz"] = "Mars"
	fmt.Println(representative)

	delete(representative, "pop music")
	fmt.Println(representative)
	val, ok := representative["pop music"]
	if ok {
		fmt.Println("the representative of pop music is",val)
	} else {
		fmt.Println("the singler of pop music is deleted.")
	}
	fmt.Println("-----------loop item by three diff ways---------------")
	for key := range representative {
		fmt.Println("key:",key,"value:",representative[key])
	}
	fmt.Println("----------------------------------")
	for key, val := range representative {
		fmt.Printf("%s -----> %s\n",key,val)
	}
	fmt.Println("----------------------------------")
}