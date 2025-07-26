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

func (n *Node) Length() int {
	l := 0
	for {
		if n == nil {
			return l
		}
		n = n.next
		l++
	}
}

func findIntersectingNode(list1 *Node, list2 *Node) *Node {
	len1 := list1.Length()
	len2 := list2.Length()
	dif := len1 - len2
	if dif < 0 {
		for i := dif; i < 0; i++ {
			list2 = list2.next
		}
	} else if dif > 0 {
		for i := 0; i < dif; i++ {
			list1 = list1.next
		}
	}
	for {
		if list1.data == list2.data {
			break
		}
		list1 = list1.next
		list2 = list2.next
	}
	return list1
}

func printLinkedList(node *Node) {
	fmt.Printf("LinkedList: ")
	for ; node != nil; node = node.next {
		fmt.Printf("%v ", node.data)
	}
	fmt.Println()
}

func main() {

	x5 := NewNode(10, nil)
	x4 := NewNode(8, x5)
	x3 := NewNode(7, x4)
	x2 := NewNode(3, x3)

	x1 := NewNode(1, x4)
	x0 := NewNode(99, x1)

	printLinkedList(x2)
	printLinkedList(x0)

	printLinkedList(findIntersectingNode(x2, x0))
}
