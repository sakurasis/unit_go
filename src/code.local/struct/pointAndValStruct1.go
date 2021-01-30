package main

import "fmt"

// Triangle the struct in mathiethic
type Triangle struct {

	width float64
	height float64
}

func (t *Triangle) calcuateArea() float64{
	return float64(1)/float64(2) * t.width * t.height
}

func (t *Triangle) changeWidth(w float64) {
	t.width = w
	return
}

func main() {
	t := Triangle { width: 2, height: 3}
	var beforeArea float64 
	beforeArea = t.calcuateArea()
	fmt.Println("before:", beforeArea)
	t.changeWidth(4) 
	beforeArea = t.calcuateArea()
	fmt.Println("after:", beforeArea)
}