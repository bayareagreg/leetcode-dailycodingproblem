// To execute Go code, please declare a func main() in a package "main"

package main

import "fmt"

type DequeData struct {
	word string
	path []string
}

type DequeNode struct {
	next *DequeNode
	prev *DequeNode
	data DequeData
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

func (d *Deque) append(val DequeData) {
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

func (d *Deque) appendLeft(val DequeData) {
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

func (d *Deque) pop() DequeData {
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

func (d *Deque) popLeft() DequeData {
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

func (d *Deque) peek() DequeData {
	if d.last == nil {
		panic("cannot peek from empty queue")
	}
	return d.last.data
}

func (d *Deque) peekLeft() DequeData {
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

func word_ladder(start string, end string, words []string) []string {
	// lets create a map for words for fast access
	mwords := make(map[string]struct{})
	for _, w := range words {
		mwords[w] = struct{}{}
	}

	queue := NewDeque()
	queue.append(DequeData{
		start,
		[]string{start},
	})

	for !queue.isEmpty() {
		data := queue.popLeft()

		if data.word == end {
			return data.path
		}

		for i := range data.word {
			for c := 'a'; c <= 'z'; c++ {
				next_word := data.word[:i] + string(c) + data.word[i+1:]
				if _, ok := mwords[next_word]; ok {
					delete(mwords, next_word)
					queue.append(DequeData{
						next_word,
						append(data.path, next_word),
					})
				}
			}
		}
	}

	return nil
}

func main() {
	fmt.Printf("%v\n", word_ladder("dog", "cat", []string{"dot", "dop", "dat", "cat"}))
}
