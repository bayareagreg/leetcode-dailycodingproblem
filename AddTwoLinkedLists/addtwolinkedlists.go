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

func addTwoLinkedLists(list1 *Node, list2 *Node) *Node {

	carry := 0
	var n *Node = nil
	var root *Node = nil
	for {
		if list1 == nil && list2 == nil && carry == 0 {
			break
		}
		d1 := 0
		if list1 != nil {
			d1 = list1.data
		}
		d2 := 0
		if list2 != nil {
			d2 = list2.data
		}
		data := (d1 + d2 + carry) % 10
		if d1+d2+carry > 10 {
			carry = 1
		} else {
			carry = 0
		}

		n2 := NewNode(data, nil)
		if n != nil {
			n.next = n2
		}
		n = n2
		if root == nil {
			root = n2
		}
		if list1 != nil {
			list1 = list1.next
		}
		if list2 != nil {
			list2 = list2.next
		}
	}
	return root
}

func main() {
	/*
		n := NewNode(9, nil)
		list1 := NewNode(9, n)

		n2 := NewNode(2, nil)
		list2 := NewNode(5, n2)

		_ = addTwoLinkedLists(list1, list2)
		printLinkedList(nX)
	*/

	x1 := NewNode(1, nil)
	x2 := NewNode(4, x1)
	x3 := NewNode(3, x2)

	y1 := NewNode(9, nil)
	y2 := NewNode(8, y1)

	z := addTwoLinkedLists(x3, y2)
	printLinkedList(z)
}
