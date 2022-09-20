//go:build ignore
// +build ignore

package main

import "fmt"

func main() {
	firstList := &LinkedList{9, &LinkedList{9, &LinkedList{9, nil}}}
	nextList := &LinkedList{1, nil}
	finalList := addTwoNumbers(firstList, nextList)
	result := fmt.Sprintf("%+v", finalList)
	fmt.Print(result)
}

// LinkedList
/**
a struct of singly-linked list.
*/
type LinkedList struct {
	Val  int
	Next *LinkedList
}

/**
addTwoNumbers 	The digits are stored in reverse order and each of their nodes contain a single digit.
				Add the two numbers and return it as a linked list.
*LinkedList		a linked list can be a single digit.
*/
func addTwoNumbers(fir *LinkedList, nxt *LinkedList) *LinkedList {
	// init the head(point to the real one)
	head := &LinkedList{Val: 0}
	firstNumb, nextNumb, temp, final := 0, 0, 0, head

	for fir != nil || nxt != nil || temp != 0 {
		//
		if fir == nil {
			firstNumb = 0
		} else {
			firstNumb = fir.Val
			fir = fir.Next
		}

		if nxt == nil {
			nextNumb = 0
		} else {
			nextNumb = nxt.Val
			nxt = nxt.Next
		}
		// get the number in the last bit
		final.Next = &LinkedList{Val: (firstNumb + nextNumb + temp) % 10}
		final = final.Next
		temp = (firstNumb + nextNumb + temp) / 10
	}
	return head.Next
}
