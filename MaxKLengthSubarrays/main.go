package main

import (
	"fmt"
	"math"
)

type Node struct {
	next *Node
	prev *Node
	data int
}

type Deque struct {
	first *Node
	last  *Node
}

func NewDeque() *Deque {
	return &Deque{
		first: nil,
		last:  nil,
	}
}

func (d *Deque) append(val int) {
	newLast := &Node{
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

func (d *Deque) appendLeft(val int) {
	newFirst := &Node{
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

func (d *Deque) pop() int {
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

func (d *Deque) popLeft() int {
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

func (d *Deque) peek() int {
	if d.last == nil {
		panic("cannot peek from empty queue")
	}
	return d.last.data
}

func (d *Deque) peekLeft() int {
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

func maxKLength(arr []int, k int) []int {
	res := make([]int, 0)
	for i := range len(arr) - 2 {
		m := math.MinInt32
		for j := i; j < i+k; j++ {
			m = max(m, arr[j])
		}
		res = append(res, m)
	}
	return res
}

func maxKLength2(arr []int, k int) []int {
	res := make([]int, 0)
	q := NewDeque()
	for i := range k {
		for !q.isEmpty() && arr[i] >= arr[q.peek()] {
			q.pop()
		}
		q.append(i)
	}

	// loop invariant: q is a list of indices where
	// their corresponding values are in descending order
	for i := k; i < len(arr); i++ {
		res = append(res, arr[q.peekLeft()])
		for !q.isEmpty() && q.peekLeft() <= i-k {
			q.popLeft()
		}
		for !q.isEmpty() && arr[i] >= arr[q.peek()] {
			q.pop()
		}
		q.append(i)
	}
	res = append(res, arr[q.peekLeft()])
	return res
}

func main() {

	input1 := []int{10, 5, 2, 7, 8, 7}
	fmt.Printf("Max K Length = %v\n", maxKLength2(input1, 3))

	/*
			q := NewDeque()
			q.append(4)
			q.append(5)
			q.appendLeft(6)

			printDeque(q)

			fmt.Printf("%v\n", q.popLeft())
			fmt.Printf("%v\n", q.pop())

			printDeque(q)
			fmt.Printf("%v\n", q.pop())

		q := NewDeque()
		q.append(1)
		q.append(2)

		printDeque(q)

		fmt.Printf("%v\n", q.pop())

		fmt.Printf("%v\n", q.pop())

		fmt.Printf("%v\n", q.isEmpty())
		//		fmt.Printf("%v\n", q.popLeft())
		//		fmt.Printf("%v\n", q.pop())
	*/

}
