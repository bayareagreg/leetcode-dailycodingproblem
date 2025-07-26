package main

import (
	"fmt"
	"strconv"
)

/**
 * Definition for singly-linked list.
 */
type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	saveHead := head
	deque := make([]*ListNode, n+1)
	i := 0
	for head != nil {
		if i >= n+1 {
			for j := 1; j <= n; j++ {
				deque[j-1] = deque[j]
			}
			i = n
		}
		deque[i] = head
		head = head.Next
		i++
	}

	if deque[n] == nil {
		return deque[1]
	} else if n >= 2 {
		deque[0].Next = deque[2]
		return saveHead
	} else if n == 1 {
		deque[0].Next = nil
		return saveHead
	} else {
		panic("Invalid n: " + strconv.Itoa(n))
	}
}

func printList(head *ListNode) {
	fmt.Printf("List: ")
	for head != nil {
		fmt.Printf("%v, ", head.Val)
		head = head.Next
	}
	fmt.Println()
}

func main() {
	//list := &ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, &ListNode{5, nil}}}}}
	//	list := &ListNode{1, &ListNode{2, &ListNode{3, nil}}}
	//	list := &ListNode{1, &ListNode{2, nil}}
	list := &ListNode{1, nil}
	printList(list)
	printList(removeNthFromEnd(list, 1))
}
