package main

import (
	"fmt"
)

func loopList(t *Node) {
	if t == nil {
		println("Empty List")
	}
	for t != nil {
		cur := t.Value
		fmt.Printf("%d -> ", cur)
		t = t.Next
	}
	fmt.Println()
}

func main() {
	head = nil
	AddNode(head, 1)
	AddNode(head, 2)
	AddNode(head, 3)
	AddNode(head, 2)
	AddNode(head, 5)
	AddNode(head, 6)
	loopList(head)
}
