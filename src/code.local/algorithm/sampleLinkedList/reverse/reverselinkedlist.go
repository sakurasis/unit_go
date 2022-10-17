package main

type ListNode struct {
	Val  int
	Next *ListNode
}

// ReverseList
func ReverseList(pHead *ListNode) *ListNode {
	// write code here
	// get the linking of a next node.
	if pHead == nil || pHead.Next == nil {
		return pHead
	}

	var pre *ListNode
	for pHead != nil {

		pHead, pHead.Next, pre = pHead.Next, pre, pHead
	}
	return pre
}

func reverse(t *ListNode) (pre *ListNode) {
	if t == nil || t.Next == nil {
		return t
	}
	cur := t.Next
	pre = reverse(cur)
	cur.Next = t
	t.Next = nil
	return pre
}
