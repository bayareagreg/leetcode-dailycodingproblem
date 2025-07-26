package main

import "fmt"

type DequeNode struct {
	next  *DequeNode
	prev  *DequeNode
	start int
	turns int
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

func (d *Deque) append(start, turns int) {
	newLast := &DequeNode{
		next:  nil,
		prev:  d.last,
		start: start,
		turns: turns,
	}
	if d.last != nil {
		d.last.next = newLast
	}
	d.last = newLast
	if d.first == nil {
		d.first = newLast
	}
}

func (d *Deque) appendLeft(start, turns int) {
	newFirst := &DequeNode{
		next:  d.first,
		prev:  nil,
		start: start,
		turns: turns,
	}
	if d.first != nil {
		d.first.prev = newFirst
	}
	d.first = newFirst
	if d.last == nil {
		d.last = newFirst
	}
}

func (d *Deque) pop() (int, int) {
	if d.last == nil {
		panic("cannot pop from empty queue")
	}
	start, turns := d.last.start, d.last.turns
	if d.first == d.last {
		d.first = nil
	}
	d.last = d.last.prev
	if d.last != nil {
		d.last.next = nil
	}
	return start, turns
}

func (d *Deque) popLeft() (int, int) {
	if d.first == nil {
		panic("cannot popLeft from empty queue")
	}
	start, turns := d.first.start, d.first.turns
	if d.first == d.last {
		d.last = nil
	}
	d.first = d.first.next
	if d.first != nil {
		d.first.prev = nil
	}
	return start, turns
}

func (d *Deque) isEmpty() bool {
	return d.first == nil
}

func (d *Deque) peek() (int, int) {
	if d.last == nil {
		panic("cannot peek from empty queue")
	}
	return d.last.start, d.last.turns
}

func (d *Deque) peekLeft() (int, int) {
	if d.first == nil {
		panic("cannot peekLeft from empty queue")
	}
	return d.first.start, d.first.turns
}

func printDeque(d *Deque) {
	fmt.Printf("Deque: ")
	for n := d.first; n != nil; n = n.next {
		fmt.Printf("(%v,%v) ", n.start, n.turns)
	}
	fmt.Println()
}

func snakesAndLadders(board [][]int) int {

	//fmt.Printf("%v\n", board)

	mb := make(map[int]int)
	end := len(board) * len(board)
	for i := 1; i <= end; i++ {
		row := len(board) - 1 - (i-1)/len(board)
		col := (i - 1) % len(board)
		if (len(board)-row)%2 == 0 {
			col = len(board) - 1 - col
		}
		if board[row][col] == -1 {
			mb[i] = i
		} else {
			mb[i] = board[row][col]
		}
	}

	fmt.Printf("%v\n", mb)

	// BFS
	path := NewDeque()
	turns := 0
	start := 1
	path.append(start, turns)
	visited := make(map[int]struct{})
	var square int
	for !path.isEmpty() {
		//printDeque(path)
		square, turns = path.popLeft()

		fmt.Printf("Starting with square: %v, turns: %v", square, turns)

		fmt.Printf(". These are all the moves I can make with %v turns: ", turns+1)
		for moveTo := square + 1; moveTo < square+7; moveTo++ {
			if moveTo >= end {
				fmt.Printf("last move: %v, square: %v\n", moveTo, square)
				return turns + 1
			}
			if _, ok := visited[moveTo]; !ok {
				visited[moveTo] = struct{}{}
				path.append(mb[moveTo], turns+1)
				fmt.Printf("%v,", mb[moveTo])
				if mb[moveTo] >= end {
					return turns + 1
				}
			}
		}
		fmt.Println()
	}
	return -1
}

func main() {

	/*
			board := [][]int{
				{-1, -1, -1, -1, -1, -1},
				{-1, -1, -1, -1, -1, -1},
				{-1, -1, -1, -1, -1, -1},
				{-1, 35, -1, -1, 13, -1},
				{-1, -1, -1, -1, -1, -1},
				{-1, 15, -1, -1, -1, -1}}

					board := [][]int{
							{-1, -1, -1},
							{-1, 9, 8},
							{-1, 8, 9},
						}
					board := [][]int{
						{2, -1, -1, -1, -1},
						{-1, -1, -1, -1, -1},
						{-1, -1, -1, -1, -1},
						{-1, -1, -1, -1, -1},
						{-1, -1, -1, -1, -1},
					}
		board := [][]int{
			{-1, -1, -1, -1, -1, -1},
			{-1, -1, -1, -1, -1, -1},
			{-1, -1, -1, -1, -1, -1},
			{-1, 35, -1, -1, 13, -1},
			{-1, -1, -1, -1, -1, -1},
			{-1, 15, -1, -1, -1, -1}}

			board := [][]int{
				{-1, -1, 19, 10, -1},
				{2, -1, -1, 6, -1},
				{-1, 17, -1, 19, -1},
				{25, -1, 20, -1, -1},
				{-1, -1, -1, -1, 15},
			}
	*/

	board := [][]int{
		{-1, 1, 2, -1},
		{2, 13, 15, -1},
		{-1, 10, -1, -1},
		{-1, 6, 2, 8},
	}

	fmt.Printf("%v\n", snakesAndLadders(board))
}
