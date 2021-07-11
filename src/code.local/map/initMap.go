package main

import "fmt"

func main() {
	fmt.Printf("--------- the first init method ----------")
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
		fmt.Println("the representative of pop music is", val)
	} else {
		fmt.Println("the singler of pop music is deleted.")
	}
	fmt.Println("-----------loop item by three diff ways---------------")
	for key := range representative {
		fmt.Println("key:", key, "value:", representative[key])
	}
	for key, val := range representative {
		fmt.Printf("%s -----> %s\n", key, val)
	}
	fmt.Println("----------------------------------")
	fmt.Println("------- the second init method --------")
	bookMap := map[string]string{}
	bookMap["spring in action"] = "java"
	bookMap["compiles:principles"] = "all programmer lang"
	fmt.Printf("the bookMap is %v\n", bookMap)
	fmt.Println("------- the third init method ---------")
	map1 := map[string]interface{}{}
	map1["name"] = "go programming design"
	map1["price"] = 1.23
	map1["description"] = []string{"scriptures", "cheaper", "understanding"}
	fmt.Printf("the map1: \t%v\n", map1)
	delete(map1, "name")
	fmt.Printf("the new map1: %v\n", map1)
}
