package main

import "fmt"

type DequeNode struct {
	next *DequeNode
	prev *DequeNode
	data *TreeNode
}

type Deque struct {
	first *DequeNode
	last  *DequeNode
}

func NewDeque() *Deque {
	return &Deque{
		first: nil,
		last:  nil,
	}
}

func (d *Deque) append(val *TreeNode) {
	newLast := &DequeNode{
		next: nil,
		prev: d.last,
		data: val,
	}
	if d.last != nil {
		d.last.next = newLast
	}
	d.last = newLast
	if d.first == nil {
		d.first = newLast
	}
}

func (d *Deque) appendLeft(val *TreeNode) {
	newFirst := &DequeNode{
		next: d.first,
		prev: nil,
		data: val,
	}
	if d.first != nil {
		d.first.prev = newFirst
	}
	d.first = newFirst
	if d.last == nil {
		d.last = newFirst
	}
}

func (d *Deque) pop() *TreeNode {
	if d.last == nil {
		panic("cannot pop from empty queue")
	}
	val := d.last.data
	if d.first == d.last {
		d.first = nil
	}
	d.last = d.last.prev
	if d.last != nil {
		d.last.next = nil
	}
	return val
}

func (d *Deque) popLeft() *TreeNode {
	if d.first == nil {
		panic("cannot popLeft from empty queue")
	}
	val := d.first.data
	if d.first == d.last {
		d.last = nil
	}
	d.first = d.first.next
	if d.first != nil {
		d.first.prev = nil
	}
	return val
}

func (d *Deque) isEmpty() bool {
	return d.first == nil
}

func (d *Deque) peek() *TreeNode {
	if d.last == nil {
		panic("cannot peek from empty queue")
	}
	return d.last.data
}

func (d *Deque) peekLeft() *TreeNode {
	if d.first == nil {
		panic("cannot peekLeft from empty queue")
	}
	return d.first.data
}

func printDeque(d *Deque) {
	fmt.Printf("Deque: ")
	for n := d.first; n != nil; n = n.next {
		fmt.Printf("%v ", n.data)
	}
	fmt.Println()
}
