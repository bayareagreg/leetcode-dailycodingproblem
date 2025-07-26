package main

import (
	"fmt"
)

type Node struct {
	next *Node
	data int
}

func NewNode(data int, next *Node) *Node {
	return &Node{
		data: data,
		next: next,
	}
}

func printLinkedList(node *Node) {
	for ; node != nil; node = node.next {
		fmt.Printf("%v\n", node.data)
	}
}

func rearrangeLinkedList(list1 *Node) {
	low := true
	for {
		if list1.next == nil {
			break
		}
		if low {
			if list1.data > list1.next.data {
				x := list1.data
				list1.data = list1.next.data
				list1.next.data = x
			}
		} else {
			if list1.data < list1.next.data {
				x := list1.data
				list1.data = list1.next.data
				list1.next.data = x
			}
		}
		list1 = list1.next
		low = !low
	}
}

func main() {

	x5 := NewNode(5, nil)
	x4 := NewNode(4, x5)
	x3 := NewNode(3, x4)
	x2 := NewNode(2, x3)
	x1 := NewNode(1, x2)

	printLinkedList(x1)
	rearrangeLinkedList(x1)
	printLinkedList(x1)
}
