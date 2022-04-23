package main

import "fmt"

type Node struct {
	Value int
	Next  *Node
}

var head = new(Node)
var end *Node

func main() {
	for i := 0; i < 5; i++ {

		n := &Node{
			Value: i + 1,
		}
		AddNode(n)
	}

	josephus(1, 3)
}

func AddNode(t *Node) {
	if end == nil {
		head = t
		t.Next = head
		end = t
	} else {
		end.Next = t
		t.Next = head
		end = t
	}
}

//josephus k is the beginning number util count the number of n is out of the collection.
func josephus(k, n int) {
	count := 1

	for i := 0; i < k-1; i++ {
		head = head.Next
		end = end.Next
	}

	for {
		count++
		head = head.Next
		end = end.Next

		if count == n {
			fmt.Println(head.Value, " is out.")
			end.Next = head.Next
			head = head.Next
			count = 1
		}

		if head == end {
			fmt.Println(head.Value, " still in, and finally alive.")
			break
		}
	}
}
