package main

// Node define the node
type Node struct {
	Value int
	Next  *Node
}

// head init the node which located in the first
var head = new(Node)

// AddNode add node to collections.
func AddNode(t *Node, v int) int {
	if head == nil {
		temp := &Node{v, nil}
		head = temp
		return 0
	}
	// the current is equals the existed node.
	if v == t.Value {
		println("the node is existed. value: ", v)
		return -1
	}
	// the condition the next node of which is null.
	if t.Next == nil {
		t.Next = &Node{v, nil}
		return -2
	}
	return AddNode(t.Next, v)
}

// FindNode find the node from the collection.
func FindNode(t *Node, v int) bool {
	if head == nil {
		t = &Node{v, nil}
		head = t
		return false
	}

	if v == t.Value {
		return true
	}

	if t.Next == nil {
		return false
	}

	return FindNode(t, v)
}

func Size(t *Node) int {
	if t == nil {
		println("it's empty.")
		return 0
	}
	len := 0
	for t != nil {
		len++
		t = t.Next
	}
	return len
}
