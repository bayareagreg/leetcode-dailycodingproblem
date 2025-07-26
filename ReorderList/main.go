package main

import "fmt"

/**
 * Definition for singly-linked list.
 */
type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseLinkedList(node *ListNode) *ListNode {
	next := node.Next
	if next == nil {
		return node
	}
	next2 := next.Next
	node.Next = nil
	for {
		next.Next = node
		if next2 == nil {
			return next
		}
		node = next
		next = next2
		next2 = next2.Next
	}
}

func reorderList(head *ListNode) {
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow, fast = slow.Next, fast.Next.Next
	}

	// reverse 2nd half
	slow = reverseLinkedList(slow)

	first := head
	for slow.Next != nil {
		first.Next, first = slow, first.Next
		slow.Next, slow = first, slow.Next
	}
}

func printList(head *ListNode) {
	fmt.Printf("List: ")
	if head == nil {
		fmt.Printf("nil")
	} else {
		for head != nil {
			fmt.Printf("%v, ", head.Val)
			head = head.Next
		}
	}
	fmt.Println()
}

func main() {
	list := &ListNode{
		1,
		&ListNode{
			2,
			&ListNode{
				3,
				&ListNode{
					4,
					nil,
				},
			}}}
	reorderList(list)
	printList(list)
}
