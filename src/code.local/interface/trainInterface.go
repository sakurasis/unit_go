package main

type Student struct{}

func Set(t interface{}) {

}

func Get(t *interface{}) {

}

func main() {
	s := Student{}

	Set(s)
	Set(&s)
}
