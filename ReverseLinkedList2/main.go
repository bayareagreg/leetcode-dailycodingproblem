package main

import "fmt"

/**
 * Definition for singly-linked list.
 */
type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseBetween(head *ListNode, left int, right int) *ListNode {
	if left == right {
		return head
	}

	saveHead := head
	pos := 1
	var prev *ListNode
	for head != nil && pos < left {
		pos++
		prev = head
		head = head.Next
	}

	stack := make([]*ListNode, right-left+1)
	i := 0
	// push nodes onto the stack
	for head != nil && pos <= right {
		stack[i] = head
		i++
		pos++
		head = head.Next
	}

	// now pop things off the stack, and set Next pointer as we do it
	for i := len(stack) - 1; i > 0; i-- {
		stack[i].Next = stack[i-1]
	}

	// point last element in the stack to the rest of the list
	stack[0].Next = head

	if left > 1 {
		// make head point to the bottom of stack
		prev.Next = stack[len(stack)-1]
	} else {
		saveHead = stack[len(stack)-1]
	}

	return saveHead
}

func printList(head *ListNode) {
	for head != nil {
		fmt.Printf("%v ", head.Val)
		head = head.Next
	}
	fmt.Println()
}

func main() {
	//head := &ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, &ListNode{5, nil}}}}}
	//head := &ListNode{5, nil}
	head := &ListNode{3, &ListNode{5, nil}}
	printList(head)
	printList(reverseBetween(head, 1, 2))

}
