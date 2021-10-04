package main

import "fmt"

type People interface {
	Speak(string) string
}

type Student struct{}

func (stu *Student) Speak(str string) (ctx string) {

	if str == "xb" {
		ctx = "学霸"
	} else {
		ctx = "学渣"
	}
	return
}

func main() {
	pep := Student{}
	ctx := "xz"
	var content string = pep.Speak(ctx)
	fmt.Println(content)
}
